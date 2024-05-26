package random

type Swapper interface {
	Len() int
	Swap(i, j int)
}

func Shuffle(r *Rand, sw Swapper) {
	for n := sw.Len(); n > 1; n-- {
		sw.Swap(r.Intn(n), n-1)
	}
}
