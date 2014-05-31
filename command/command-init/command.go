package commandinit

import (
	"github.com/mitchellh/packer/packer"
	"io/ioutil"
	"strings"
)

type Command struct{}

func (Command) Help() string {
	return strings.TrimSpace(helpString)
}

func (c Command) Synopsis() string {
	return "Create the minimum files and directories needed"
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (c Command) Run(env packer.Environment, args []string) int {
	ui := env.Ui()
	allGreen := []byte("All green! ^^")
	// Create default tempalte file
	err := ioutil.WriteFile("./default.json",[]byte(defaultTemplate),0644)
	check(err)
	ui.Say(string(allGreen))
	return 0
}
