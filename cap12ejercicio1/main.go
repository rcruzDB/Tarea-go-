package main

import (
	
)

func main() {
	display.Display("map with struct key", map[struct{ x, y int }]int{
		struct{ x, y int }{}:      0,
		struct{ x, y int }{1, 10}: 100,
	})

	display.Display("map with array key", map[[3]int]int{
		[3]int{}:        100,
		[3]int{0, 0, 1}: 100,
		[3]int{1, 4, 7}: 100,
	})
}