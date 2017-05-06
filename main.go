package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/logrusorgru/aurora"
)

func handleCtrlC(c chan os.Signal) {
	sig := <-c
	fmt.Println(aurora.Green("\rSignal: "), sig)
	if sig == os.Interrupt {
		fmt.Println(aurora.Magenta("\rGoodbye"))
		os.Exit(0)
	}
}

func main() {
	fmt.Println(aurora.Magenta("Starting watcher."))

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go handleCtrlC(c)

	NewWatcher(os.Args[1], os.Args[2:]).Start()
}
