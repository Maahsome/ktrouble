# ktrouble

A CLI tool for launching troubleshooting docker images into a kubernetes cluster

## Installation

```bash
brew install maahsome/tap/ktrouble --formula
```

## How everything is put together

The `ktrouble` tool keeps a list of utility definitions, and some tool
configuration in a local config file.  Each definition has a configuration that
looks like this:

```yaml
utilitydefinitions:
    - name: dnsutils
      repository: gcr.io/kubernetes-e2e-test-images/dnsutils:1.3
      execcommand: /bin/sh
      requiresecrets: false
      requireconfigmaps: false
      excludefromshare: false
      hidden: false
      hint: Image has basic DNS investigation tools like nslookup and dig
      environments: []
```

The `repository` property here should likely be named `image`, since it points
to the image itself.  In order to support custom docker registries, we use an
`environments` mechanism.  There are environments defined in the config file
like this:

```yaml
environments:
    - excludefromshare: false
      hidden: false
      name: lowers
      repository: us-docker.pkg.dev/lowers-repo
    - excludefromshare: false
      hidden: false
      name: uppers
      repository: us-docker.pkg.dev/uppers-repo
```

This might represent dev/prod, uppers/lowers, however you like to define each
environment. Once these  environments are defined, you can assign environments
to a `utility definition`.  When we go to launch a POD, the utility will be
presented multiple times, prefixed with each of the `environment` names.  The
`image` in the POD manifest will combine the `repository` from the `environment`
with the `repository` from the `utility` definition, separated by a `/` of
course.

### git repository

As an initial version, the interaction with the defined git repository was to
simply store each utility definition as a `.yaml` file in the root of the
repository.  As features have been added that required a change in the config
structures, the `v0.n.n` version of `ktrouble` maintains the list in the root of
the repository.  When the `major` of the application `semver` changes, a new
directory wil be created in the repository, eg: "v1" for `v1.n.n`, and "v2" for
`v2.n.n`.  The `ktrouble` tool will prompt to perform a migration, in which it
will read all the files from the previous directory, modify the definitions, and
write them out to the new directory, where `ktrouble` will interact for that
`version`.

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

# to get a verbose listing of definitions
ktrouble get util -f name,repository,exec,hidden,excluded,source

# an example of a docker image that might live in two different registries
ktrouble add environment -n lowers -r "us-docker.pkg.dev/lowers-repo"
ktrouble add environment -n uppers -r "us-docker.pkg.dev/uppers-repo"
ktrouble get environments
ktrouble add utility -u custom-utility -r "docker/custom-utility:0.0.1" -e 'lowers,uppers' -c '/bin/bash'
ktrouble get utility
# show the environments for all the defined utilities
ktrouble get utility --fields 'NAME,REPOSITORY,ENVIRONMENTS'

# Interactions with 'maahsome/ktrouble-utils' repository
# setup the git credentials
# if you have an environment variable that contains the token, eg. GITHUB_TOKEN
ktrouble set config --tokenvar GITHUB_TOKEN --user cmaahs
# if you would rather store the token in the config.yaml file
ktrouble set config --token "<your token>" --user cmaahs

# once configured, you can run the commands to interact with the repository
# to pull new items into your local config.yaml file
# this will display a list of utility definitions not already in your local config
ktrouble pull
# a status of utility definitions
ktrouble status
# to pull items that are listed as "different"
ktrouble pull -a
# get a list of differences for items marked as "different" in the status
ktrobule diff
# to push one of your local utility definitions up to the common repository
ktrouble push

# the environments can also be stored and sourced from the defined git repository.
# we have extended the status,pull,push,diff methods to handle the environments
# See the status of defined environments
ktrouble status --env
# pull environment definitions from git, this simply pulls them all
ktrouble pull --env
# push environments to upstream, this also pushes all environments that are not
# set to be excluded
ktrouble push --env
# get a list of differences for items marked as "different" in the status
ktrouble diff --env

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
stored.

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
- [ ] KT-57:  Create docker image: aws+basic-tools
- [ ] KT-58:  Create docker image: gcloud+basic-tools
- [ ] KT-59:  Create docker image: azure+basic-tools
- [ ] KT-60:  Create docker image:helm+kubectl+basic-tools
- [ ] KT-63:  Add "get hint" and then prompt for the utility
- [ ] KT-65:  Add config for handling dynamic GAR and other repositories.
- [ ] KT-67:  Add a Utility definition to handle fsGroup, runAsUser, and runAsGroup
- [ ] KT-70:  When mounting configmaps with a `.` in the name, a naming error occurs
- [ ] KT-72:  Add the ability to change the "sleep" list from the CLI
- [ ] KT-73:  Add --build-cmd switch for launch and attach
- [ ] KT-76:  Better defaults around git configuration
- [ ] KT-82:  Create a PUBLIC docker image that can be used to test the ingress/service feature
- [ ] KT-85:  fetch an environments definition from the upstream git repository, since environments will be specific.
- [ ] KT-86:  Add the logger to the Config so it can be passed around
- [ ] KT-87:  Need a way to DELETE an upstream utility
- [ ] KT-88:  Need a way to DELETE an upstream environment
- [ ] KT-89:  Add a config option to set the FIELDS to be outputted for each object type
- [ ] KT-90:  Update ALL objects to use c.Fields to drive output fields

### Completed

- [x] KT-62:  Add "requireSecret", "requireConfigmap", and "containerHint" as YAML fields in the utility definitions
- [x] KT-66:  Fix: volume names can only be 63 characters long
- [x] KT-74:  Add environment support to support same images in different registries
