package main

import "fmt"

func swap(x, y *int) {
	temp := *x;
	*x = *y;
	*y = temp;
}

func main() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Printf("x = %d, y = %d\n", x, y)
}
