package lesson1

import "testing"

func Test_isPassBlockThrowHole(t *testing.T) {
	type args struct {
		block []int
		hole  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"casePass",
			args{
				block: []int{1, 1, 1},
				hole: []int{1, 1},
			},
			true,
		},
		{
			"caseNotPass",
			args{
				block: []int{2, 2, 2},
				hole: []int{1, 1},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPassBlockThrowHole(tt.args.block, tt.args.hole); got != tt.want {
				t.Errorf("isPassBlockThrowHole() = %v, want %v", got, tt.want)
			}
		})
	}
}