package lesson1

import "testing"

func Test_findingSolution(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"caseNoSolution",
			args{1, 2, -3},
			"NO SOLUTION",
		},
		{
			"caseManySolution",
			args{0, 4, 2},
			"MANY SOLUTIONS",
		},
		{
			"case1",
			args{1, 0, 0},
			"0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findingSolution(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("findingSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
