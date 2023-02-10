package pokemon

type Pokemon struct {
	id                   int
	name                 string
	primaryType          *PokemonType
	pokemonSecondaryType *PokemonType
}

func NewPokemon(
	id int,
	name string,
	primaryType *PokemonType,
	secondaryType *PokemonType,
) *Pokemon {
	return &Pokemon{
		id:                   id,
		name:                 name,
		primaryType:          primaryType,
		pokemonSecondaryType: secondaryType,
	}
}

func (p *Pokemon) Id() int {
	return p.id
}

func (p *Pokemon) Name() string {
	return p.name
}

func (p *Pokemon) PrimaryType() *PokemonType {
	return p.primaryType
}

func (p *Pokemon) SecondaryType() *PokemonType {
	return p.pokemonSecondaryType
}
