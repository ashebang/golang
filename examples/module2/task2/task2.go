package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

/*
1.接收客户端request，并将request 中带的header 写入response header
2.读取当前系统的环境变量中的VERSION 配置，并写入response header
3.Server 端记录访问日志包括客户端IP，HTTP 返回码，输出到server 端的标准输出
4.当访问localhost/healthz 时，应返回200
*/

func index(w http.ResponseWriter, req *http.Request) {
	//	fmt.Fprintf(w, "hello\n")

	for name, index := range req.Header {
		for _, h := range index {
			fmt.Fprintf(w, "%v: %v\n", name, index)
			w.Header().Set(name, h)
		}
	}

	// 2.读取当前系统的环境变量中的VERSION 配置，并写入response header
	// $ VERSION=1.20 go run <my.go> &
	//	os.Setenv("VERSION", "1.23")
	version := os.Getenv("VERSION")
	fmt.Println(version)
	w.Header().Set("VERSION", version)

	// 3.Server 端记录访问日志包括客户端IP，HTTP 返回码，输出到server 端的标准输出

	clientIP := GetClientIP(req)
	log.Printf("Response Code: %d", http.StatusOK)
	log.Printf("Client IP: %s", clientIP)

}

func GetClientIP(req *http.Request) string {
	ipaddr := req.Header.Get("X-Real-IP")
	if ipaddr == "" {
		ipaddr = req.Header.Get("X-Forwarder-For")
	}
	if ipaddr == "" {
		//fmt.Println(req.RemoteAddr)
		ipaddr = req.RemoteAddr
	}
	host, _, _ := net.SplitHostPort(ipaddr)
	return host
}

func main() {

	//	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
