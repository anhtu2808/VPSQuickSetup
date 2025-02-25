package cli

import "github.com/anhtu2808/VPSQuickSetup/src/config"

type Cli interface {
	// Set data for the script | if need there are some vars use fmt.Scan
	SetData()
	// Build the script
	ScriptBuilder() (string, error)
	// Execute the script
	ScriptExecutor(*config.Config) error
}
