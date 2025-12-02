package dial

type Decoder struct {
	Dial *Dial
}

func NewDecoder(dial *Dial) *Decoder {
	return &Decoder{
		Dial: dial,
	}
}

func (d *Decoder) TurnDial(input string) {

}

func (d *Decoder) GetPassword() int {
	return d.Dial.GetPassword()
}
