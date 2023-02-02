package testdata

import (
	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
	"github.com/redbeestudios/go-seed/pkg"
)

func Pokemon() *pokemon.Pokemon {
	return pokemon.NewPokemon(
		3,
		"venusaur",
		pokemon.MustBuildPokemonType(pokemon.Grass.String()),
		pkg.ToPointer(pokemon.MustBuildPokemonType(pokemon.Rock.String())),
	)
}
