package dial

import (
	"reflect"
	"testing"
)

func TestNewDial(t *testing.T) {
	type args struct {
		initialPosition int
	}
	tests := []struct {
		name string
		args args
		want *Dial
	}{
		{
			name: "Returns dial at position 0",
			args: args{0},
			want: &Dial{Position: 0},
		},
		{
			name: "Returns dial at position 50",
			args: args{50},
			want: &Dial{Position: 50},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDial(tt.args.initialPosition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDial_TurnRight(t *testing.T) {
	type fields struct {
		Position int
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "Turn right landing on 0",
			fields: fields{Position: 99},
			args:   args{1},
			want:   0,
		},
		{
			name:   "Turn right rolling over",
			fields: fields{Position: 95},
			args:   args{6},
			want:   1,
		},
		{
			name:   "Turn right yields expected value",
			fields: fields{Position: 0},
			args:   args{1},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Dial{
				Position: tt.fields.Position,
			}
			d.TurnRight(tt.args.value)
			if got := d.Position; got != tt.want {
				t.Errorf("Dial position = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDial_TurnLeft(t *testing.T) {
	type fields struct {
		Position int
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Turn left landing on 0",
			fields: fields{
				Position: 2,
			},
			args: args{2},
			want: 0,
		},
		{
			name: "Turning left rolling over",
			fields: fields{
				Position: 2,
			},
			args: args{3},
			want: 99,
		},
		{
			name: "Turn left returns expected value",
			fields: fields{
				Position: 2,
			},
			args: args{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dial{
				Position: tt.fields.Position,
			}
			d.TurnLeft(tt.args.value)
			if got := d.Position; got != tt.want {
				t.Errorf("Dial position = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDial_GetPassword(t *testing.T) {
	type fields struct {
		Position int
	}
	tests := []struct {
		name      string
		fields    fields
		direction []int
		want      int
	}{
		{
			name:      "Initializes at 0",
			fields:    fields{50},
			direction: []int{0},
			want:      0,
		},
		{
			name:      "Left turn leads to 1 hit",
			fields:    fields{50},
			direction: []int{-50},
			want:      1,
		},
		{
			name:      "Right turn leads to 1 hit",
			fields:    fields{50},
			direction: []int{50},
			want:      1,
		},
		{
			name:      "Complex turns",
			fields:    fields{50},
			direction: []int{50, 33, -33, 0, 50, 50},
			want:      3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dial{
				Position: tt.fields.Position,
			}
			for _, val := range tt.direction {
				if val < 0 {
					d.TurnLeft(-val)
				} else {
					d.TurnRight(val)
				}
			}
			if got := d.GetPassword(); got != tt.want {
				t.Errorf("GetPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
