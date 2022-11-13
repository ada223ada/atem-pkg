package generic

type List[T comparable] struct {
	l []T
}

func NewList[T comparable](t []T) *List[T] {
	return &List[T]{l: t}
}

func (l *List[T]) Push(v T) {
	l.l = append(l.l, v)
}

func (l *List[T]) Insert(v T) *List[T] {
	l.InsertAt(0, v)
	return l
}

func (l *List[T]) InsertAt(pos int, v T) *List[T] {
	l.l = append(l.l[:pos+1], l.l[pos:]...)
	l.l[0] = v
	return l
}

func (l *List[T]) Remove(i int) *List[T] {
	l.l = l.l[:i+copy(l.l[i:], l.l[i+1:])]
	return l
}

func (l *List[T]) Equals(rhs *List[T]) bool {
	if len(l.l) != len(rhs.l) {
		return false
	}
	for i := 0; i < len(l.l); i++ {
		if l.l[i] != rhs.l[i] {
			return false
		}
	}
	return true
}

func (l *List[T]) Clone() *List[T] {
	ll := &List[T]{l: make([]T, len(l.l))}
	copy(ll.l, l.l)
	return ll
}

func (l *List[T]) Slice() []T {
	return l.l
}
