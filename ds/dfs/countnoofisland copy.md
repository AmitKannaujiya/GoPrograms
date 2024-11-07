# Question : [Number of Island - (Leetcode : 200)](https://leetcode.com/problems/number-of-islands/description/)

Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

### Example 1

```
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1

```

### Example 2

```
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3

```

### Constraints

-   m == grid.length
-   n == grid[i].length
-   1 <= m, n <= 300
-   grid[i][j] is '0' or '1'.


## Solution

```GO
package main
import "fmt"

func numIslands(grid [][]byte) int {
    count := 0
    for i:= 0; i<len(grid); i++ {
        for j:=0; j< len(grid[i]); j++ {
            if grid[i][j] == '1' {
                count++
                dfsUtill(grid, i, j)
            }
        }
    }
    return count
}

func dfsUtill(grid [][]byte, i, j int) {
    grid[i][j] = '2'
    x := []int{-1, 1, 0, 0}
    y := []int{0, 0, -1, 1}
    for d:= 0; d< 4; d++ {
        dx := i + x[d]
        dy := j + y[d]
        if isSafe(grid, dx, dy) {
            dfsUtill(grid, dx, dy)
        }
    }
}

func isSafe(grid [][]byte, i, j int) bool {
    if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
        return false
    }
    if grid[i][j] != '1' {
        return false
    }
    return true
}

// Function to initialize the recursion with starting indices of 0
func main() {
    grid1 := [][]byte{
		{'1','1','1','1','0'},
		{'1','1','0','1','0'},
		{'1','1','0','0','0'},
		{'0','0','0','0','0'},
	}

	grid2 := [][]byte{
		{'1','1','0','0','0'},
		{'1','1','0','0','0'},
		{'0','0','1','0','0'},
		{'0','0','0','1','1'},
	}
    fmt.Println(numIslands(grid1)) //  print : 1
    fmt.Println(numIslands(grid2)) //  print : true
}

Time Complexity = O(n)
Space Complexity = O(1)
=> where n is the lengths of the array, respectively.
```