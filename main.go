package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Duroktar/trabWatcher/watcher"
	"github.com/logrusorgru/aurora"
)

func handleCtrlC(c chan os.Signal, w *watcher.Watcher) {
	sig := <-c
	if sig != os.Interrupt {
		return
	}
	fmt.Println(aurora.Green("\rSignal: "), sig)
	fmt.Println(aurora.Magenta("Goodbye"))
	w.Stop()
	os.Exit(0)
}

func main() {
	fmt.Println(aurora.Magenta("Starting watcher."))

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	w := watcher.NewWatcher(os.Args[1], os.Args[2:])
	go handleCtrlC(c, w)

	w.Start()
}
