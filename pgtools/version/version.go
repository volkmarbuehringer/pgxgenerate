package version

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/sirupsen/logrus"
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
		syscall.SIGQUIT,
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

			case syscall.SIGQUIT:
				fmt.Println("stop and core dump")
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

	switch os.Getenv("PU_LOGLEVEL") {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	if t, ok := os.LookupEnv("PU_DEBUG"); ok {
		debug, _ := strconv.ParseBool(t)
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}

	}

	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Println(program, githash, githashtool, buildstamp)
		os.Exit(0)
	}

}
