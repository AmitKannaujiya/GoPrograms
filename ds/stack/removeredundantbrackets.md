# Question : [Expression contains redundant bracket or not - (GFG)](https://www.geeksforgeeks.org/problems/expression-contains-redundant-bracket-or-not/1)

Given a string of balanced expression str, find if it contains a redundant parenthesis or not. A set of parenthesis are redundant if the same sub-expression is surrounded by unnecessary or multiple brackets. Return 1 ifit contains a redundant parenthesis, else 0.
Note: Expression may contain + , - , *, and / operators. Given expression is valid and there are no white spaces present.

### Example 1

```
Input: exp = ((a+b))
Output: Yes
Explanation: ((a+b)) can reduced to (a+b).

```

### Example 2

```
Input: exp = (a+b+(c+d))
Output: No
Explanation: (a+b+(c+d)) doesn't have any redundant or multiple brackets.

```

### Example 3

```
Input: exp = a+b+(c+d)
Output: No
Explanation: a+b+(c+d) doesn't have any redundant or multiple brackets.

```
### Example 3

```
Input: exp = (a+b+(c+d)
Output: No
Explanation: a+b+(c+d) doesn't have any redundant or multiple brackets.

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