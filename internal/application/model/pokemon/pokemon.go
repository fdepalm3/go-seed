package pokemon

type Pokemon struct {
	id                   int
	name                 string
	primaryType          Type
	pokemonSecondaryType *Type
}

func NewPokemon(
	id int,
	name string,
	primaryType Type,
	secondaryType *Type,
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

func (p *Pokemon) PrimaryType() Type {
	return p.primaryType
}

func (p *Pokemon) SecondaryType() *Type {
	return p.pokemonSecondaryType
}
