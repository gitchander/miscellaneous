package random

func RandByCorpus[T any](r *Rand, vs []T) T {
	return vs[r.Intn(len(vs))]
}
