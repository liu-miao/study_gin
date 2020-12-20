package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//有返回值不想处理 就可以 用这个接收    _,_
	_, _ = fmt.Fprintln(w, "Hello Golang!")

}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server faild,err:%v\n", err)
		return

	}

}
