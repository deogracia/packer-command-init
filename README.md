packer-command-init
===================

Packer's Custom Command 

It creates all the default files and directories. For now, there all blanks.

## Install ##


Not sure if it's necessary.
> Download and build Packer from source as described [here](https://github.com/mitchellh/packer#developing-packer).

Next, clone this repository into `$GOPATH/src/github.com/deogracia/packer-command-init`.  Then build the packer-provisioner-ansible-local binary:


    go build -o $GOPATH/packer-command-init \
    plugin/packer-command-init/main.go


Now [configure Packer](http://www.packer.io/docs/other/core-configuration.html) to pick up the new command:


    {
      "commands": {
        "init": "packer-command-init"
      }
    }
