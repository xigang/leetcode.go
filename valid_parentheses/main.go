package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(isValid("[]{}()"))
}

func isValid(s string) bool {
	hashMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	stack := NewStack()
	for _, v := range s {
		c := string(v)
		switch c {
		case "(", "[", "{":
			stack.Push(c)
			break
		case ")", "]", "}":
			if stack.Empty() || hashMap[c] != stack.Pop() {
				return false
			}
		}
	}
	return stack.Empty()
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
