// os.Args和flag测试
/*
1. go build os_xxx.go

2. go_xxx.exe arg1 arg2 arg3 arg4...
*/

package main

import (
	"fmt"
	// "os"
	"flag"
)

func main() {
	/*os.args测试*/
	// args1 := os.Args
	// args2 := os.Args[1:]
	// args3 := os.Args[3]
	// fmt.Println(args1)
	// fmt.Println(args2)
	// fmt.Println(args3)

	/*flag测试*/
	s := flag.String("1", "a", "first")
	i := flag.Int("2", 0, "second")
	// b := flag.Bool("3", false, "third")

	flag.Parse()

	fmt.Printf("第一个--%s\n", *s)
	fmt.Printf("第二个--%d\n", *i)
	fmt.Println("args", flag.Args())
	fmt.Println("Nargs", flag.NArg())
	fmt.Println("Nflag", flag.NFlag())
}
