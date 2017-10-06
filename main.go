package main

import (
	"net/http"
	"log"
	"thinkgo/router"
	_ "thinkgo/controller"
)


func main(){
	http.HandleFunc("/", router.HttpRouter)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("main.go main() err := http.ListenAndServe() error")
	}
}


/*
export GOROOT=/usr/local/go/
export GOPATH=/root/lesliehuang/goproj
export PATH=$PATH:/usr/local/go/bin/
*/
