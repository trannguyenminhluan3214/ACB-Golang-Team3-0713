package helper

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	type args struct {
		t interface{}
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test",
			args: args{
				t: "",
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				t: "abc",
			},
			want: false,
		},
		{
			name: "Test 3",
			args: args{
				t: 0,
			},
			want: true,
		},
		{
			name: "Test 4",
			args: args{
				t: nil,
			},
			want: true,
		},
		{
			name: "Test 5",
			args: args{
				t: []int{},
			},
			want: true,
		},
		{
			name: "Test 6",
			args: args{
				t: []int{1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.t); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
