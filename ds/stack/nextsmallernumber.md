# Question : [Generate next smaller element - (practice)]()

Given a array of nums generate the next smaller elements for elements

### Example 1

```
Input: array = [6, 5, 8, 3, 4, 2, 1, 4, 7]
Output: [5, 3, 3, 2, 2, 1, -1, -1, -1]

```

### Example 2

```
Input: array = [1, 2, 3, 4]
Output: [-1, -1, -1, -1]

```
### Example 3

```
Input: array = [4, 3, 2, 1]
Output: [3, 2, 1, -1]

```

**Expected Time Complexity**: O(N)
**Expected Auxiliary Space**: O(N)

### Constraints

-   1<=|str|<=10^4

<!-- **Follow up:** Can you solve it in *O(n)* time and *O(1)* space? -->

## Solution

```GO
package main
import "fmt"

type Stack[T any] struct {
	Push    func(T)
	Pop     func() (T, error)
	IsEmpty func() bool
	Top     func() (T, error)
	Len     func() int
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
			t = elements[n-1]
			return t, nil
		},
		Len: func() int {
			return len(elements)
		},
	}
}

func nextSmallerElements(nums []int) []int {
	result := make([]int, len(nums))
	stack := CreateStack[int]()
	for i := len(nums) - 1; i >= 0; i-- {
		nse := -1
		for !stack.IsEmpty() {
			top, _ := stack.Top()
			if nums[top] < nums[i] {
				break
			}
			stack.Pop()
		}
		if !stack.IsEmpty() {
			top, _ := stack.Top()
			nse = nums[top]
		}
		stack.Push(i)
		result[i] = nse
	}
	return result
}

func main() {
    fmt.Println(nextSmallerElements([]int{1, 2, 3, 4})) //  print : -1, -1, -1, -1
    fmt.Println(nextSmallerElements([]int{6, 5, 8, 3, 4, 2, 1, 4, 7})) //  print : 5, 3, 3, 2, 2, 1, -1, -1, -1
    fmt.Println(nextSmallerElements([]int{4, 3, 2, 1})) //  print : 3, 2, 1, -1
}

Time Complexity = O(n)
Space Complexity = O(n)
=> where n is the lengths of the array, respectively.
```