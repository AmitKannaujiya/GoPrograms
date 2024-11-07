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

func surroundedRegion(board [][]byte)  {
    // first dp dfs on boundry 0
    for i:= 0; i<len(board); i++ {
        for j:=0; j< len(board[i]); j++ {
            if (i == 0 || j == 0 || i == len(board)-1 || j == len(board[i])-1) && board[i][j] == 'O' {
                dfsUtillSurroundedRegion(board, i, j)
            }
        }
    }
    // now make every 0 as X and every * as 0
    for i:= 0; i<len(board); i++ {
        for j:=0; j< len(board[i]); j++ {
            if board[i][j] == 'O' {
                board[i][j] = 'X'
            }
            if board[i][j] == '*' {
                board[i][j] = 'O'
            }
        }
    }
}

func dfsUtillSurroundedRegion(board [][]byte, i, j int) {
    //if board[i][j] == 'O' {
        board[i][j] = '*'
    //}
    x := []int{-1, 1, 0, 0}
    y := []int{0, 0, -1, 1}
    for d:= 0; d< 4; d++ {
        dx := i + x[d]
        dy := j + y[d]
        if isSafeBoard(board, dx, dy) {
            dfsUtillSurroundedRegion(board, dx, dy)
        }
    }
}

func isSafeBoard(board [][]byte, i, j int) bool {
    if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) {
        return false
    }
    if board[i][j] != 'O' {
        return false
    }
    return true
}