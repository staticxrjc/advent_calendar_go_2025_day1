package dial

import (
	"reflect"
	"testing"
)

func TestDecoder_GetPassword(t *testing.T) {
	type fields struct {
		Dial *Dial
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Decoder{
				Dial: tt.fields.Dial,
			}
			if got := d.GetPassword(); got != tt.want {
				t.Errorf("GetPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_TurnDial(t *testing.T) {
	type fields struct {
		Dial *Dial
	}
	type args struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Decoder{
				Dial: tt.fields.Dial,
			}
			d.TurnDial(tt.args.input)
		})
	}
}

func TestNewDecoder(t *testing.T) {
	type args struct {
		dial *Dial
	}
	tests := []struct {
		name string
		args args
		want *Decoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDecoder(tt.args.dial); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}
