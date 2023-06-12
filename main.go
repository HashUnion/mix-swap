package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/DwGoing/mix-swap/internal/app"
	"github.com/alibaba/ioc-golang"
)

func main() {
	if err := ioc.Load(); err != nil {
		panic(err)
	}
	_, err := app.GetAppSingleton()
	if err != nil {
		panic(err)
	}

	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig
}
