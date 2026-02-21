package entity

type player struct {
	name   string
	symbol string
}

type Player interface {
	GetName() string
	GetSymbol() string
}

func NewPlayer(name, symbol string) *player {
	return &player{
		name:   name,
		symbol: symbol,
	}
}

func (p *player) GetName() string {
	return p.name
}

func (p *player) GetSymbol() string {
	return p.symbol
}
