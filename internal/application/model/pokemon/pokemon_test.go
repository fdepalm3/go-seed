package pokemon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPokemon(t *testing.T) {
	id := 3
	name := "venusaur"
	primaryType := MustBuildPokemonType(Grass.String())
	secondaryType := MustBuildPokemonType(Rock.String())

	p := NewPokemon(id, name, primaryType, &secondaryType)

	assert.Equal(t, id, p.Id())
	assert.Equal(t, name, p.Name())
	assert.Equal(t, primaryType, p.PrimaryType())
	assert.Equal(t, &secondaryType, p.SecondaryType())
}
