package main


import (
	"fmt"
	"os"
)


func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)		// 先读取行列

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		// 读取每个数据
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}


type point struct {
	i, j int
}


var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}


func (p point) add(r point) point {
	return point{p.i+r.i, p.j+r.j}
}


// bool表示这个点是否有效
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}


func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))		// 记录走过的路 以及最终路径
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}  // 先放入一个起始点

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]	// 出队

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			val, ok := next.at(maze)
			// 如果点无效 或是障碍物
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			// 如果已经走过了
			if !ok || val != 0 {
				continue
			}

			// 回到原点
			if next == start {
				continue
			}

			// 当前步数加一
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			// 放入队列
			Q = append(Q, next)
		}
	}

	return steps
}


func main() {
	maze := readMaze("maze.txt")

	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}

	steps := walk(maze, point{0, 0}, point{len(maze)-1, len(maze[0])-1 })

	fmt.Println("--------------")
	for _, row := range steps {
		for _, val := range row {
			// 3位对齐（数据最多两位）
			fmt.Printf("%3d", val)	
		}
		fmt.Println()
	}
}