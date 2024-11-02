# Question : [Generate next greater element - (GFG)](https://www.geeksforgeeks.org/problems/expression-contains-redundant-bracket-or-not/1)

Given a array of nums generate the next greater elements for elements

### Example 1

```
Input: array = [6, 5, 8, 3, 4, 2, 1, 4, 7]
Output: [8, 8, -1, 4, 7, 4, 4, 7, -1]

```

### Example 2

```
Input: array = [1, 2, 3, 4]
Output: [2, 3, 4, -1]

```
### Example 3

```
Input: array = [4, 3, 2, 1]
Output: [-1, -1, -1, -1]

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

// Function to initialize the recursion with starting indices of 0
func main() {
    fmt.Println(checkRedundantBracket("(a+b)")) //  print : false
    fmt.Println(checkRedundantBracket("((a+b))")) //  print : true
    fmt.Println(checkRedundantBracket("c+d+(a+b)")) //  print : false
    fmt.Println(checkRedundantBracket("(c+d+(a+b)")) //  print : false
}

Time Complexity = O(n)
Space Complexity = O(n)
=> where n is the lengths of the array, respectively.
```