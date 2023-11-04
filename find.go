package highorderfuncs

func Find[E any](l []E, f func(e E) bool) (find E) {
	for _, ll := range l {
		if f(ll) {
			return ll
		}
	}

	return
}
