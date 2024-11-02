package stack

import (
	"errors"
	"fmt"
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

func minAddToMakeValid(s string) int {
    stack := []byte{}
    count := 0
    for i:=0; i < len(s); i++ {
        ch := s[i]
        if ch == '(' {
            stack = append(stack, ch)
        } else {
            // empty stack, increment count
            if len(stack) == 0 {
                count++
            } else {
                // pop the opening bracket
                stack = stack[:len(stack) - 1]
            }
        }
    }
    // total matching brackets + remainig brackets
    return count + len(stack)
}

func removeConsecutiveSame(words []string) int {
	stack := []string{}
	for _, word :=  range words {
		if (len(stack) > 0  && stack[len(stack) - 1] == word) {
			stack = stack[:len(stack) - 1] //  pop last element
		} else {
			stack = append(stack, word) //  push word
		}
	}
	return len(stack)
}

func backspaceCompare(s string, t string) bool {
    stack1 := []byte{}
    for i:=0; i < len(s); i++ {
        ch := s[i]
        if ch == '#' && len(stack1) > 0{
            stack1 = stack1[:len(stack1) - 1]
        } else if ch != '#' {
            stack1 = append(stack1 , ch)
        }
    }
    stack2 := []byte{}
    for i:=0; i < len(t); i++ {
        ch := t[i]
        if ch == '#' && len(stack2) > 0{
            stack2 = stack2[:len(stack2) - 1]
        } else if ch != '#' {
            stack2 = append(stack2 , ch)
        }
    }
    return string(stack1) == string(stack2)
}

func checkRedundantBracket(s string) bool {
	stack := []byte{}
	for i:=0; i< len(s); i++ {
		if s[i] == '(' || s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			stack = append(stack, s[i])
		} else if s[i] == ')' {
			size := len(stack)
			if stack[size - 1] == '(' {
				return true
			}
			for size > 0 && stack[size - 1] != '(' {
				stack = stack[:size - 1]
				size--
			}
			if size > 0 && stack[size - 1] == '(' {
				stack = stack[:size - 1]
			}
		}
	}
	return false
}

func generateNextGreaterElementFirst(nums []int) []int{
	result := make([]int, len(nums)) 
	for i:=0; i< len(nums); i++ {
		num := nums[i]
		j := i + 1
		for ; j < len(nums); j++ {
			if num < nums[j] {
				result[i] = nums[j]
				break
			}
		}
		if j == len(nums) {
			result[i] = -1
		}
	}
	return result
}

func generateNextGreaterElementSecond(nums []int) []int{
	stack := []int{}
	result := make([]int, len(nums)) 
	for i:=0; i < len(nums); i++ {
		nge := -1
		size := len(stack)-1
		for size >= 0 && nums[stack[size]] < nums[i] {
			fmt.Println(nums[stack[size]], nums[i])
			stack = stack[:size]
			size--
		}
		if size >= 0 {
			nge = nums[stack[size]]
		}
		stack = append(stack, i)
		
		result[i] = nge
	}
	return result
}