package pokemon

import (
	"testing"

	"github.com/redbeestudios/go-seed/testdata"
	"github.com/stretchr/testify/assert"
)

func TestFromDomain(t *testing.T) {
	pokemon := testdata.Pokemon()

	dto := pokemonFromDomain(pokemon)

	assert.Equal(t, pokemon.Id(), dto.Id)
	assert.Equal(t, pokemon.Name(), dto.Name)
	assert.Equal(t, pokemon.PrimaryType().String(), dto.Type)
}
