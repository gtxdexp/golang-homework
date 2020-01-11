package main

import (
	"fmt"
	"os"
)

func multiply(a int, b int) int{
	var temp int
	for i := 1; i <= b; i++{
		temp += a
	}
	return temp
}

func main(){
	var a, b int
    fmt.Fscan(os.Stdin, &a)
    fmt.Fscan(os.Stdin, &b)
	fmt.Println(multiply(a, b))
}
