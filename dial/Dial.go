package dial

type Dial struct {
	Position int
	NumHits  int
}

func NewDial(initialPosition int) *Dial {
	return &Dial{
		initialPosition,
		0,
	}
}

func (d *Dial) TurnRight(value int) {
	d.Position = (value + d.Position) % 100
	if d.Position == 0 && value != 0 {
		d.NumHits += 1
	}
}

func (d *Dial) TurnLeft(value int) {
	if d.Position < value {
		d.Position = 100 + d.Position - value
	} else {
		d.Position = d.Position - value
	}
	if d.Position == 0 && value != 0 {
		d.NumHits += 1
	}
}

func (d *Dial) GetPassword() int {
	return d.NumHits
}
