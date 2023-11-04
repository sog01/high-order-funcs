package highorderfuncs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduceInt(t *testing.T) {
	type args struct {
		l            []int
		f            func(res, currentValue int) int
		initialValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sum",
			args: args{
				l: []int{1, 2, 3, 4, 5},
				f: func(res, currentValue int) int {
					return res + currentValue
				},
				initialValue: 0,
			},
			want: 15,
		},
		{
			name: "sum initial 10",
			args: args{
				l: []int{1, 2, 3, 4, 5},
				f: func(res, currentValue int) int {
					return res + currentValue
				},
				initialValue: 10,
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.args.l, tt.args.f, tt.args.initialValue)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReduceToStringMap(t *testing.T) {
	type args struct {
		l            []string
		f            func(res map[string]struct{}, currentValue string) map[string]struct{}
		initialValue map[string]struct{}
	}
	tests := []struct {
		name string
		args args
		want map[string]struct{}
	}{
		{
			name: "sum",
			args: args{
				l: []string{"john", "doe"},
				f: func(res map[string]struct{}, currentValue string) map[string]struct{} {
					res[currentValue] = struct{}{}
					return res
				},
				initialValue: map[string]struct{}{},
			},
			want: map[string]struct{}{
				"john": {},
				"doe":  {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.args.l, tt.args.f, tt.args.initialValue)
			assert.Equal(t, tt.want, got)
		})
	}
}
