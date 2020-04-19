package main

import "fmt"

func average(arg1, arg2, arg3 float64) float64 {

	return (arg1 + arg2 + arg3) / 3
}

func main() {
	fmt.Println(average(3.4, 5.6, 7.8))
}
