package highorderfuncs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterInt(t *testing.T) {
	type args struct {
		l []int
		f func(e int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "filter greater than 10",
			args: args{
				l: []int{1, 2, 3, 11, 12, 10},
				f: func(e int) bool {
					return e > 10
				},
			},
			want: []int{11, 12},
		},
		{
			name: "no filter match",
			args: args{
				l: []int{1, 2, 3, 8, 9, 10},
				f: func(e int) bool {
					return e > 10
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.args.l, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}
