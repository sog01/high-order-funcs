package highorderfuncs

func Reduce[E any, D any](l []E, f func(res D, currentValue E) D, initialValue D) D {
	res := initialValue
	for _, ll := range l {
		res = f(res, ll)
	}
	return res
}
