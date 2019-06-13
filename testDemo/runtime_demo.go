// runtime包测试

package main

import (
	"fmt"
	"runtime"
)

func goro1() {
	fmt.Print("hello goroutine 1")
}

func main() {
	go goro1()
	fmt.Print(runtime.NumCPU())        // 4
	fmt.Print(runtime.NumGoroutine())  // 2
	fmt.Println(runtime.GOMAXPROCS(2)) // 4
	i := 2
	fmt.Println(runtime.GOMAXPROCS(i)) // 2
}
