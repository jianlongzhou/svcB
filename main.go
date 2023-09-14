package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// 定义路由处理函数
	http.HandleFunc("/service/b", helloHandler)
	fmt.Println("start service: B")

	go func() {
		for {
			request()
			time.Sleep(time.Second)
		}
	}()

	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("start service err: ", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get request header: ", r.Header)
	fmt.Fprint(w, "B, Hello, World! version: ", os.Getenv("VERSION"))
}

func request() {
	host := os.Getenv("HOST")
	resp, err := http.Get(fmt.Sprintf("%s/service/c", host))
	if err != nil {
		fmt.Println("call service c err: ", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read service rsp body err: ", err)
		return
	}
	fmt.Println("svcB -> svcC: ", string(body))
}
