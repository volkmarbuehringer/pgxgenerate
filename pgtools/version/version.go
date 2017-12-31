package version

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var githash string
var githashtool string
var buildstamp string
var program string

//Contexter installiert ein signal-handler und den cancel-Context
func Contexter() (context.Context, func()) {
	ctx, cancel := context.WithCancel(context.Background())

	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {
		for {
			s := <-signal_chan
			switch s {
			// kill -SIGHUP XXXX
			case syscall.SIGHUP:
				fmt.Println("hungup")
				cancel()

			// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				fmt.Println("Warikomi")
				cancel()

			// kill -SIGTERM XXXX
			case syscall.SIGTERM:
				fmt.Println("force stop")
				cancel()

			default:
				fmt.Println("Unknown signal.")
			}
		}
	}()
	return ctx, cancel
}

func init() {

	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Println(program, githash, githashtool, buildstamp)
		os.Exit(0)
	}

}
