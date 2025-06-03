package registry

import (
	"github.com/spf13/cobra"
)

var AllCommands = []*cobra.Command{}

// RegisterCommand adds a command to the AllCommands slice
func RegisterCommand(cmd *cobra.Command) {
	AllCommands = append(AllCommands, cmd)
}

// GetAllCommands returns all registered commands
func GetAllCommands() []*cobra.Command {
	return AllCommands
}
