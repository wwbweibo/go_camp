package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return initAndListenHttpServer(ctx)
	})

	group.Go(func() error {
		return signalHandler(ctx)
	})

	err := group.Wait()
	log.Printf("errgroup err received err %v, exited", err)
}

func initAndListenHttpServer(ctx context.Context) error {
	log.Printf("init http server")
	server := &http.Server{
		Addr: "0.0.0.0:8000",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello from server"))
		}),
	}
	go func() {
		select {
		case <-ctx.Done():
			log.Printf("context done receive, exiting http server")
			_ = server.Close()
		}
	}()
	log.Printf("starting http server")
	return server.ListenAndServe()
}

// signalHandler will listen system signal and handle it.
func signalHandler(ctx context.Context) error {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	select {
	case <-ctx.Done():
		log.Printf("context done receive, exiting system signal listen")
		return ctx.Err()
	case sig := <-sigs:
		log.Printf("receive sig %s, exiting...", sig)
		return fmt.Errorf("receive sig %v, exiting", sig)
	}

}
