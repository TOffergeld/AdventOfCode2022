package Day9

import (
	"image"
	"testing"
)

func Test_pInR(t *testing.T) {
	type args struct {
		p image.Point
		r image.Rectangle
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"zero", args{image.Point{0, 0}, image.Rect(-1, -1, 1, 1)}, true},
		{"outside", args{image.Point{2, 0}, image.Rect(-1, -1, 1, 1)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pInR(tt.args.p, tt.args.r); got != tt.want {
				t.Errorf("pInR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sign(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{
			-1,
		}, want: -1},
		{
			"",
			args{
				v: 1,
			},
			1,
		},
		{
			name: "",
			args: args{
				v: -42,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sign(tt.args.v); got != tt.want {
				t.Errorf("sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay9(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"default"},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Day9()
		})
	}
}