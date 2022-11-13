package generic

type Ordered interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

func Max[T Ordered](a []T) T {
	m := a[0]
	for _, v := range a {
		if m < v {
			m = v
		}
	}
	return m
}

func Min[T Ordered](a []T) T {
	m := a[0]
	for _, v := range a {
		if m >= v {
			m = v
		}
	}
	return m
}