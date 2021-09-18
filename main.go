package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Runtime: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
}