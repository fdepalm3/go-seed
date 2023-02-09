package pokemon

import "fmt"

type Type struct {
	Name        string
	Weaknesses  []*Type
	Resistances []*Type
	Inmunities  []*Type
}

// 	Water    Type = "water"
// 	Normal   Type = "normal"
// 	Flying   Type = "flying"
// 	Fighting Type = "fighting"
// 	Poison   Type = "poison"
// 	Electric Type = "electric"
// 	Ground   Type = "ground"
// 	Rock     Type = "rock"
// 	Psychic  Type = "psychic"
// 	Ice      Type = "ice"
// 	Bug      Type = "bug"
// 	Ghost    Type = "ghost"
// 	Steel    Type = "steel"
// 	Dragon   Type = "dragon"
// 	Dark     Type = "dark"
// 	Fairy    Type = "fairy"
// 	Invalid  Type = ""
// )

var allowedTypes = map[string]*Type{

	"fire": &Type{
		Name: "fire",
	},
	"grass": &Type{
		Name: "grass",
	},
	// Water.String():    Water,
	// Normal.String():   Normal,
	// Flying.String():   Flying,
	// Fighting.String(): Fighting,
	// Poison.String():   Poison,
	// Electric.String(): Electric,
	// Ground.String():   Ground,
	// Rock.String():     Rock,
	// Psychic.String():  Psychic,
	// Ice.String():      Ice,
	// Bug.String():      Bug,
	// Ghost.String():    Ghost,
	// Steel.String():    Steel,
	// Dragon.String():   Dragon,
	// Dark.String():     Dark,
	// Fairy.String():    Fairy,
	// Invalid.String():  Invalid,
}

func NewPokemonType(pokemonType string) (*Type, error) {
	if tipo, existe := allowedTypes[pokemonType]; existe {
		return tipo, nil
	}
	return nil, fmt.Errorf("Invalid pokemon type: %s", pokemonType)
}

func InitPokemonTypes() {
	allowedTypes["fire"].Resistances = []*Type{
		allowedTypes["grass"],
	}
	allowedTypes["grass"].Weaknesses = []*Type{
		allowedTypes["fire"],
	}
}

// Testing purpose functions/methods
func MustBuildPokemonType(pokemonType string) Type {
	p, err := NewPokemonType(pokemonType)
	if err != nil {
		panic(err)
	}
	return *p
}
