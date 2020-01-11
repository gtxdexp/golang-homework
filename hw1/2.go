package main

import (
	"fmt"
	"os"
)

func fibonacci1(n int) int{
	var f int = 0 
	var f1 int = 1
	var f2 int = 1
	
	if (n == 1 || n == 2) {
        return 1
    }
    for i := 3; i <= n; i++ {
            f = f1 + f2
            f1 = f2
            f2 = f
    }
    return f
}

func main(){
	var n int
    fmt.Fscan(os.Stdin, &n)
	fmt.Println(fibonacci1(n))
}
