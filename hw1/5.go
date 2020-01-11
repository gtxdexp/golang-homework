package main

import (
	"fmt"
)

func unique_count (n*[] int) int{
	var temp int = 1
	bubble_sort(n)
	for i:=1; i<len(*n); i++{
		if (*n)[i-1] != (*n)[i]{
			temp++
		}
	}
	return temp
}

func bubble_sort(n *[]int) {
	for i := 0; i < len(*n); i++ {
		for j := 0; j < len(*n)-1-i; j++ {
			if (*n)[j] > (*n)[j+1] {
				(*n)[j], (*n)[j+1] = (*n)[j+1], (*n)[j]
			}
		}
	}
}

func main() {
	n := []int{1,2,3,4,2,1,2,3,4,2,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3}
    unique_count(&n)
    fmt.Println(unique_count(&n))
}