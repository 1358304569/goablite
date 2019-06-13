/*
3. C/S版--服务器端

参考来源：https://blog.csdn.net/dalerkd/article/details/78291214


*/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数
	if openLog {
		fmt.Println(r.Form)
		fmt.Println("path", r.URL.Path)
		fmt.Println("scheme", r.URL.Scheme)
		fmt.Println(r.Form["url_long"])
	}
	if openLog {
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}
	}

	io.WriteString(w, "Hello,Test Me!")
	fmt.Fprintf(w, "Hello!")

	My_Mem_Prof()
	/* 不退出测试长连接性能 */
	chLongLink <- true
	<-chExitThreads
}

var chExitThreads chan bool
var chLongLink = make(chan bool, 1000*100)

func openService() {
	http.HandleFunc("/MyWeb", helloHandler)
	err := http.ListenAndServe(":8686", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}

const openLog bool = false

var ch_sync chan bool

func calc() {
	number := 0
	for {
		_, ok := <-chLongLink
		if !ok {
			break
		}
		number++
		if (number % 100) == 0 {
			fmt.Println("now long link:", number)
		}
	}
	fmt.Println("All long link number is:", number)
	close(ch_sync)
	fmt.Println("End")
}
func My_Mem_Prof() {
	fm, err := os.OpenFile("./tmp/mem.out", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	pprof.WriteHeapProfile(fm)
	fm.Close()
}

func My_Cpu_Prof() {
	f, err := os.OpenFile("./tmp/cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	<-chExitThreads
	pprof.StopCPUProfile()
	fmt.Println("End CPU Prof")
	f.Close()
	close(chExitProf)
}

var chExitProf chan bool

func main() {
	/*初始化channel通知*/
	chExitThreads = make(chan bool)
	ch_sync = make(chan bool)
	chExitProf = make(chan bool)

	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	fmt.Println("CPU Number:", runtime.NumCPU())
	fmt.Println("Thread Number:", runtime.NumCPU()*2)
	go My_Cpu_Prof()

	exp := `Please visit: "http://localhost:8686/MyWeb" `
	/////////////////
	fmt.Println(exp)
	go calc()
	go openService()

	/* 直到用户输入才开始退出协程 */
	input := make([]byte, 1024)
	os.Stdin.Read(input)
	close(chLongLink)
	<-ch_sync
	fmt.Println("OK")
	close(chExitThreads)
	<-chExitProf
}
