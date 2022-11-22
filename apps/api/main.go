package main

import (
	"wheel/config"
	"wheel/db"
	"wheel/redis"
	"wheel/server"
)

func main() {
	config.Init()
	db.Init()
	redis.Init()

	server.Init()
}
