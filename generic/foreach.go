package generic

func ForEach[T any](a []T, f func(T) bool) {
	for _, e := range a {
		if !f(e) {
			break
		}
	}
}
