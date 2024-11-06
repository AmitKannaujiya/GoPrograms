package dfs

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