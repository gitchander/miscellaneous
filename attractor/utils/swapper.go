package utils

type Swapper interface {
	Len() int
	Swap(i, j int)
}

// a b c d e f
//  / / / / /
// b c d e f a
func RotateDown(s Swapper) {
	n := s.Len()
	for i := 0; i < n-1; i++ {
		s.Swap(i, i+1)
	}
}

// a b c d e f
//  \ \ \ \ \
// f a b c d e
func RotateUp(s Swapper) {
	n := s.Len()
	for i := n - 1; i > 0; i-- {
		s.Swap(i, i-1)
	}
}
