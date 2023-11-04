package highorderfuncs

func Filter[E any](l []E, f func(e E) bool) []E {
	var res []E
	for _, ll := range l {
		if f(ll) {
			res = append(res, ll)
		}
	}
	return res
}
