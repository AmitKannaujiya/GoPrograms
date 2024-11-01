# Question : [String Manipulation - (GFG)](https://www.geeksforgeeks.org/problems/string-manipulation3706/1)

Tom is a string freak. He has got sequences of words arr[] to manipulate. If in a sequence, two same words come together then Tom destroys each other. Find the number of words left in the sequence after this pairwise destruction.

### Example 1

```
Input: arr[] = ["ab", "aa", "aa", "bcd", "ab"]
Output: 3
Explanation: After the first iteration, we'll have: ab bcd ab. We can't further destroy more strings and hence we stop and the result is 3. 

```
### Example 2

```
Input: arr[] = ["tom", "jerry", "jerry", "tom"]
Output: 0
Explanation: After the first iteration, we'll have: tom tom. After the second iteration: 'empty-array' .Hence, the result is 0.

```

-  *Expected Time Complexity:* O(n)
-  *Expected Auxiliary Space:* O(n)

### Constraints

-   1 ≤ arr.size() ≤10^6
-   1 ≤ |arri| ≤ 50


## Solution

```GO
package main
import (
	"fmt"
	"string"
)

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

// Function to initialize the recursion with starting indices of 0
func main() {
    words := []string{"ab", "aa", "aa", "bcd", "ab"}
    fmt.Println(removeConsecutiveSame(words)) //  print : 3
}

Time Complexity = O(n)
Space Complexity = O(n)
=> where n is the lengths of the array, respectively.
```