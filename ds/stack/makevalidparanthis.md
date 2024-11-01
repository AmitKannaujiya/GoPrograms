# Question : [Minimum add to make parentheses valid - (Leetcode : 921)](https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/description/)

A parentheses string is valid if and only if:

- It is the empty string,
- It can be written as AB (A concatenated with B), where A and B are valid strings, or
- It can be written as (A), where A is a valid string.
You are given a parentheses string s. In one move, you can insert a parenthesis at any position of the string.

- For example, if s = "()))", you can insert an opening parenthesis to be "(()))" or a closing parenthesis to be "())))".
Return the minimum number of moves required to make s valid.

### Example 1

```
Input: s = "())"
Output: 1

```

### Example 2

```
Input: s = "((("
Output: 3

```

### Constraints

-    n == nums.length
-    1 <= k <= n <= 10^5
-    -10^4 <= nums[i] <= 10^4


## Solution

```GO
package main
import "fmt"

func minAddToMakeValid(s string) int {
    stack := []byte{}
    count := 0
	// if opening bracket insert it else pop and count it 
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

// Function to initialize the recursion with starting indices of 0
func main() {
    fmt.Println(minAddToMakeValid("(((()")) //  print : 3
}

Time Complexity = O(n)
Space Complexity = O(n)
=> where n is the lengths of the array, respectively.
```