package core

import (
	"reflect"
	"testing"
)

func Test_reverseInts(t *testing.T) {
	type args struct {
		src [Size]int
	}
	tests := []struct {
		name string
		args args
		want [Size]int
	}{
		{
			"Reverse [0, 0, 0, 0]",
			args{src: [Size]int{0, 0, 0, 0}},
			[Size]int{0, 0, 0, 0},
		},
		{
			"Reverse [2, 0, 0, 0]",
			args{src: [Size]int{2, 0, 0, 0}},
			[Size]int{0, 0, 0, 2},
		},
		{
			"Reverse [0, 0, 2, 0]",
			args{src: [Size]int{0, 0, 2, 0}},
			[Size]int{0, 2, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseInts(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
