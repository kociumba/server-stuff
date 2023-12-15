package main

import (
	"fmt"
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

	icon, err := http.Get("https://raw.githubusercontent.com/getlantern/systray/main/icon.ico")

	systray.SetIcon([]byte(icon))
	systray.SetTooltip("balls")
	exitApp := systray.AddMenuItem("Exit", "Exit the app")
	go func() {
		<-exitApp.ClickedCh
		fmt.Println("exiting...")
		systray.Quit()
		os.Exit(0)
	}()
}
