# Question : [Backspace String Compare - (Leetcode : 844)](https://leetcode.com/problems/backspace-string-compare/description/)

Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.

**Note** that after backspacing an empty text, the text will continue empty.

### Example 1

```
Input: s = "ab#c", t = "ad#c"
Output: true
Explanation: Both s and t become "ac".

```

### Example 2

```
Input: s = "ab##", t = "c#d#"
Output: true
Explanation: Both s and t become "".

```

### Example 3

```
Input: s = "a#c", t = "b"
Output: false
Explanation: s becomes "c" while t becomes "b".

```

### Constraints

-   1 <= s.length, t.length <= 200
-   s and t only contain lowercase letters and '#' characters.

**Follow up:** Can you solve it in *O(n)* time and *O(1)* space?

## Solution

```GO
package main
import "fmt"

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

// Function to initialize the recursion with starting indices of 0
func main() {
    fmt.Println(backspaceCompare("ab##", "ad#c")) //  print : true
    fmt.Println(backspaceCompare("ab##", "c#d#")) //  print : true
    fmt.Println(backspaceCompare("a#c", "b")) //  print : false
}

Time Complexity = O(n)
Space Complexity = O(n)
=> where n is the lengths of the array, respectively.
```