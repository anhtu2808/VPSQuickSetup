package app

import (
	"fmt"
	"log"

	"github.com/anhtu2808/VPSQuickSetup/src/cli"
	"github.com/anhtu2808/VPSQuickSetup/src/config"
	"github.com/anhtu2808/VPSQuickSetup/src/constants"
	"github.com/anhtu2808/VPSQuickSetup/src/utils/helper"
)

func GetHostIP() {
	config, err := config.LoadConfig(constants.ConfigPath)
	if err != nil {
		log.Fatalf("Load config error: %v", err)
	}
	for {
		// Clear the screen
		helper.ClearScreen()
		fmt.Print(constants.GetHostIPMenu)
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 0:
			return
		case 1:
			app := &cli.ShowIP{}
			app.SetData()
			app.ScriptExecutor(config)
			helper.PauseScreen()
		}
	}
}
