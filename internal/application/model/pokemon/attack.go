package pokemon

type Move struct {
	id         int
	name       string
	attackType *Type
}

func (a *Move) Id() int {
	return a.id
}

func (a *Move) Name() string {
	return a.name
}

func (a *Move) MoveType() *Type {
	return a.attackType
}

func NewMove(
	id int,
	name string,
	attackType *Type,
) *Move {
	return &Move{
		id:         id,
		name:       name,
		attackType: attackType,
	}
}
