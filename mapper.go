package highorderfuncs

func Map[E any, D any](l []E, f func(i int, e E) D) []D {
	var res []D
	for i, ll := range l {
		res = append(res, f(i, ll))
	}
	return res
}
