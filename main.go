//go:generate rsrc -ico icon.ico -o resources.syso
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/getlantern/systray"
)

func main() {
	go Initiate()

	onExit := func() {
		now := time.Now()
		os.WriteFile(fmt.Sprintf(`on_exit_%d.txt`, now.UnixNano()), []byte(now.String()), 0644)
	}

	go systray.Run(onReady, onExit)

	wait := make(chan struct{})
	exit := make(chan struct{})
	go user_input(wait, exit)

	select {
	case <-wait:
		go user_input(wait, exit)
	case <-exit:
		os.Exit(0)
	}
}
