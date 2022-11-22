package main

import (
	"wheel/db"
	"wheel/redis"
	"wheel/server"
)

func main() {
	db.Init()
	redis.Init()

	server.Init()
}
