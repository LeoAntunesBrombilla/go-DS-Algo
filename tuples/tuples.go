package main

import "fmt"

func powerSeries(a int) (int, int, error) {
	var square = a * a
	var cube = square * a

	return square, cube, nil
}

func main() {
	fmt.Println(powerSeries(5))
}
