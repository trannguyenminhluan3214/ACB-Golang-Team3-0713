package helper

import (
	"reflect"
	"testing"
)

func TestMax(t *testing.T) {
	type args struct {
		list interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Case 1",
			args: args{
				list: []int{1, 2, 3, 4, 5, 7, 6},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
