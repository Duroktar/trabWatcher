package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

// NewWatcher runs given cmd when any file given in watched changes
func NewWatcher(cmd string, watched []string) *Watcher {
	var w Watcher
	w.modTimes = make(map[string]string)
	w.Targets = watched
	w.Command = cmd
	w.tick = 1
	return &w
}

// Watcher object
type Watcher struct {
	Command  string
	Targets  []string
	process  *exec.Cmd
	modTimes map[string]string
	tick     int
}

// Start the watcher
func (w Watcher) Start() {
	w.spawnProcess()
	for {
		w.checkFiles()
		time.Sleep(time.Duration(w.tick) * time.Second)
	}
}

// SetTick sets how often in seconds files are checked for changes
func (w Watcher) SetTick(rate int) {
	w.tick = rate
}

func (w Watcher) checkFiles() {
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

func (w Watcher) spawnProcess() {
	if &w.process == nil {
		w.killProcess()
	}
	w.getLastModifiedTimes()
	cmd, err := launchCommand(w.Command)
	// if err != nil {
	// fmt.Println(aurora.Red("Error executing command").Bold())
	// panic(err)
	// }
	if err == nil {
		w.process = cmd
	}
}

func (w Watcher) getLastModifiedTimes() {
	for _, target := range w.Targets {
		modTime, err := lastModified(target)
		if err != nil {
			fmt.Println(aurora.Red("Error initializing.").Bold())
		}
		w.modTimes[target] = modTime
	}
}

func (w Watcher) killProcess() {
	if w.process.ProcessState.Exited() != true {
		err := w.process.Process.Kill()
		if err != nil {
			panic(err)
		}
	}
}
