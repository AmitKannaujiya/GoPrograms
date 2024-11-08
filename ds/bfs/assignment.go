package bfs

type Queue[T any] struct {
	Enqueue func(T)
	Dequeue func()T
	Front func() T
	Back func() T
	IsEmpty func() bool
	Size func() int
}

func CreateQueue[T any]() *Queue[T] {
	elements := make([]T, 0)
	return &Queue[T]{
		Enqueue: func(t T) {
			elements = append(elements, t)
		},
		Dequeue: func() T {
			t := elements[0]
			elements = elements[1:]
			return t
		},
		Front: func() T {
			return elements[0]
		},
		Back: func() T {
			return elements[len(elements) - 1]
		},
		IsEmpty: func() bool {
			return len(elements) == 0
		},
		Size: func() int {
			return len(elements)
		},
	}
}

func numIslands(grid [][]byte) int {
	count := 0
    for i:= 0; i<len(grid); i++ {
        for j:=0; j< len(grid[i]); j++ {
            if grid[i][j] == '1' {
                count++
                bfsUtill(grid, i, j)
            }
        }
    }
    return count
}
type Data struct {
	row int
	col int
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
func bfsUtill(grid [][]byte, i, j int) {
	//grid[i][j] = '2'
    x := []int{-1, 1, 0, 0}
    y := []int{0, 0, -1, 1}
	queue := CreateQueue[Data]()
	queue.Enqueue(Data{i, j})

	for !queue.IsEmpty() {
		size := queue.Size()
		for i:= 0; i< size; i ++ {
			data := queue.Dequeue()
			grid[data.row][data.col] = '2'
			for d := 0; d < 4; d++ {
				dx := data.row + x[d]
				dy := data.col + y[d]
				if isSafe(grid, dx, dy) {
					queue.Enqueue(Data{dx, dy})
				}
			}
		}
	}
}

func LevelOrderTravesal()