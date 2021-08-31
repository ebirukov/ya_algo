package lesson1

import (
	"fmt"
	"testing"
)

func Test_calcNumOfParts(t *testing.T) {
	type args struct {
		n int
		k int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{10, 5, 2},
			4,
		},
		{
			"case3",
			args{13, 5, 3},
			3,
		},
		{
			"case2",
			args{14, 5, 3},
			4,
		},
		{
			"case4",
			args{3, 5, 14},
			0,
		},

		{
			"case5",
			args{0, 0, 0},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcNumOfParts(tt.args.n, tt.args.k, tt.args.m); got != tt.want {
				t.Errorf("calcNumOfParts() = %v, want %v", got, tt.want)
			}
		})
	}
}

var res int

func BenchmarkCalcNumOfParts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//calcNumOfParts(199, 117, 13)
		res = calcNumOfParts(200, 199, 200)
		//calcNumOfParts(199, 99, 49)
	}
	fmt.Printf("%v", res)
}
