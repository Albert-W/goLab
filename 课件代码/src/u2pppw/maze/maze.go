package main

import (
	"fmt"
	"os"
)

//返加二维slice
func readMaze(filename string) [][]int {
	//打开文件
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	//6 行；
	maze := make([][]int, row)
	for i := range maze {
		//5列
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

//用x, y容易混
type point struct {
	i, j int
}
//四个方向
var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

//找下个点
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	//防止越界
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	//返加 对应的标记
	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {

	//新建record slice; 初始化都是0；
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	//初始队列
	Q := []point{start}

	//队列不空时
	for len(Q) > 0 {
		//取当前主node 为 cur
		cur := Q[0]
		Q = Q[1:]
		//发现终点，退出
		if cur == end {
			break
		}

		for _, dir := range dirs {
			//找到下个点，为next
			next := cur.add(dir)

			//maze at next is 1
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			// steps at next isn't 0
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}
			//找到当前的步数，并记录
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] =
				curSteps + 1
			//添加候选节点
			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze := readMaze("src/u2pppw/maze/maze.in")
	//二维slice; 调用walk函数
	steps := walk(maze, point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})

	//打印迷宫
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val) //3位对齐
		}
		fmt.Println()
	}

	// TODO: construct path from steps
}
