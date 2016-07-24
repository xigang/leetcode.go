package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(isValid("[]{}()"))
}

func isValid(s string) bool {
	stack := NewStack()
	for _, v := range s {
		stack.Push(string(v))
	}

	stackLen := stack.Len()
	var valid bool
	if stackLen > 0 {

	}

}

type Stack struct {
	l *list.List
}

func NewStack() *Stack {
	stack := list.New()
	return &Stack{l: stack}
}

func (stack *Stack) Push(value interface{}) {
	stack.l.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.l.Back()
	if e != nil {
		stack.l.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{} {
	e := stack.l.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

func (stack *Stack) Len() int {
	return stack.l.Len()
}

func (stack *Stack) Empty() bool {
	return stack.l.Len() == 0
}
