package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func main() {
	systray.Run(onReady, onExit)
}

func getTopResourcesConsumption() string {
	//var b bytes.Buffer

	result, _ := exec.Command("bash", "-c", "ps -eo pid,%mem --sort=-%mem | head").Output()
	return string(result)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Pretty awesome")
	menuData := getTopResourcesConsumption()
	fmt.Println(menuData)
	menuDataSlice := strings.Split(menuData, "\n")

	for i, s := range menuDataSlice {
		createButton(s, string(i))

	}
	go func() {
		<-ClickedCh
		fmt.Println("Requesting quit")
	}()
	/*systray.AddMenuItem("Quit", "Quit the whole app")
	for {
		select {
		case <-button1.ClickedCh:
			fmt.Println("ca affiche!!!")
		case <-button2.ClickedCh:
			fmt.Println("ca affiche!!!")
		case <-button3.ClickedCh:
			fmt.Println("ca affiche!!!")
		case <-button4.ClickedCh:
			fmt.Println("ca affiche!!!")
		case <-button5.ClickedCh:
			fmt.Println("ca affiche!!!")
		case <-button6.ClickedCh:
			fmt.Println("ca affiche!!!")
		}
	}*/
	//}()
	//fmt.Println(*systray.parent)
}

func onExit() {
	// clean up here
}

func createButton(name string, hint string) {
	systray.AddMenuItem(name, hint)
}
