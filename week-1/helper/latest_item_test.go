package helper

import (
	"reflect"
	"testing"
)

func TestLast(t *testing.T) {
	type args struct {
		lst interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Case 1",
			args: args{
				lst: []int{1, 2, 3, 4, 5, 7, 6},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Last(tt.args.lst); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Last() = %v, want %v", got, tt.want)
			}
		})
	}
}
