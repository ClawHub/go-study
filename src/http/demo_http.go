package http

import (
	"context"
	"fmt"
	"go-study/src/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//测试
func DemoHttp() {
	log.HttpLogger.Info("-------------http-------------")
	startWebService("", "9997", "", "", false)
}

//开始web服务
func startWebService(host, port, certFile, keyFile string, enableSSL bool) {
	log.HttpLogger.Info("http server starting......")

	http.HandleFunc("/welcome", welcome)

	hostPort := host + ":" + port
	//http服务配置
	server := &http.Server{
		Addr:           hostPort,
		Handler:        http.DefaultServeMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//启动新协程
	go func() {
		var err error
		//如果允许ssl
		if enableSSL {
			err = server.ListenAndServeTLS(certFile, keyFile)
		} else {
			//ListenAndServe监听TCP网络地址。然后调用来处理传入连接上的请求。已接受的连接被配置为启用TCP keep-alives
			err = server.ListenAndServe()
		}
		if err == nil {
			log.HttpLogger.Info(fmt.Sprintf("success to bind %s", hostPort))
		} else {
			log.HttpLogger.Info(fmt.Sprintf("fail to bind %s %s", hostPort, err))
		}
	}()
	//获取os.Signal类型信道
	signalChannel := make(chan os.Signal)
	//监听
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	//如果没有相关事件时，阻塞
	<-signalChannel
	//关闭web服务
	err := server.Shutdown(context.Background())
	if err != nil {
		log.HttpLogger.Error("server shutdown fail")
	}
	log.HttpLogger.Info("http server closing......")
}

//welcome
func welcome(rw http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintln(rw, "welcome")
}
