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

	button1 := systray.AddMenuItem(string(menuDataSlice[0]), " ")
	button2 := systray.AddMenuItem(string(menuDataSlice[1]), " ")
	button3 := systray.AddMenuItem(string(menuDataSlice[2]), " ")
	button4 := systray.AddMenuItem(string(menuDataSlice[3]), " ")
	button5 := systray.AddMenuItem(string(menuDataSlice[4]), " ")
	button6 := systray.AddMenuItem(string(menuDataSlice[5]), " ")

	/*for i, s := range menuDataSlice {
	    println(i)

	}*/

	systray.AddMenuItem("Quit", "Quit the whole app")
	for {
		select {
		case <-button1.ClickedCh:
			fmt.Println("Process kill")
		case <-button2.ClickedCh:
			fmt.Println("Process kill")
		case <-button3.ClickedCh:
			fmt.Println("Process kill")
		case <-button4.ClickedCh:
			fmt.Println("Process kill")
		case <-button5.ClickedCh:
			fmt.Println("Process kill")
		case <-button6.ClickedCh:
			fmt.Println("Process kill")
		}
	}
	//}()
	//fmt.Println(*systray.parent)
}

func onExit() {
	// clean up here
}
