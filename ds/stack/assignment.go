package stack

import (
	"errors"
	_ "fmt"
	"go-program/ds/utill"
)

// implement stack in golang
type Stack[T any] struct {
	Push func(T)
	Pop func()(T, error)
	IsEmpty func() bool
	Top func() (T, error)
	Len func() int
}

func CreateStack[T any]() *Stack[T] {
	elements := make([]T, 0)
	return &Stack[T]{
		Push: func(t T) {
			elements = append(elements, t)
		},
		Pop: func() (T, error) {
			var t T
			if len(elements) == 0 {
				return t, errors.New(utill.EMPTY_STACK)
			}
			n := len(elements)
			t = elements[n-1]
			elements = elements[:n-1]
			return t, nil
		},
		IsEmpty: func() bool {
			return len(elements) == 0
		},
		Top: func() (T, error) {
			var t T
			if len(elements) == 0 {
				return t, errors.New(utill.EMPTY_STACK)
			}
			n := len(elements)
			t = elements[n - 1]
			return t, nil
		},
		Len: func() int {
			return len(elements)
		},
	}
}

func findMiddleOfStack(stack *[]int, mid int) int {
	old := *stack
	// base case
	if mid == 0 {
		return old[len(old)-1]	
	}
	// 1 case
	size := len(old)
	top := old[size-1] // pull top elm
	*stack = old[:size-1]
	m := findMiddleOfStack(stack, mid-1) // reduce mid and call find middle
	*stack = append(*stack, top) // back track
	return m
}

func checkSortedRec(stack *[]int, element int) bool {
	old := *stack
	size := len(old)
	// base case
	if size == 0  || size == 1{
		return true
	}
	// 1 case 
	top := old[size-1] // pull top elm
	*stack = old[:size-1]
	if top > element {
		return false
	}

	recAns := checkSortedRec(stack, top)
	// rec
	return recAns
}

func checkSortedIterative(stack []int) bool {
	size := len(stack)
	if size <= 1{
		return true
	}
	top := stack[size-1]
	for size > 0 {
		next := stack[size -1]
		if top < next {
			return false
		}
		top = next
		size--
	}
	return true
}

func insertBottom(stack *[]int, element int) {
	// empty stack insert
	if len(*stack) == 0 {
		*stack = append(*stack, element)
		return
	}
	old := *stack
	top := old[len(old) - 1] // pull top
	*stack = old[:len(old) - 1]
	insertBottom(stack, element) // insert element recursively
	*stack = append(*stack, top) // now put top back in stack
}

func reverseStack(stack *[]int) {
	//base case
	if len(*stack) == 0 {
		return
	}
	old := *stack
	//1 case
	top := old[len(old) - 1] // pull top element
	*stack = old[:len(old) - 1]

	reverseStack(stack) // recursive call for rest of element

	insertBottom(stack, top) // insert element at the end

	// rec
}