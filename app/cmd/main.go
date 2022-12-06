package main

import (
	"os"
	"strconv"
)

const initFlag = "--init"
const envRedisAddrKey = "redisAddr"
const envRedisPassKey = "redisPass"
const envRedisDbKey = "redisDb"
const envListenAddrKey = "listenAddr"

func main() {
	opts := new(ServerOptions)
	opts.RedisAddr = os.Getenv(envRedisAddrKey)
	opts.RedisDB, _ = strconv.Atoi(os.Getenv(envRedisDbKey))
	opts.RedisPass = os.Getenv(envRedisPassKey)
	opts.ServerListenAddr = os.Getenv(envListenAddrKey)

	if len(os.Args) > 1 && os.Args[1] == initFlag {
		err := FillDB(opts)
		if err != nil {
			os.Exit(-1)
		}
		return
	}

	server, err := CreateServer(opts)
	if err != nil {
		return
	}

	server.ListenAndServe()
}
