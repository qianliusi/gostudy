package leetcode

import (
	"container/list"
)

type Point struct {
	x, y int
}

func Sum(n int) (res int) {
	for n >= 10 {
		mod := n % 10
		n /= 10
		res += mod
	}
	res += n
	return
}

func MovingCount(m int, n int, k int) int {
	queue := list.New()
	//正确
	visit := make([][]bool, m)
	for a := range visit {
		visit[a] = make([]bool, n)
	}
	point := Point{0, 0}
	queue.PushBack(point)
	x, y := 0, 0
	res := 0
	for x < m && y < n && queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		p := front.Value.(Point)
		if visit[p.x][p.y] {
			continue
		}
		visit[p.x][p.y] = true
		res++
		x, y = p.x, p.y
		if x+1 < m && Sum(x+1)+Sum(y) <= k {
			queue.PushBack(Point{x + 1, y})
		}
		if y+1 < n && Sum(x)+Sum(y+1) <= k {
			queue.PushBack(Point{x, y + 1})
		}
	}
	return res
}
