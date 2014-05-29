package commandinit

import (
	//"flag"
	//"fmt"
	"github.com/mitchellh/packer/packer"
	//"log"
	//"sort"
	"strings"
)

type Command struct{}

func (Command) Help() string {
	return strings.TrimSpace(helpString)
}

func (c Command) Synopsis() string {
	return "Create the minimum files and directories needed"
}

func (c Command) Run(env packer.Environment, args []string) int {
	ui := env.Ui()
	ui.Say("All green! ^^")
	return 0
}
