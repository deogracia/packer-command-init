package main

import (
	"github.com/mitchellh/packer/packer/plugin"
	"github.com/deogracia/packer-command-init/command/command-init"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterCommand(new(commandinit.Command))
	server.Serve()
}