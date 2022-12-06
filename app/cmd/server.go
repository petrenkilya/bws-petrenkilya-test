package main

import (
	"bws_test/app/handlers"
	"bws_test/app/repositories/redis"
	"bws_test/app/usecase/impl"
	redis2 "github.com/go-redis/redis"
	"github.com/gofiber/fiber"
	"log"
)

const hackersPath = "/json/hackers"
const hackersRedisKey = "hackers"
const redisDefaultAddr = "127.0.0.1:6379"
const serverDefaultListenAddr = ":8010"

type Server struct {
	app        *fiber.App
	logger     *log.Logger
	listenAddr string
}

type ServerOptions struct {
	RedisAddr        string
	RedisPass        string
	RedisDB          int
	ServerListenAddr string
}

func parseOpts(opt *ServerOptions) *ServerOptions {
	if opt == nil {
		opt = &ServerOptions{RedisAddr: redisDefaultAddr, ServerListenAddr: serverDefaultListenAddr}
		return opt
	}
	if len(opt.RedisAddr) == 0 {
		opt.RedisAddr = redisDefaultAddr
	}
	if len(opt.ServerListenAddr) == 0 {
		opt.ServerListenAddr = serverDefaultListenAddr
	}
	return opt
}

func CreateServer(opt *ServerOptions) (*Server, error) {
	opt = parseOpts(opt)
	app := fiber.New(&fiber.Settings{
		Prefork:      true,
		ServerHeader: "Fiber",
	})
	logger := log.Default()

	redisClient := redis2.NewClient(&redis2.Options{Addr: opt.RedisAddr, Password: opt.RedisPass, DB: opt.RedisDB})
	err := redisClient.Ping().Err()
	if err != nil {
		logger.Printf("redisClient ping error: %v", err)
		return nil, err
	}

	hackersRepo := redis.CreateHackersRepo(redisClient, hackersRedisKey)
	hackersUseCase := impl.CreateHackersUseCase(hackersRepo)
	hackersHandler := handlers.CreateHackersHandler(hackersUseCase, logger)

	app.Get(hackersPath, hackersHandler.Get)

	return &Server{app: app, logger: logger, listenAddr: opt.ServerListenAddr}, nil
}

func (s *Server) ListenAndServe() {
	s.logger.Fatal(s.app.Listen(s.listenAddr))
}
