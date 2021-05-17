package main

type Task interface {
	Question() string
	CheckAnswer(answer string) (success bool, failMessage string)
}

type Tasker interface {
	Next() (Task, bool)
}
