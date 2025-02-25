package cli

import (
	"github.com/anhtu2808/VPSQuickSetup/src/config"
	"github.com/anhtu2808/VPSQuickSetup/src/handler"
)

type InstallNginx struct {
	Command string
}

func (i *InstallNginx) SetData() {
	i.Command = "sudo apt update && sudo apt install nginx -y"
}

func (i *InstallNginx) ScriptBuilder() (string, error) {
	return i.Command, nil
}

func (i *InstallNginx) ScriptExecutor(config *config.Config) error {
	return handler.RunScriptRemote(i.ScriptBuilder, config)
}
