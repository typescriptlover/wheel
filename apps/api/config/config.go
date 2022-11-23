package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type config struct {
	DB_URI     string `env:"DB_URI"`
	JWT_SECRET string `env:"JWT_SECRET"`
}

var Config *config

func Init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	c := &config{}
	opts := env.Options{RequiredIfNoDef: true}

	if err := env.Parse(c, opts); err != nil {
		log.Fatalf("%+v\n", err)
	}

	Config = c
}

func GetConfig() *config {
	return Config
}
