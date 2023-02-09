package cmd

import (
	"github.com/go-chi/chi/v5"
)

func InitRoutes(dependencies *Dependencies) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/pokemon/{name}", dependencies.PokemonController.GetPokemon)
	r.Get("/attack/{name}", dependencies.PokemonController.GetAttack)
	r.Post("/dumpPokemons", dependencies.PokemonController.DumpPokemons)
	r.Post("/dumpPokemonsGoRoutine", dependencies.PokemonController.DumpPokemonsWithGoRoutines)

	return r
}
