package main

import (
	"fmt"

	"github.com/anhtu2808/VPSQuickSetup/src/app"
	"github.com/anhtu2808/VPSQuickSetup/src/constants"
	"github.com/anhtu2808/VPSQuickSetup/src/utils/helper"
)

func main() {

	for {
		helper.ClearScreen()
		fmt.Print(constants.MainMenu)
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			app.GetHostIP()
		case 2:
			app.InstallNginx()
		case 0:
			return
		}

	}

}
