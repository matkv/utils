package cmd

import (
	"github.com/matkv/utils/internal/registry"
	"github.com/spf13/cobra"
)

// GetAllCommands returns all registered commands from the registry
func GetAllCommands() []*cobra.Command {
	return registry.GetAllCommands()
}
