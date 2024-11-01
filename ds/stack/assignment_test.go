package stack

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateStack(t *testing.T) {
	stack := CreateStack[int]()
	_, err := stack.Pop()
	assert.NotNil(t, err)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	last, _ := stack.Pop()
	assert.Equal(t, 3, last)
	second, _ := stack.Pop()
	assert.Equal(t, 2, second)
	first, _ := stack.Pop()
	assert.Equal(t, 1, first)
}

func TestStackMiddle(t *testing.T) {
	stack := []int{}
	stack = append(stack, 2)
	stack = append(stack, 4)
	stack = append(stack, 7)
	stack = append(stack, 8)
	stack = append(stack, 1)

	mid := len(stack) / 2

	midElement := findMiddleOfStack(&stack, mid)
	assert.Equal(t, 7, midElement)
}

func TestCheckSorted(t *testing.T) {
	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 4)
	stack = append(stack, 7)
	stack = append(stack, 8)
	stack = append(stack, 10)

	assert.Equal(t, true, checkSortedIterative(stack))
	assert.Equal(t, true, checkSortedRec(&stack, math.MaxInt))
}

func TestReverseStack(t *testing.T) {
	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 4)
	stack = append(stack, 7)
	stack = append(stack, 8)
	stack = append(stack, 10)
	reverseStack(&stack)
	
	assert.ElementsMatch(t, stack, []int{10, 8, 7, 4, 1})
}
