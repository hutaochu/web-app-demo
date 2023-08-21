package main

import (
	"context"
	"flag"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/hutaochu/web-app-demo/myserver/docs"
	"github.com/hutaochu/web-app-demo/myserver/handlers"
	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	flag.Set("logtostderr", "false")
	flag.Set("log_file", "myfile.log")
	flag.Parse()
	defer klog.Flush()

	router := gin.Default()
	handlers.Run(router)

	addr := ":8080"
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	go func() {
		if err := http.ListenAndServe(":8081", http.DefaultServeMux); err != nil && err != http.ErrServerClosed {
			klog.Fatal("ListenAndServe error: ", err)
		}
	}()
	klog.Info("pprof server listening and serving HTTP on :8081")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			klog.Fatal("server listenAndServe failed, err:", err)
		}
	}()
	klog.Info("Listening and serving HTTP on ", addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	klog.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		klog.Fatal("Server Server error: ", err)
	}
	klog.Info("Server exiting")
}
