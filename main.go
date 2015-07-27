package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/kisom/goutils/die"
)

var (
	waitFor time.Duration
	verbose bool
	cmdName = "xscreensaver-command"
)

func scanForXScreenSaver() {
	_, err := exec.LookPath(cmdName)
	if err != nil {
		die.With("xscreensaver-command not found: please install it via your package manager.")
	}
}

func heartbeat() {
	for {
		<-time.After(waitFor)
		cmd := exec.Command(cmdName, "-deactivate")
		fmt.Printf("%+v\n", cmd)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "[!] %v\n", err)
		}

		if verbose {
			fmt.Printf("%s\n", bytes.TrimSpace(out))
		}
	}
}

func help() {
	fmt.Fprintf(os.Stderr, `hbxss [-i interval] [-t time] [-v]

	-f    	  	Force hbxss to run indefinitely.
	
	-i interval	Specify the interval between heartbeats. This
			should follow the form <number><unit>, where
			unit should be one of 's', 'm', or 'h' for
			seconds, minutes, or hours, respectively.

	-t time		Specify how long the program should run for;
	   		the default is two hours.

	-v		Print each heartbeat as it occurs.
`)
}

func init() {
	flag.Usage = help
}

func main() {
	var runFor time.Duration

	forceForever := flag.Bool("f", false, "Force hbxss to run forever.")
	showHelp := flag.Bool("h", false, "Display a short usage message and exit.")

	flag.DurationVar(&waitFor, "i", 5*time.Minute, "Time between heartbeats.")
	flag.DurationVar(&runFor, "t", 2 * time.Hour, "Duration program should run.")
	flag.BoolVar(&verbose, "v", false, "Print each heartbeat.")
	flag.Parse()

	if *showHelp {
		help()
		os.Exit(0)
	}

	scanForXScreenSaver()
	go heartbeat()

	if *forceForever {
		fmt.Fprintf(os.Stderr, "!!! Warning: xscreensaver will be indefinitely suppressed !!!\n")
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, os.Kill, os.Interrupt, syscall.SIGTERM)
		<-sigc
	} else {
		<-time.After(runFor)
	}
}
