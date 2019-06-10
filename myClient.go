// 客户端程序

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

func client(exit_ch chan int, webPath string) {

	var err error

	resp, err := http.Get(webPath)
	if err != nil {
		// handle error
		fmt.Println(err.Error())
		goto threadEnd
	}

	/*body*/
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err.Error())
		goto threadEnd
	}
	defer resp.Body.Close()
	//fmt.Println(string(body))
threadEnd:
	exit_ch <- 1

}

func main() {
	var webPath string
	flag.StringVar(&webPath, "webpath", "http://baidu.com", "Tested Web Path")
	test_time_work := flag.Int("time", 50, "Test times")
	flag.Parse()
	fmt.Println(webPath)
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)

	exit_ch := make(chan int)

	test_time := time.Now()

	for i := 0; i < *test_time_work; i++ {
		go client(exit_ch, webPath)
	}

	for i := 0; i < *test_time_work; i++ {
		<-exit_ch
		fmt.Println("Success time!:", i+1)
	}
	test_time_end := time.Since(test_time)
	fmt.Println("End*******************")
	fmt.Println("CPU Number:", runtime.NumCPU())
	fmt.Println(test_time_end)
}
