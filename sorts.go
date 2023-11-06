package highorderfuncs

import "sort"

func Sort[E any](l []E, f func(before, after E) bool) []E {
	sort.SliceStable(l, func(i, j int) bool {
		return f(l[i], l[j])
	})
	return l
}
