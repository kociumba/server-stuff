package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/getlantern/systray"
)

// onReady initializes the system tray display when the app opens.
//
// This function does not take any parameters.
// It does not return anything.
func onReady() {
	// this works only if the image is in the same directory
	//
	// icon, err := os.ReadFile("icon.ico")
	// if err != nil {
	// 	panic(err)
	// }

	resp, err := http.Get("https://raw.githubusercontent.com/kociumba/server-stuff/main/icon.ico")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	icon, err := io.ReadAll(resp.Body)

	systray.SetIcon(icon)
	systray.SetTooltip("balls")
	exitApp := systray.AddMenuItem("Exit", "Exit the app")
	go func() {
		<-exitApp.ClickedCh
		fmt.Println("exiting...")
		systray.Quit()
		os.Exit(0)
	}()
}
