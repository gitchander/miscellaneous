package optional

type OptInt64 struct {
	Present bool
	Value   int64
}

func PresentInt64(a int64) OptInt64 {
	return OptInt64{
		Present: true,
		Value:   a,
	}
}

func (p *OptInt64) Reset() {
	*p = OptInt64{}
}
