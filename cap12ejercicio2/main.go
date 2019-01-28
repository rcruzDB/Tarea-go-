package main

import (
	
)

type cycle struct {
	value int
	tail  *cycle
}

func main() {
	var c cycle
	c = cycle{42, &c}

	display.Display("cycle", c)
}