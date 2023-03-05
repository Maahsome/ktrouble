# ktrouble

A CLI tool for launching troubleshooting docker images into a kubernetes cluster

## Installation

```bash
brew install maahsome/tap/ktrouble --formula
```

## TODO

- [ ] Validate the name of the utility container?
- [x] Display a list of choices of utility container if NO args[0] is passed in
- [x] Display a list of choices of K8s ServiceAccounts if no args[1] is passed in
- [x] Provide a switch to directly submit the manifest to the KUBECONFIG defined current context
  - this is just default, no switch
- [x] Add a deeper definition for utility containers, specifying sizes (requests/limits)
- Add a check to see if the POD has already been created
- [x] Add a "delete" container command
- [ ] Add a config.yaml based list of container details

## PERSONAL JIRA LIST

```zsh
switch-jira kt
jira readme
```

### To Do

- [ ] KT-1:    Add EXAMPLES and Documentation 
- [ ] KT-3:    Find and add an ansible container 
- [ ] KT-6:    Convert to real OUTPUT formats 
- [ ] KT-7:    Replace logrus with common.Logger 
- [ ] KT-8:    In the delete command, when no pods are running, exit with that description 
- [ ] KT-9:    Extend the delete command to look at the first param after delete and use that as the delete POD name 
- [ ] KT-10:   Fix a bug where the utilitydefinitions are detected as empty, and defaults are written to config.yaml 
- [ ] KT-11:   Read the utilitydefinitions into a global variable rather than re-read config all the time (both MAP and ARRAY) 

### Done

- [x] KT-4:    Add a LIST for running PODs
- [x] KT-5:    Add a LIST for defined container images
- [x] KT-2:    Move container list details to config.yaml, create an initial version
- [x] KT-12:   Add a basic-tools image combining some of the others

