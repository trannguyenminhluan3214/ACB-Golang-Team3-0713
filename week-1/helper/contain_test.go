package helper

import "testing"

func TestContains(t *testing.T) {
	type args struct {
		list interface{}
		v    interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{
				list: []int{1, 2, 3, 4, 5, 7, 6},
				v:    2,
			},
			want: true,
		},
		{
			name: "Case 1",
			args: args{
				list: []int{1, 2, 3, 4, 5, 7, 6},
				v:    10,
			},
			want: false,
		},
		{
			name: "Case 1",
			args: args{
				list: "html",
				v:    "l",
			},
			want: true,
		},
		{
			name: "Case 1",
			args: args{
				list: "html",
				v:    "a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.list, tt.args.v); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
