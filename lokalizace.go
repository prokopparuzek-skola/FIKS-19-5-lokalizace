package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rectangle struct {
	ru Point
	ld Point
}

func find(rects *[3]Rectangle) (out int) {
	for i := 0; i < 3; i++ {
		var a, b int
		a = rects[i].ru.x - rects[i].ld.x
		b = rects[i].ru.y - rects[i].ld.y
		out += a * b
	}
	var intersects []Rectangle
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 3; j++ {
			var tmp Rectangle
			var over bool
			tmp, over = intersect(rects[i], rects[j])
			if over {
				var a, b int
				a = tmp.ru.x - tmp.ld.x
				b = tmp.ru.y - tmp.ld.y
				out -= a * b
				intersects = append(intersects, tmp)
			}
		}
	}
	for i := 1; i < len(intersects); i++ {
		var tmp Rectangle
		var over bool
		tmp, over = intersect(intersects[0], intersects[i])
		if over {
			var a, b int
			a = tmp.ru.x - tmp.ld.x
			b = tmp.ru.y - tmp.ld.y
			out += a * b
			break
		}
	}
	return
}

func intersect(rect1, rect2 Rectangle) (inter Rectangle, over bool) {
	switch { // překrývají se?
	case rect2.ld.x >= rect1.ru.x:
		fallthrough
	case rect2.ld.y >= rect1.ru.y:
		fallthrough
	case rect2.ru.x <= rect1.ld.x:
		fallthrough
	case rect2.ru.y <= rect1.ld.y:
		over = false
		return
	}
	inter.ld.x = max(rect1.ld.x, rect2.ld.x)
	inter.ld.y = max(rect1.ld.y, rect2.ld.y)
	inter.ru.x = min(rect1.ru.x, rect2.ru.x)
	inter.ru.y = min(rect1.ru.y, rect2.ru.y)
	over = true
	return
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var T int
	fmt.Scan(&T)
	for ; T > 0; T-- {
		var rects [3]Rectangle
		for i := 0; i < 3; i++ {
			fmt.Scanln(&rects[i].ld.x, &rects[i].ld.y, &rects[i].ru.x, &rects[i].ru.y)
		}
		fmt.Println(find(&rects))
	}
}
