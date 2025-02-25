package cli

import (
	"github.com/anhtu2808/VPSQuickSetup/src/config"
	"github.com/anhtu2808/VPSQuickSetup/src/handler"
)

type ShowIP struct {
	Command string
}

func (s *ShowIP) SetData() {
	s.Command = "hostname -I | awk '{print $1}'"
}

func (s *ShowIP) ScriptBuilder() (string, error) {
	return s.Command, nil
}

func (s *ShowIP) ScriptExecutor(config *config.Config) error {
	return handler.RunScriptRemote(s.ScriptBuilder, config)
}
