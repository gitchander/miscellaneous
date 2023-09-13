package optional

type Optional[T any] struct {
	Present bool
	Value   T
}

func MakePresent[T any](v T) Optional[T] {
	return Optional[T]{
		Present: true,
		Value:   v,
	}
}

func (p *Optional[T]) Reset() {
	*p = Optional[T]{}
}

func (p *Optional[T]) SetValue(v T) {
	*p = MakePresent(v)
}

func (o Optional[T]) GetValue() (T, bool) {
	if o.Present {
		return o.Value, true
	}
	var zv T // zero value
	return zv, false
}

func (o Optional[T]) If(f func(T)) {
	if o.Present {
		f(o.Value)
	}
}
