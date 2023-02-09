package pokemon

import "github.com/redbeestudios/go-seed/internal/application/model/pokemon"

type pokemonResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type moveResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func pokemonFromDomain(pokemon *pokemon.Pokemon) *pokemonResponse {
	return &pokemonResponse{
		Id:   pokemon.Id(),
		Name: pokemon.Name(),
		Type: pokemon.PrimaryType().Name,
	}
}

func moveFromDomain(move *pokemon.Move) *moveResponse {
	return &moveResponse{
		Id:   move.Id(),
		Name: move.Name(),
		Type: move.MoveType().Name,
	}
}
