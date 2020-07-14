package 算法

import "math"

/*
A*算法
*/

type Grid struct {
	x      int   //x坐标
	y      int   //y坐标
	f      int   // G+H
	g      int   //从起点到当前格子花费了多少步
	h      int   //在不考虑障碍的情况下，从当前各自走到目标格子的距离，也就是距离目标格子还有多元
	parent *Grid //父节点
}

func (grid *Grid) InitGrid(x, y int, parent *Grid, end *Grid) {
	grid.x = x
	grid.y = y
	grid.parent = parent
	if parent == nil {
		grid.g = 1
	} else {
		grid.g = parent.g + 1
	}
	grid.h = int(math.Abs(float64(grid.x-end.x)) + math.Abs(float64(grid.y-end.y)))
	grid.f = grid.g + grid.h
}

func aStarSearch(start Grid, end Grid){

}
