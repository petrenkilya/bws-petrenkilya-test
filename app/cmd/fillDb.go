package main

import (
	"bws_test/app/models"
	"bws_test/app/repositories/redis"
	redis2 "github.com/go-redis/redis"
)

var hackersList = [...]models.Hacker{{"Alan Turing", 1912}, {"Claude Shannon", 1916},
	{"Alan Kay", 1940}, {"Richard Stallman", 1953}, {"Yukihiro Matsumoto", 1965},
	{"Linus Torvalds", 1969}}

func FillDB(opts *ServerOptions) error {
	opts = parseOpts(opts)
	redisClient := redis2.NewClient(&redis2.Options{Addr: opts.RedisAddr, Password: opts.RedisPass, DB: opts.RedisDB})
	err := redisClient.Ping().Err()
	if err != nil {
		return err
	}

	hackersRepo := redis.CreateHackersRepo(redisClient, hackersRedisKey)
	for _, hacker := range hackersList {
		err = hackersRepo.Create(hacker)
		if err != nil {
			return err
		}
	}
	return nil
}
