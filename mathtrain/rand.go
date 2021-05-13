package main

import (
	"math/rand"
	"time"
)

func newRandSeed(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

func newRandTime(t time.Time) *rand.Rand {
	return newRandSeed(t.UnixNano())
}

func newRandNow() *rand.Rand {
	return newRandTime(time.Now())
}

func randIntByInterval(r *rand.Rand, v Interval) int {
	if v.Empty() {
		panic("interval is empty")
	}
	return v.Min + r.Intn(v.Max-v.Min)
}

func randArgumentsAdd(r *rand.Rand) (x, y int) {

	var v = Interval{
		Min: 0,
		Max: 100,
	}

	x = randIntByInterval(r, v)
	// y = randIntByInterval(r, v)
	y = randIntByInterval(r,
		Interval{
			Min: 0,
			Max: (v.Max - x),
		})

	return x, y
}

func randArgumentsSub(r *rand.Rand) (x, y int) {

	var v = Interval{
		Min: 0,
		Max: 100,
	}

	y = randIntByInterval(r, v)
	//z := randIntByInterval(r, v) // z = x - y
	z := randIntByInterval(r,
		Interval{
			Min: 0,
			Max: (v.Max - y),
		})

	x = z + y

	return x, y
}

func randArgumentsMul(r *rand.Rand) (x, y int) {

	var v = Interval{
		Min: 2,
		Max: 10,
	}

	x = randIntByInterval(r, v)
	y = randIntByInterval(r, v)

	return x, y
}

func randArgumentsDiv(r *rand.Rand) (x, y int) {

	var v = Interval{
		Min: 2,
		Max: 10,
	}

	y = randIntByInterval(r, v)
	z := randIntByInterval(r, v) // z = x / y

	x = z * y

	return x, y
}

func randArguments(r *rand.Rand, bo BinaryOperation) (x, y int) {
	switch bo.(type) {
	case Add:
		return randArgumentsAdd(r)
	case Sub:
		return randArgumentsSub(r)
	case Mul:
		return randArgumentsMul(r)
	case Div:
		return randArgumentsDiv(r)
	}
	return
}

func randBinaryOperation(r *rand.Rand, bos []BinaryOperation) BinaryOperation {
	switch n := len(bos); n {
	case 0:
		return nil
	case 1:
		return bos[0]
	default:
		return bos[r.Intn(n)]
	}
}

var correctPhrases = []string{
	"Correct",
	"Right",
	"Bingo!",
	"Very well!",
	"Excellent!",
	"Well done!",
	"Splendid!",
	"That's famous!",
}

var wrongPhrases = []string{
	"Wrong",
	"Incorrectly",
	"Mistaken",
}

func randString(r *rand.Rand, corpus []string) string {
	return corpus[r.Intn(len(corpus))]
}

func randCorrectPhrase(r *rand.Rand) string {
	return randString(r, correctPhrases)
}

func randWrongPhrase(r *rand.Rand) string {
	return randString(r, wrongPhrases)
}
