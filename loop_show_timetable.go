package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/codeskyblue/go-sh"
)

func loopShowTimeTable() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	err := sh.Command("tput", "civis").Run()
	die(err)

	err = sh.Command("tput", "smcup").Run()
	die(err)

	go func() {
		<-c

		err := sh.Command("tput", "clear").Run()
		die(err)

		err = sh.Command("tput", "rmcup").Run()
		die(err)

		err = sh.Command("tput", "cnorm").Run()
		die(err)

		os.Exit(0)
	}()

	for {
		err := sh.Command("tput", "home").Run()
		die(err)

		showTimeTable()
		time.Sleep(time.Second)
	}
}
