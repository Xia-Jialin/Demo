package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func init() {
	// file := "./" + "message" + ".log"
	// logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	// if err != nil {
	// 	panic(err)
	// }
	// log.SetOutput(logFile) // 将文件设置为log输出的文件
	// log.SetPrefix("[qSkipTool]")
	// log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	// return
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(20 * time.Second)
	fmt.Fprintf(w, "Hello Word!")
}
func main() {
	http.HandleFunc("/", helloHandler)
	ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	// 启动http1
	http1 := http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}
	// 监听sig信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func(c chan (os.Signal)) {
		for {
			select {
			case <-c:
				defer cancel()
				err := http1.Shutdown(context.TODO())
				if err != nil {
					fmt.Println(err)
				}
				return
			}
		}
	}(quit)

	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		if err := http1.ListenAndServe(); err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Println("all exit: ", err)
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("服务退出")
			return
		}
	}

}
