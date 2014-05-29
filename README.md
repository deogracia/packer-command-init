packer-command-init
===================

Packer's Custom Command 

It creates all the default files and directories. For now, there all blanks.

## Install ##

### Compilation ###
Not sure if it's necessary.
> Download and build Packer from source as described [here](https://github.com/mitchellh/packer#developing-packer).

Next, clone this repository into `$GOPATH/src/github.com/deogracia/packer-command-init`.  Then build the packer-provisioner-ansible-local binary:


    go build -o $GOPATH/bin/packer-command-init \
    plugin/packer-command-init/main.go

### Configuration ###
Source : [Packer's documentation](http://www.packer.io/docs/extend/plugins.html#toc_2)

Now [configure Packer](http://www.packer.io/docs/other/core-configuration.html) to pick up the new command:


    {
      "commands": {
        "init": "packer-command-init"
      }
    }

Run packer without command should give you :

    usage: packer [--version] [--help] <command> [<args>]
    
    Available commands are:
        build       build image(s) from template
        fix         fixes templates from old versions of packer
        init        Create the minimum file and directories needed
        inspect     see components of a template
        validate    check that a template is valid
    
    Globally recognized options:
        -machine-readable    Machine-readable output format.



## Thanks to ... ###
- [Kelsey Hightower](https://github.com/kelseyhightower), his [ansible provisionner](https://github.com/kelseyhightower/packer-provisioner-ansible-local) helped me to grasp the packer's plugion hierarchy.
- [Mitchell Hashimoto](https://github.com/mitchellh)
