package utils

func FindFirstUsing[T any](ss []T, fn func(value T) bool) *T {
	for _, value := range ss {
		if fn(value) {
			return &value
		}
	}
	return nil
}

func Reduce[T any, R any](ss []T, reducer func(R, T) R, initialValue ...R) (el R) {
	if len(ss) == 0 {
		return
	}
	if len(initialValue) > 0 {
		el = initialValue[0]
	}
	for _, s := range ss {
		el = reducer(el, s)
	}
	return
}
