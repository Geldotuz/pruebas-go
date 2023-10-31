package players

type Player struct {
	Name     string
	LastName string
	Age      int64
}

func NewPlayer(name, LastName string, age int64) *Player {
	return &Player{
		Name:     name,
		LastName: LastName,
		Age:      age,
	}
}

func (p *Player) Cumpleaños() {
	p.Age++
}

func (p *Player) CambioApellido(apellido string) {
	p.LastName = (apellido)
}
