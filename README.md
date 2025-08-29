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
      image: gcr.io/kubernetes-e2e-test-images/dnsutils
      execcommand: /bin/sh
      requiresecrets: false
      requireconfigmaps: false
      excludefromshare: false
      hidden: false
      hint: Image has basic DNS investigation tools like nslookup and dig
      environments: []
      tags:
        - 1.2
        - 1.3
```

In order to support custom docker registries, we use an `environments`
mechanism.  There are environments defined in the config file
like this:

```yaml
environments:
    - excludefromshare: false
      hidden: false
      name: lowers
      removeupstream: false
      repository: us-docker.pkg.dev/lowers-repo
    - excludefromshare: false
      hidden: false
      name: uppers
      removeupstream: false
      repository: us-docker.pkg.dev/uppers-repo
```

This might represent dev/prod, uppers/lowers, however you like to define each
environment. Once these  environments are defined, you can assign environments
to a `utility definition`.  When we go to launch a POD, the utility will be
presented multiple times, prefixed with each of the `environment` names.  The
`image` in the POD manifest will combine the `repository` from the `environment`
with the `image/tags` from the `utility` definition, separated by a `/` of
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
ktrouble add utility -u helm-kubectl -c "/bin/bash" -i "dtzar/helm-kubectl" --tags '3.11,3.19'
# modifying an item to keep from prompting to upload on push
ktrouble update utility -u helm-kubectl --toggle-exclude
# modify an item to change the command
ktrouble update utility -u helm-kubectl -c '/bin/sh'
# remove a local item
ktrouble remove utility -u helm-kubectl
# hide a utility
# this will remove the item from the prompt when using 'ktrouble launch'
# and 'ktrouble get utils' commands
ktrouble update utility -u alpine3 --toggle-hidden
ktrouble get utils
ktrouble get utils --show-hidden

# to get a verbose listing of definitions
ktrouble get util -f name,image,tags,hidden,requiresecrets,requireconfigmaps
# the list of field names for each command can be learned from
ktrouble fields

# an example of a docker image that might live in two different registries
ktrouble add environment --name lowers -r "us-docker.pkg.dev/lowers-repo"
ktrouble add environment --name uppers -r "us-docker.pkg.dev/uppers-repo"
ktrouble get environments
ktrouble add utility -u custom-utility -i "docker/custom-utility" --tags '0.0.1' -e 'lowers,uppers' -c '/bin/bash'
ktrouble get utility
# show the environments for all the defined utilities
ktrouble get utility --fields 'NAME,IMAGE,TAGS,ENVIRONMENTS'

# Interactions with 'maahsome/ktrouble-utils' repository
# setup the git credentials
# if you have an environment variable that contains the token, eg. GITHUB_TOKEN
ktrouble set config --tokenvar GITHUB_TOKEN --user "<yourgithub user>"
# if you would rather store the token in the config.yaml file
ktrouble set config --token "<your token>" --user "<your github user>"

# once configured, you can run the commands to interact with the repository
# a status of utility definitions
ktrouble status
# to pull new items into your local config.yaml file
ktrouble pull
# to pull items that are listed as "different"
ktrouble pull -a
# get a list of differences for items marked as "different" in the status
ktrobule diff
# to push one of your local utility definitions up to the common repository
ktrouble push
# if you want to remove an upstream utility definition, you will need to use the
# '--remove-upstream' parameter when you remove it
# this will set the 'hidden' and also 'removeupstream' settings
# the utility will be removed from both the upstream and local on the
# next 'ktrouble push' that is run
ktrouble remove utility -u netshoot --remove-upstream
ktrouble push
# similarly environments are removed from the upstream repository the same way
ktrouble remove environment -e lowers --remove-upstream
ktrouble push --env

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

# You can also create your own empty git repository and set the URL to that
# repository, just replace the maahsome repository with your own
ktrouble set config --giturl " https://github.com/maahsome/ktrouble-utils.git"
# you will be prompted to run 'ktrouble migrate' which will create a 'v2'
# directory in your repository, once created you wil be able to interact with
# your repository

# and or course
ktrouble --help
```
