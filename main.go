package main

import (
	"fmt"

	"github.com/redbeestudios/go-seed/cmd"
	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
	"github.com/redbeestudios/go-seed/pkg"
)

func main() {
	env, err := pkg.NewEnv("dev")
	if err != nil {
		panic(fmt.Sprintf("error creating environment: %s", err.Error()))
	}

	config := cmd.InitConfig(env)
	deps := cmd.InitDependencies(config)
	router := cmd.InitRoutes(deps)
	pokemon.InitPokemonTypes()

	cmd.StartServer(config, router)
}
