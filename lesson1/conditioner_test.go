package lesson1

import "testing"

func Test_work(t *testing.T) {
	type args struct {
		room int
		cond int
		mode string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"caseHeat",
			args{10, 20, "heat"},
			20,
		},
		{
			"caseFreeze",
			args{10, 20, "freeze"},
			10,
		},
		{
			"caseAuto1",
			args{30, 40, "auto"},
			40,
		},
		{
			"caseAuto2",
			args{30, 20, "auto"},
			20,
		},
		{
			"caseFan",
			args{10, 0, "fan"},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := work(tt.args.room, tt.args.cond, tt.args.mode); got != tt.want {
				t.Errorf("work() = %v, want %v", got, tt.want)
			}
		})
	}
}
