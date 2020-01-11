package main

import (
	"fmt"
)

func bubble_sort(n *[]float64) {
	for i := 0; i < len(*n); i++ {
		for j := 0; j < len(*n)-1-i; j++ {
			if (*n)[j] > (*n)[j+1] {
				(*n)[j], (*n)[j+1] = (*n)[j+1], (*n)[j]
			}
		}
	}
}

func main() {
	n := []float64{13.2, 3.2, 73, -0.1, 2.1, -5.2, -1.1, 0.2}
    bubble_sort(&n)
    fmt.Println(n)
}


