package lesson1

import (
	"reflect"
	"testing"
)

func Test_solveSysEquations(t *testing.T) {
	type args struct {
		a float32
		b float32
		c float32
		d float32
		e float32
		f float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			"caseNoSolution",
			args{1, 1, 1, 1, -300, 300},
			[]float32{0},
		},
		{
			"caseNoSolution2",
			args{0, 1, 0, 1, -300, 300},
			[]float32{0},
		},
		{
			"caseNoSolution3",
			args{1, 0, 1, 0, -300, 300},
			[]float32{0},
		},
		{
			"simple_inf_y0_any_x",
			args{0, 0, 0, 1, -300, 300},
			[]float32{4, 300},
		},
		{
			"simple_inf_x0_any_y",
			args{0, 0, 1, 0, 300, -300},
			[]float32{3, -300},
		},
		{
			"simple2_inf_x0_any_y",
			args{1, 0, 0, 0, 300, -300},
			[]float32{3, 300},
		},
		{
			"case1",
			args{1, 0, 0, 1, 3, 3},
			[]float32{2, 3, 3},
		},
		{
			"inf_kx+b",
			args{1, 1, 2, 2, 1, 2},
			[]float32{1, -1, 1},
		},
		{
			"inf_y0_any_x",
			args{0, 2, 0, 4, 1, 2},
			[]float32{4, 0.5},
		},
		{
			"inf_x0_any_y",
			args{2, 0, 4, 0, 1, 2},
			[]float32{3, 0.5},
		},
		{
			"inf_all",
			args{0, 0, 0, 0, 3, 3},
			[]float32{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveSysEquations(tt.args.a, tt.args.b, tt.args.c, tt.args.d, tt.args.e, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveSysEquations() = %v, want %v", got, tt.want)
			}
		})
	}
}
