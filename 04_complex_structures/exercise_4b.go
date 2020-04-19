package main

import "fmt"

func average(args ...float64) float64 {

	sum := 0.0

	for _, value := range args {
		sum += value
	}
	return sum / float64(len(args))
}

func main() {
	fmt.Println(average(3.4, 5.6, 7.8))
}
