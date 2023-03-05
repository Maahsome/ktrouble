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

- [ ] KT-1:    Add EXAMPLES and Documentation
- [ ] KT-3:    Find and add an ansible container
- [ ] KT-8:    In the delete command, when no pods are running, exit with that description
- [ ] KT-9:    Extend the delete command to look at the first param after delete and use that as the delete POD name
- [ ] KT-10:   Fix a bug where the utilitydefinitions are detected as empty, and defaults are written to config.yaml
- [ ] KT-16:   Start adding godoc comments  (In Progress)
- [ ] KT-18:   Add command line parameters to `launch` rather than just using the prompts

### Done

- [x] KT-4:    Add a LIST for running PODs
- [x] KT-5:    Add a LIST for defined container images
- [x] KT-2:    Move container list details to config.yaml, create an initial version
- [x] KT-12:   Add a basic-tools image combining some of the others
- [x] KT-6:    Convert to real OUTPUT formats
- [x] KT-11:   Read the utilitydefinitions into a global variable rather than re-read config all the time (both MAP and ARRAY)
- [x] KT-7:    Replace logrus with common.Logger
- [x] KT-13:   Add a bash:// column to the get pods output
- [x] KT-14:   Add get sizes to display the request/limits for each size
- [x] KT-15:   Turn off auto-format for headers, change header names
- [x] KT-17:   Refactor for clarity
