/*
2. 单机版，flag取参数

参考来源：https://blog.csdn.net/oLeiShen/article/details/84106157

参数：
	-c:连接数
	-t:并发数
	-u:url

使用方法：
		xx.exe -c=200 -t=3 -u=http://www.baidu.com

*/

package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	c = flag.Int("c", 1, "Plz input client quantity")
	t = flag.Int("t", 1, "Plz input times quantity")
	u = flag.String("u", "http://www.baidu.com", "Plz input url")
)

var (
	total    = 0.0
	about    = 0.0
	success  = 0.0
	failure  = 0.0
	use_time = 0.0
)

var wg sync.WaitGroup

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
	total += float64(num)

}

func main() {

	start_time := time.Now().UnixNano()

	flag.Parse()

	if *c == 0 || *t == 0 || *u == "" {
		flag.PrintDefaults()
		return
	}

	fmt.Println("c:", *c, ",t:", *t)

	for i := 0; i < *c; i++ {
		wg.Add(1)
		go run(*t)
	}

	wg.Wait()
	end_time := time.Now().UnixNano()

	fmt.Println("PreTotal:", (*c)*(*t))
	fmt.Println("Total:", total)
	fmt.Println("Success:", success)
	fmt.Println("Failure:", failure)
	fmt.Println("SuccessRate:", fmt.Sprintf("%.2f", ((success/total)*100.0)), "%")
	fmt.Println("UseTime:", fmt.Sprintf("%.4f", float64(end_time-start_time)/1e9), "s")
}
