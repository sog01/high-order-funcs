package highorderfuncs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortsStruct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	type args struct {
		l []User
		f func(before, after User) bool
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{
			name: "sort age ascending",
			args: args{
				l: []User{
					{
						Name: "John",
						Age:  20,
					},
					{
						Name: "Doe",
						Age:  19,
					},
					{
						Name: "Rocky",
						Age:  18,
					},
				},
				f: func(before, after User) bool {
					return before.Age < after.Age
				},
			},
			want: []User{
				{
					Name: "Rocky",
					Age:  18,
				},
				{
					Name: "Doe",
					Age:  19,
				},
				{
					Name: "John",
					Age:  20,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sort(tt.args.l, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}
