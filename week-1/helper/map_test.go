package helper

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		list         interface{}
		iterateeFunc interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "case 1",
			args: args{
				list: []int{1, 2, 3},
				iterateeFunc: func(elem int) int {
					return elem * 3
				},
			},
			want: []int{3, 6, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.list, tt.args.iterateeFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
