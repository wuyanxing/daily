package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var server *http.Server


type myHandler struct {
}

func (myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Method)
	fmt.Println(r.Header)
	var p []byte
	r.Body.Read(p)
	fmt.Println(p)

	w.Write([]byte("hello root"))
}

type inter interface {
	Speak()
}
type name struct {
	s string
}

func (n name) Speak()  {
	fmt.Println(n.s)
}

func test (f inter)  {
	f.Speak()
}
func main()  {

	nn := name{"name"}
	test(nn)


	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})

	server = &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		TLSConfig:         nil,
		ReadTimeout:       3,
		ReadHeaderTimeout: 0,
		WriteTimeout:      3,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	err := server.ListenAndServe()

	if err != nil{
		if err == http.ErrServerClosed{
			log.Fatal("closed")
		}else {
			log.Fatal("err closed")
		}
	}

}

/*
总结：
	在golang中没有thread的概念，goruntine来代替golang，用法 go func()。
	chan 为golang的内置变量，代表一个管道，可以读写，channel = make(chan string)
	在sync包中，有一个WaitGroup函数用来管理goruntine，
		sw := sync.WaitGroup{3}  初始化一个三个协程的sw
		sw.Done() 释放资源，当sw变成0时，sw.Wait()退出
		sw.Wait()  阻塞等待协程结束，类似于read.join()

	goruntine存在panic机制，panic是指协程出现异常，当发生panic时，则整个进程会挂掉。但是可以使用recover函数进行捕捉，
	类似于 try{}catch{}finally{}中的finally

	注意::golang中存在大小写的区别，大写开头的变量或方法可以被外部包识别，小写开头的只能在本文件内使用

	http:
	http.NewServeMux()为http后台代理处理器，可以根据路由和uri来定位具体的处理函数。
	Handle函数为自定义的处理函数，用于处理url路径和具体处理函数，Handle接收pattern和Handler参数，其中pattern为具体的路径
	Handler为具体处理函数，Handler(w http.ResponseWriter, r *http.Request),可以通过自定义struct并实现ServeHTTP接口来重载
	Handler函数

	http.ServeHTTP为自带的httpserver，可以设置具体启动参数，并通过ListenAndServe()来启动server。默认的http.ServerHTTP实现了
	ListenAndServe()方法，可以new一个http.ServeHTTP实例并通过ListenAndServe()来启动
*/
