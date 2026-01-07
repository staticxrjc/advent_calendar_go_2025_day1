package dial

import (
	"fmt"
	"strconv"
)

type Decoder struct {
	Dial *Dial
}

func NewDecoder(dial *Dial) *Decoder {
	return &Decoder{
		Dial: dial,
	}
}

func (d *Decoder) TurnDial(input string) error {
	firstChar := input[0]
	val, err := strconv.Atoi(input[1:])
	if err != nil {
		return fmt.Errorf("Input needs to be in format L(INT)\n")
	}
	if firstChar == 'L' {
		d.Dial.TurnLeft(val)
	} else if firstChar == 'R' {
		d.Dial.TurnRight(val)
	}

	return nil
}

func (d *Decoder) GetPassword() int {
	return d.Dial.GetPassword()
}
