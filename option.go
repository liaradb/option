package option

type Option[T any] struct {
	t *T
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func Some[T any](t T) Option[T] {
	return Option[T]{
		t: &t,
	}
}

func (o Option[T]) Value() (T, bool) {
	if o.t == nil {
		var t T
		return t, false
	}

	return *o.t, true
}

func (o Option[T]) Pointer() *T {
	return o.t
}

func (o Option[T]) Unsafe() T {
	return *o.t
}
