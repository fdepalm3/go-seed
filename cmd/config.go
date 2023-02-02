package cmd

import (
	"fmt"

	"github.com/redbeestudios/go-seed/pkg"
)

type Config struct {
	Server  pkg.ServerConfig     `json:"server"`
	PokeApi pkg.RestClientConfig `json:"poke_api"`
	Redis   pkg.RedisConfig      `json:"redis"`
}

func InitConfig(env pkg.Env) *Config {
	config := &Config{}

	if err := pkg.ReadJsonFile(env.String(), config); err != nil {
		panic(fmt.Sprintf("Error initializing config: %s", env.String()))
	}

	return config
}
