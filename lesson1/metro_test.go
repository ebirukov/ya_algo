package lesson1

import (
	"reflect"
	"testing"
)

func Test_timesOfWatch(t *testing.T) {
	type args struct {
		a int
		b int
		n int
		m int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			"valid interval",
			args{1,3,3,2},
			[]int{5,7},
			false,
		},
		{
			"valid interval equal",
			args{3,3,1,1},
			[]int{1,7},
			false,
		},
		{
			"valid interval edge",
			args{1,1,4,3},
			[]int{7,7},
			false,
		},
		{
			"valid interval one",
			args{1,1,1,1},
			[]int{1,3},
			false,
		},
		{
			"valid interval in",
			args{11,1,1,10},
			[]int{19,21},
			false,
		},
		{
			"invalid interval",
			args{1,5,1,2},
			nil,
			true,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := timesOfWatch(tt.args.a, tt.args.b, tt.args.n, tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("timesOfWatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("timesOfWatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}
