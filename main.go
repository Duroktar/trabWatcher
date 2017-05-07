package main

import (
	"errors"
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

func handleArgs() (string, []string, error) {
	if len(os.Args) < 3 {
		return "", nil, errors.New("Not enough args")
	}
	return os.Args[1], os.Args[2:], nil
}

func main() {
	command, files, err := handleArgs()
	if err != nil {
		fmt.Println(aurora.Red("usage: trabWatcher \"command\" [files]").Bold())
		os.Exit(1)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	w := watcher.NewWatcher(command, files)
	go handleCtrlC(c, w)

	fmt.Println(aurora.Magenta("Starting watcher."))
	w.Start()
}
