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

func TestMakeValidParentheses(t *testing.T) {
	assert.Equal(t, 3, minAddToMakeValid("(((()"))
	assert.Equal(t, 4, minAddToMakeValid("(((("))
	assert.Equal(t, 0, minAddToMakeValid(""))
}

func TestConsicutiveSame(t *testing.T) {
	words := []string{"tom", "jerry", "jerry", "tom"}
	assert.Equal(t, 0, removeConsecutiveSame(words))
	words = []string{"ab", "aa", "aa", "bcd", "ab"}
	assert.Equal(t, 3, removeConsecutiveSame(words))
}

func TestBackSpaceStringCompare(t *testing.T) {
	assert.Equal(t, false, backspaceCompare("ab##", "ad#c"))
	assert.Equal(t, true, backspaceCompare("ab##", "c#d#"))
	assert.Equal(t, false, backspaceCompare("a#c", "b"))
}

func TestCheckRedundantBracket(t *testing.T) {
	assert.Equal(t, false, checkRedundantBracket("(a+b)"))
	assert.Equal(t, true, checkRedundantBracket("((a+b))"))
	assert.Equal(t, false, checkRedundantBracket("c+d+(a+b)"))
	assert.Equal(t, false, checkRedundantBracket("(c+d+(a+b)"))
}

func TestGenerateNextGreaterElementFirst(t *testing.T) {
	assert.Equal(t, []int{2, 3, 4, -1}, generateNextGreaterElementFirst([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{8, 8, -1, 4, 7, 4, 4, 7, -1}, generateNextGreaterElementFirst([]int{6, 5, 8, 3, 4, 2, 1, 4, 7}))
	assert.Equal(t, []int{-1, -1, -1, -1}, generateNextGreaterElementFirst([]int{4, 3, 2, 1}))
}

func TestGenerateNextGreaterElementSecond(t *testing.T) {
	assert.Equal(t, []int{2, 3, 4, -1}, generateNextGreaterElementSecond([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{8, 8, -1, 4, 7, 4, 4, 7, -1}, generateNextGreaterElementSecond([]int{6, 5, 8, 3, 4, 2, 1, 4, 7}))
	assert.Equal(t, []int{-1, -1, -1, -1}, generateNextGreaterElementSecond([]int{4, 3, 2, 1}))
}

func TestGenerateNextGreaterElement2Second(t *testing.T) {
	assert.Equal(t, []int{2,-1,2}, nextGreaterElements([]int{1,2,1}))
	assert.Equal(t, []int{2,3,4,-1,4}, nextGreaterElements([]int{1,2,3,4,3}))
	assert.Equal(t, []int{-1, 4, 4, 4}, nextGreaterElements([]int{4, 3, 2, 1}))
}