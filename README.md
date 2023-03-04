# ktrouble

A CLI tool for launching troubleshooting docker images into a kubernetes cluster

## TODO

- Validate the name of the utility container
- Display a list of choices of utility container if NO args[0] is passed in
- Display a list of choices of K8s ServiceAccounts if no args[1] is passed in
- Provide a switch to directly submit the manifest to the KUBECONFIG defined current context
- Add a deeper definition for utility containers, specifying sizes (requests/limits)
- Add a check to see if the POD has already been created
- Add a "delete" container switch
