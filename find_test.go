package highorderfuncs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindInt(t *testing.T) {
	type args struct {
		l []int
		f func(e int) bool
	}
	tests := []struct {
		name     string
		args     args
		wantFind int
	}{
		{
			name: "find number 10",
			args: args{
				l: []int{1, 2, 3, 10},
				f: func(e int) bool {
					return e == 10
				},
			},
			wantFind: 10,
		},
		{
			name: "no find match",
			args: args{
				l: []int{1, 2, 3, 11},
				f: func(e int) bool {
					return e == 10
				},
			},
			wantFind: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFind := Find(tt.args.l, tt.args.f)
			assert.Equal(t, tt.wantFind, gotFind)
		})
	}
}
