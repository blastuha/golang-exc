package main

import (
	"fmt"
)

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println("X,Y", x, y)
}
