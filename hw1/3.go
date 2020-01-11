package main

import (
	"fmt"
	"os"
)

func fibonacci2(n int) int {
    if n < 3 { 
    	return 1 
    }
    return fibonacci2(n - 2) + fibonacci2(n - 1)
}

func main(){
	var n int
    fmt.Fscan(os.Stdin, &n)
	fmt.Println(fibonacci2(n))
}
