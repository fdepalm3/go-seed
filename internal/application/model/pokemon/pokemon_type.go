package pokemon

type PokemonType struct {
	Name        string
	Weaknesses  []*PokemonType
	Resistances []*PokemonType
	Immunities  []*PokemonType
}

const (
	Fire     string = "fire"
	Grass    string = "grass"
	Water    string = "water"
	Normal   string = "normal"
	Flying   string = "flying"
	Fighting string = "fighting"
	Poison   string = "poison"
	Electric string = "electric"
	Ground   string = "ground"
	Rock     string = "rock"
	Psychic  string = "psychic"
	Ice      string = "ice"
	Bug      string = "bug"
	Ghost    string = "ghost"
	Steel    string = "steel"
	Dragon   string = "dragon"
	Dark     string = "dark"
	Fairy    string = "fairy"
	invalid  string = ""
)

var allowedTypes = map[string]*PokemonType{
	Fire: {
		Name: Fire,
	},
	Grass: {
		Name: Grass,
	},
	Water: {
		Name: Water,
	},
	Normal: {
		Name: Normal,
	},
	Flying: {
		Name: Flying,
	},
	Fighting: {
		Name: Fighting,
	},
	Poison: {
		Name: Poison,
	},
	Electric: {
		Name: Electric,
	},
	Ground: {
		Name: Ground,
	},
	Rock: {
		Name: Rock,
	},
	Psychic: {
		Name: Psychic,
	},
	Ice: {
		Name: Ice,
	},
	Bug: {
		Name: Bug,
	},
	Ghost: {
		Name: Ghost,
	},
	Steel: {
		Name: Steel,
	},
	Dragon: {
		Name: Dragon,
	},
	Dark: {
		Name: Dark,
	},
	Fairy: {
		Name: Fairy,
	},
	invalid: {Name: invalid},
}

func (p *PokemonType) Multiplier(t *PokemonType) float64 {
	multiplier := 1.0
	for _, v := range p.Resistances {
		if v == t {
			multiplier = multiplier / 2
			return multiplier
		}
	}
	for _, v := range p.Weaknesses {
		if v == t {
			multiplier = multiplier * 2
			return multiplier
		}
	}

	for _, v := range p.Immunities {
		if v == t {
			multiplier = 0
			return multiplier
		}
	}

	return multiplier
}

func Type(pokemonType string) *PokemonType {
	if tipo, existe := allowedTypes[pokemonType]; existe {
		return tipo
	}
	return allowedTypes[invalid]
}

func InitPokemonTypes() {
	Type(Fire).Weaknesses = []*PokemonType{Type(Water), Type(Ground), Type(Rock)}
	Type(Fire).Resistances = []*PokemonType{Type(Bug), Type(Steel), Type(Fire), Type(Grass), Type(Ice), Type(Fairy)}
	Type(Grass).Weaknesses = []*PokemonType{Type(Flying), Type(Poison), Type(Bug), Type(Fire), Type(Ice)}
	Type(Grass).Resistances = []*PokemonType{Type(Ground), Type(Water), Type(Grass), Type(Electric)}
	Type(Water).Weaknesses = []*PokemonType{Type(Grass), Type(Electric)}
	Type(Water).Resistances = []*PokemonType{Type(Steel), Type(Fire), Type(Water), Type(Ice)}
	Type(Normal).Weaknesses = []*PokemonType{Type(Fighting)}
	Type(Normal).Immunities = []*PokemonType{Type(Ghost)}
	Type(Flying).Weaknesses = []*PokemonType{Type(Rock), Type(Electric), Type(Ice)}
	Type(Flying).Resistances = []*PokemonType{Type(Fighting), Type(Bug), Type(Grass)}
	Type(Flying).Immunities = []*PokemonType{Type(Ground)}
	Type(Fighting).Weaknesses = []*PokemonType{Type(Flying), Type(Psychic), Type(Fairy)}
	Type(Fighting).Resistances = []*PokemonType{Type(Rock), Type(Bug), Type(Dark)}
	Type(Poison).Weaknesses = []*PokemonType{Type(Psychic), Type(Ground)}
	Type(Poison).Resistances = []*PokemonType{Type(Fighting), Type(Poison), Type(Bug), Type(Grass), Type(Fairy)}
	Type(Electric).Weaknesses = []*PokemonType{Type(Ground)}
	Type(Electric).Resistances = []*PokemonType{Type(Flying), Type(Steel), Type(Electric)}
	Type(Ground).Weaknesses = []*PokemonType{Type(Water), Type(Grass), Type(Ice)}
	Type(Ground).Resistances = []*PokemonType{Type(Poison), Type(Rock)}
	Type(Ground).Immunities = []*PokemonType{Type(Electric)}
	Type(Rock).Weaknesses = []*PokemonType{Type(Fighting), Type(Ground), Type(Steel), Type(Water), Type(Grass)}
	Type(Rock).Resistances = []*PokemonType{Type(Normal), Type(Flying), Type(Poison), Type(Fire)}
	Type(Psychic).Weaknesses = []*PokemonType{Type(Bug), Type(Ghost), Type(Dark)}
	Type(Psychic).Resistances = []*PokemonType{Type(Fighting), Type(Psychic)}
	Type(Ice).Weaknesses = []*PokemonType{Type(Fighting), Type(Rock), Type(Steel), Type(Fire)}
	Type(Ice).Resistances = []*PokemonType{Type(Ice)}
	Type(Bug).Weaknesses = []*PokemonType{Type(Flying), Type(Fire), Type(Rock)}
	Type(Bug).Resistances = []*PokemonType{Type(Fighting), Type(Ground), Type(Grass)}
	Type(Ghost).Weaknesses = []*PokemonType{Type(Ghost), Type(Dark)}
	Type(Ghost).Resistances = []*PokemonType{Type(Poison), Type(Bug)}
	Type(Ghost).Immunities = []*PokemonType{Type(Fighting), Type(Normal)}
	Type(Steel).Weaknesses = []*PokemonType{Type(Fighting), Type(Ground), Type(Fire)}
	Type(Steel).Resistances = []*PokemonType{Type(Normal), Type(Flying), Type(Rock), Type(Bug), Type(Steel), Type(Grass), Type(Psychic), Type(Ice), Type(Dragon), Type(Fairy)}
	Type(Steel).Immunities = []*PokemonType{Type(Poison)}
	Type(Dragon).Weaknesses = []*PokemonType{Type(Ice), Type(Dragon), Type(Fairy)}
	Type(Dragon).Resistances = []*PokemonType{Type(Water), Type(Electric), Type(Fire), Type(Grass)}
	Type(Dark).Weaknesses = []*PokemonType{Type(Fighting), Type(Bug), Type(Fairy)}
	Type(Dark).Resistances = []*PokemonType{Type(Ghost), Type(Dark)}
	Type(Dark).Immunities = []*PokemonType{Type(Psychic)}
	Type(Fairy).Weaknesses = []*PokemonType{Type(Poison), Type(Steel)}
	Type(Fairy).Resistances = []*PokemonType{Type(Fighting), Type(Bug), Type(Dark)}
	Type(Fairy).Immunities = []*PokemonType{Type(Dragon)}
}

// Testing purpose functions/methods
func MustBuildPokemonType(pokemonType string) PokemonType {
	p := Type(pokemonType)

	return *p
}
