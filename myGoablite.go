/*

 */

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
)

// 参数
var (
	n = flag.Int("n", 20, "Number of requests to run.")
	c = flag.Int("c", 5, "Number of requests to run concurrently.\nTotal number of requests cannot be smaller than the concurrency level.")
	p = flag.Int("p", runtime.GOMAXPROCS(-1), "Number of used cpu cores")
	u = flag.String("u", "", "the url you want to connect.")
)

// 结果
var (
	success = 0.0
	failure = 0.0
	useTime = 0.0
)

// 帮助信息
var usage = `Usage: hey [options...] <url>

Options:
  -n  Number of requests to run. Default is 20.
  -c  Number of requests to run concurrently. Total number of requests cannot
      be smaller than the concurrency level. Default is 5.
  -p  Number of used cpu cores.(default for current machine is %d cores)
`

var wg sync.WaitGroup

// 主函数
func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, runtime.NumCPU()))
	}
	flag.Parse()
	// fmt.Println("开始解析了么？？？？")

	// 判断参数是否正常
	if flag.NFlag() < 1 {
		usageAndExit("")
	}
	runtime.GOMAXPROCS(*p)
	num := *n
	conc := *c
	url := *u

	if url == "" || len(url) == 0 {
		usageAndExit("Please input the url you want to connect.")
	}
	if num <= 0 || conc <= 0 {
		usageAndExit("-n and -c cannot be smaller than 1.")
	}

	if num < conc {
		usageAndExit("-n cannot be less than -c.")
	}
	if *p > runtime.NumCPU() {
		fmt.Printf("your computer have %d CPU cores\n", runtime.NumCPU())
		usageAndExit("-p cannot bigger than that your computer have.")

	}

	startTime := time.Now().UnixNano()

	// 并发开始
	for i := 0; i < conc; i++ {
		wg.Add(1)
		go run(num / conc)
	}
	// fmt.Println("主程序开始wait")
	wg.Wait()
	endTime := time.Now().UnixNano()
	useTime = float64(endTime-startTime) / 1e9

	// 输出结果
	fmt.Println("Plan_Total:", *n)
	fmt.Println("Concurrency Level:", *c)
	fmt.Println("CPU Core:", *p)
	fmt.Println()
	fmt.Println("Complete requests:", success)
	fmt.Println("Failed requests:", failure)
	// fmt.Println("SuccessRate:", fmt.Sprintf("%.2f", ((success/total)*100.0)), "%")
	fmt.Println("UseTime:", fmt.Sprintf("%.4f", useTime), "s")
	fmt.Println("Requests per second:", fmt.Sprintf("%.4f", float64(*n)/useTime))

}

// 运行每个客户端
func run(num int) {

	defer wg.Done()

	no := 0.0
	ok := 0.0

	for i := 0; i < num; i++ {
		resp, err := http.Get(*u)

		if err != nil {
			no += 1
			continue
		}

		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			no += 1
			continue
		}

		ok += 1
		continue
	}

	success += ok
	failure += no
	// total += float64(num)

}

// 提示信息
func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}
