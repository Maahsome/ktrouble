# ktrouble

A CLI tool for launching troubleshooting docker images into a kubernetes cluster

## Installation

```bash
brew install maahsome/tap/ktrouble --formula
```

## PERSONAL JIRA LIST

```zsh
switch-jira kt
jira readme
```

### To Do

- [ ] KT-1:   Add EXAMPLES and Documentation
- [ ] KT-3:   Find and add an ansible container
- [ ] KT-8:   In the delete command, when no pods are running, exit with that description
- [ ] KT-9:   Extend the delete command to look at the first param after delete and use that as the delete POD name
- [ ] KT-10:  Fix a bug where the utilitydefinitions are detected as empty, and defaults are written to config.yaml
- [ ] KT-16:  Start adding godoc comments (In Progress)
- [ ] KT-18:  Add command line parameters to the launch command
- [ ] KT-29:  Add rebase command to pull ALL remote items, overwriting local versions
- [ ] KT-30:  Add --all to pull command, to prompt to select from ALL upstream utilities to pull from

### Done

- [x] KT-4:   Add a LIST for running PODs
- [x] KT-5:   Add a LIST for defined container images
- [x] KT-2:   Move container list details to config.yaml, create an initial version
- [x] KT-12:  Add a basic-tools image combining some of the others
- [x] KT-6:   Convert to real OUTPUT formats
- [x] KT-11:  Read the utilitydefinitions into a global variable rather than re-read config all the time (both MAP and ARRAY)
- [x] KT-7:   Replace logrus with common.Logger
- [x] KT-13:  Add a bash:// column to the get pods output
- [x] KT-14:  Add get sizes to display the request/limits for each size
- [x] KT-15:  Turn off auto-format for headers, change header names
- [x] KT-17:  Refactor for clarity
- [x] KT-19:  Add an add command to add a utility definition
- [x] KT-27:  Update the update of the "source" property, so that it only changes ones that are "" blank AND also in the "default" list, and setting all others to "local"
- [x] KT-21:  Add a remove command to remove an existing utility definition
- [x] KT-22:  Add a local property "excludeFromShare" to exclude an item from being pushed
- [x] KT-23:  Add a "source" property, indicating if a utility object is from the upstream source
- [x] KT-20:  Add an update command to update an existing utility definition
- [x] KT-24:  Add a pull command to display a list of utilities that are on the upstream source, but not downloaded locally, allow an "All" or "multi-select" to choose which to pull
- [x] KT-28:  Add set gituser, set gittokenvar, and set gittoken to facilitate setting these config.yaml settings for interation with git
- [x] KT-26:  Add a status command that will compare your local config.yaml definitions with the upstream source
- [x] KT-25:  Add a push command to push items that are not marked as "excludeFromShare"
