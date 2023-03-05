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


