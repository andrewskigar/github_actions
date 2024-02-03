package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("\n=======\n%s\n=======\n", runtime.GOOS)
}
