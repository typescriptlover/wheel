package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type config struct {
	DB_URI string `env:"DB_URI"`
}

var Config *config

func Init() {
	c := &config{}
	opts := env.Options{RequiredIfNoDef: true}

	if err := env.Parse(c, opts); err != nil {
		log.Fatalf("%+v\n", err)
	}

	Config = c
}
