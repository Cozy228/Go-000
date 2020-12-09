package main

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	g, cancelCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return listenSignal(cancelCtx)
	})
	g.Go(func() error {
		return startServer(cancelCtx, ":8080")
	})
	if err := g.Wait(); err != nil {
		log.Println("error group return err:", err.Error())
	}

	log.Println("DONE!")
}

func listenSignal(ctx context.Context) error {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case <-signalChan:
		return errors.New("server close -> close signal received")
	case <-ctx.Done():
		return errors.New("server close -> context done")
	}
}

func startServer(ctx context.Context, addr string) error {
	svr := &http.Server{Addr: addr, Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "Hello, Ziyu! ID:G20200607010463\n")
	})}
	go func() {
		select {
		case <-ctx.Done():
			shutdownServer(svr, ctx)
		}
	}()
	log.Println("HTTP Server ready to start...")
	return svr.ListenAndServe()
}

func shutdownServer(server *http.Server, ctx context.Context) {
	_ = server.Shutdown(ctx)
}
