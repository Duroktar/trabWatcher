package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/logrusorgru/aurora"
)

func lastModified(file string) (string, error) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return "", err
	}
	return fileInfo.ModTime().String(), nil
}

func launchCommand(command string) (*exec.Cmd, error) {
	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd, cmd.Run()
}

// newWatcher runs given cmd when any file given in watched changes
func newWatcher(cmd string, watched []string) *watcher {
	var w watcher
	w.modTimes = make(map[string]string)
	w.Targets = watched
	w.Command = cmd
	w.tick = 1
	return &w
}

// watcher object
type watcher struct {
	Command  string
	Targets  []string
	process  *exec.Cmd
	modTimes map[string]string
	tick     int
}

// start the watcher
func (w watcher) start() {
	w.spawnProcess()
	for {
		w.checkFiles()
		time.Sleep(time.Duration(w.tick) * time.Second)
	}
}

func (w watcher) checkFiles() {
	for _, target := range w.Targets {
		modTime, err := lastModified(target)
		if err != nil {
			continue
		}
		if modTime != w.modTimes[target] {
			w.spawnProcess()
		}
	}
}

func (w watcher) spawnProcess() {
	if &w.process == nil {
		w.killProcess()
	}
	w.getLastModifiedTimes()
	cmd, err := launchCommand(w.Command)
	if err != nil {
		return
	}
	w.process = cmd
}

func (w watcher) getLastModifiedTimes() {
	for _, target := range w.Targets {
		modTime, err := lastModified(target)
		if err != nil {
			fmt.Println(aurora.Red("Error initializing.").Bold())
			continue
		}
		w.modTimes[target] = modTime
	}
}

func (w watcher) killProcess() {
	if w.process.ProcessState.Exited() {
		return
	}
	err := w.process.Process.Kill()
	if err != nil {
		panic(err)
	}
}

func handleCtrlC(c chan os.Signal) {
	sig := <-c
	if sig != os.Interrupt {
		return
	}
	fmt.Println(aurora.Green("\rSignal: "), sig)
	fmt.Println(aurora.Magenta("Goodbye"))
	os.Exit(0)
}

func main() {
	fmt.Println(aurora.Magenta("Starting watcher."))

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go handleCtrlC(c)

	newWatcher(os.Args[1], os.Args[2:]).start()
}
