// os.Args测试
/*
1. go build os_xxx.go

2. go_xxx.exe arg1 arg2 arg3 arg4...
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	args1 := os.Args
	args2 := os.Args[1:]
	args3 := os.Args[3]
	fmt.Println(args1)
	fmt.Println(args2)
	fmt.Println(args3)
}
