package newmath

import "testing"

func TestSqrt(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"整数4", args{4}, 2},
		{"整数9", args{9}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sqrt(tt.args.x); got != tt.want {
				t.Errorf("Sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
