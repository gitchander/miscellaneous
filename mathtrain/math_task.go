package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type mathTask struct {
	bo      BinaryOperation
	x, y, z int // x @ y = z, @ - binary operation
}

var _ Task = &mathTask{}

func randMathTask(r *rand.Rand, bos []BinaryOperation) *mathTask {

	bo := randBinaryOperation(r, bos)
	x, y := randArguments(r, bo)

	z := bo.Do(x, y)

	return &mathTask{
		bo: bo,
		x:  x,
		y:  y,
		z:  z,
	}
}

func (p *mathTask) Question() string {
	return fmt.Sprintf("%d %s %d = ", p.x, p.bo.Symbol(), p.y)
}

func (p *mathTask) CheckAnswer(answer string) (success bool, failMessage string) {

	answer = strings.TrimSpace(answer)

	result, err := strconv.Atoi(answer)
	if err != nil {
		failMessage = fmt.Sprintf("Invalid input data: %q", answer)
		return false, failMessage
	}

	if result != p.z {
		failMessage = fmt.Sprintf("Wrong result: have %d, want %d", result, p.z)
		return false, failMessage
	}

	return true, ""
}

// ------------------------------------------------------------------------------
type mathTasker struct {
	r   *rand.Rand
	bos []BinaryOperation
}

var _ Tasker = &mathTasker{}

func newMathTasker(r *rand.Rand, bos []BinaryOperation) *mathTasker {
	return &mathTasker{
		r:   r,
		bos: bos,
	}
}

func (p *mathTasker) Next() (Task, bool) {
	task := randMathTask(p.r, p.bos)
	return task, true
}
