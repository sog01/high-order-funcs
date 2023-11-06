package highorderfuncs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapToStruct(t *testing.T) {
	type D struct {
		Id   int
		Name string
	}

	type args struct {
		l []string
		f func(i int, e string) D
	}
	tests := []struct {
		name string
		args args
		want []D
	}{
		{
			name: "Map to struct",
			args: args{
				l: []string{"john", "doe"},
				f: func(id int, name string) D {
					return D{
						Id:   id,
						Name: name,
					}
				},
			},
			want: []D{
				{
					Id:   0,
					Name: "john",
				},
				{
					Id:   1,
					Name: "doe",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(tt.args.l, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}
