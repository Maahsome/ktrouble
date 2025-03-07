# ktrouble

A CLI tool for launching troubleshooting docker images into a kubernetes cluster

## Installation

```bash
brew install maahsome/tap/ktrouble --formula
```

## Getting Started

Once you have `ktrouble` installed, here are the quick getting running steps:

```zsh
# run version command, really any command, will build the default `config.yaml` file
ktrouble version
# at this point you can run quite a few commands as long as you are connected
# to a kubernetes cluster (current context)
# this will list the utilities stored in the `config.yaml` file
ktrouble get utility
# this command will prompt to launch a utility
ktrouble launch       # also ktrouble l
# adding a new locally defined utility pod
ktrouble add utility -u helm-kubectl311 -c "/bin/bash" -r "dtzar/helm-kubectl:3.11"
# modifying an item to keep from prompting to upload on push
ktrouble update utility -u helm-kubectl311 --toggle-exclude
# modify an item to change the command
ktrouble update utility -u helm-kubectl311 -c '/bin/sh'
# remove a local item
ktrouble remove utility -u helm-kubectl311
# hide an item that has an upstream source 'ktrouble-utils' repository
# this will remove the item from the prompt when using 'ktrouble launch'
ktrouble remove utility -u alpine3
ktrouble update utility -u alpine3 --toggle-hidden

# Interactions with 'futurama/farnsworth/tools/ktrouble-utils' repository
# setup the git credentials
# if you have an environment variable that contains the token, eg. GITLAB_TOKEN
ktrouble set config --tokenvar GITHUB_TOKEN --user christopher.maahs
# if you would rather store the token in the config.yaml file
ktrouble set config --token "<your token>" --user christopher.maahs

# once configured, you can run the commands to interact with the repository
# to pull new items into your local config.yaml file
# this will display a list of utility definitions not already in your local config
ktrouble pull
# a status of utility definitions
ktrouble status
# to pull items that are listed as "different"
ktrouble pull -a
# to get a verbose listing of definitions
ktrouble get util -f ktrouble get util -f name,repository,exec,hidden,excluded,source

# and or course
ktrouble --help
```

## Methodology Change

Originally just running `ktrouble` would start the prompt process.  This required
breaking the cobra methodologies which resulted in the `--help` features not being
built correctly.  Usage discovery through `--help` is an important feature to me, so
I have refactored to using a command to start the launching process `ktrouble launch`
or shortcut `ktrouble l`.

## Sharing of Utility Definitions

I figure there may as well be an easy way to share definitions.  So some commands
will be added to assist with CRUD operations on the local `config.yaml` file for
utility definitions.  Then `push|pull|status` commands will be added that will
interact with a git repository where a list of utility pod definitions will be
stored.  The initial population of the `config.yaml` utility definitions will
also attempt to pull from the central repository before defaulting to the defaults
in the `defaults` package in the code.

## PERSONAL JIRA LIST

```zsh
switch-jira kt
jira readme
```

### In Flight

- [ ] KT-16:  Start adding godoc comments (In Progress)

### To Do

- [ ] KT-1:   Add EXAMPLES and Documentation
- [ ] KT-3:   Find and add an ansible container
- [ ] KT-8:   In the delete command, when no pods are running, exit with that description
- [ ] KT-9:   Extend the delete command to look at the first param after delete and use that as the delete POD name
- [ ] KT-10:  Fix a bug where the utilitydefinitions are detected as empty, and defaults are written to config.yaml
- [ ] KT-18:  Add command line parameters to the launch command
- [ ] KT-29:  Add rebase command to pull ALL remote items, overwriting local versions
- [ ] KT-31:  Sort the LISTS, all of them
- [ ] KT-32:  Add a description/tools field, to house what tools/operation the utility definition is meant to solve
- [ ] KT-45:  Add a "defaultTemplate" setting in the config.yaml, and default to "default", and use that as the template when no --template/-t is provided
- [ ] KT-49:  Add create volume, with a --dir or --file parameter to populate
- [ ] KT-50:  Add populate volume with a --dir or --file parameter
- [ ] KT-51:  Add --volumes to the launch command to list volumes with app=ktrouble label
- [ ] KT-52:  Examine the exec parameter in kubectl code, can we use that?

