/*
goroutine测试

*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1) // 相当于串行
	// runtime.GOMAXPROCS(2) // 并行
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("a--i: ", i)
			wg.Done()
		}()
		time.Sleep(time.Second) // 这句注释掉就只能看到a--i:10
		/*因为协程与主协程是并发，等打印的时候，在主协程中的i早就运行完了，只能得到10*/
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("b--i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
