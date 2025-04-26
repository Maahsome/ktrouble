# ktrouble help for all commands

## TOC

- [_main_](#ktrouble)
- [add](#add)
- [add utility](#add-utility)
- [changelog](#changelog)
- [delete](#delete)
- [diff](#diff)
- [edit](#edit)
- [edit config](#edit-config)
- [edit template](#edit-template)
- [fields](#fields)
- [get](#get)
- [get configs](#get-configs)
- [get namespace](#get-namespace)
- [get node](#get-node)
- [get nodelabels](#get-nodelabels)
- [get running](#get-running)
- [get serviceaccount](#get-serviceaccount)
- [get sizes](#get-sizes)
- [get templates](#get-templates)
- [get utilities](#get-utilities)
- [launch](#launch)
- [pull](#pull)
- [push](#push)
- [remove](#remove)
- [remove utility](#remove-utility)
- [set](#set)
- [set config](#set-config)
- [status](#status)
- [update](#update)
- [update utility](#update-utility)
- [version](#version)

## ktrouble

```plaintext
EXAMPLE:
  Simply run the 'launch' command and you will be prompted for all of the
  required details.
    - Utility Pod Selection
    - Namespace
    - Service Account
    - Node Selector
    - Resource Sizing

  > ktrouble launch

Usage:
  ktrouble [command]

Available Commands:
  add         Add various objects for ktrouble
  changelog   Get changelog information
  completion  Generate the autocompletion script for the specified shell
  delete      Delete PODs that have been created by ktrouble
  diff        Get a context diff on each utility definition
  edit        Edit various objects for ktrouble
  fields      Display a list of valid fields to use with the --fields/-f parameter for each command
  genhelp     Output help from all the sub commands
  get         Get various internal configuration and kubernetes resource listings
  help        Help about any command
  launch      Launch a kubernetes troubleshooting pod
  pull        Pull utility definitions from git
  push        Push local objects to upstream git repository
  remove      Remove various objects for ktrouble
  set         Set various objects for ktrouble
  status      Get a comparison of the local utility definitions with the upstream one
  update      Update various objects for ktrouble
  version     Express the 'version' of ktrouble

Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")

Use "ktrouble [command] --help" for more information about a command.
```

[TOC](#TOC)

## add

```plaintext
EXAMPLE:
  Use the "add utility" command to add a new utility definition to your 'config.yaml'

    > ktrouble add utility --help

Usage:
  ktrouble add [flags]
  ktrouble add [command]

Available Commands:
  utility     Add a utility definition to the ktrouble configuration

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")

Use "ktrouble add [command] --help" for more information about a command.
```

[TOC](#TOC)

## add utility

```plaintext
EXAMPLE:
  Use 'add utility' command to add a new utility definition to your 'config.yaml'
  file

    > ktrouble add utility -u helm-kubectl311 -c "/bin/bash" -r "dtzar/helm-kubectl:3.11"

Usage:
  ktrouble add utility [flags]

Flags:
  -c, --cmd string           Default shell/command to use when 'exec'ing into the POD (default "/bin/sh")
  -e, --exclude              Exclude from 'push' to central repository
      --hint-file string     Specify a file containing the hint text
  -u, --name string          Unique name for your utility pod
  -r, --repository string    Repository and tag for your utility container, eg: cmaahs/basic-tools:latest
      --require-configmaps   Set the Utilty to always prompt for configmaps
      --require-secrets      Set the Utilty to always prompt for secrets

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## changelog

```plaintext
EXAMPLE:
  Get just the latest changelog entry

    > ktrouble changelog

EXAMPLE:
  Get all the changelog entries

    > ktrouble changelog --all

Usage:
  ktrouble changelog [flags]

Aliases:
  changelog, cl, changes

Flags:
      --all              Specify this switch to show ALL of the changelog entries
      --version string   Sepecify the version to display the changelog for

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## delete

```plaintext
EXAMPLE:
  Delete a running POD.  This will prompt with a list of PODs that are running
  and were launched using ktrouble.

    > ktrouble delete

Usage:
  ktrouble delete [flags]

Flags:
  -a, --all   Choose from a list of running PODs from ALL users

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## diff

```plaintext
EXAMPLE:
  The 'diff' command will list the differences between your local 'config.yaml'
  file 'utilities' definitions and the remote repository.

  > ktrouble diff

Usage:
  ktrouble diff [flags]

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## edit

```plaintext
EXAMPLE:

    > ktrouble edit config --help
    > ktrouble edit template --help

Usage:
  ktrouble edit [flags]
  ktrouble edit [command]

Available Commands:
  config      Edit the default config, or specified in KTROUBLE_CONFIG
  template    Edit the default template, or specified one via --template/-t

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")

Use "ktrouble edit [command] --help" for more information about a command.
```

[TOC](#TOC)

## edit config

```plaintext
EXAMPLE:
  The default config can be hand edited
  
    > ktrouble edit config

EXAMPLE:
  Edit a secondary NON default config file

    > export KTROUBLE_CONFIG=cmaahs-config.yaml
    > ktrouble edit config

Usage:
  ktrouble edit config [flags]

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## edit template

```plaintext
EXAMPLE:
  
    > ktrouble edit template --template christmas

Usage:
  ktrouble edit template [flags]

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## fields

```plaintext
EXAMPLE:
  The 'fields' command will list the valid fields that can be used with various
  commands that accept the --fields/-f parameter.

    > ktrouble fields

Usage:
  ktrouble fields [flags]

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## get

```plaintext
EXAMPLE:
  These are mostly utility commands to review things important to running ktrouble.
  Allowing a display of various items stored in config.yaml and listing various
  kubernetes resources.

    > ktrouble get configs --help
    > ktrouble get namespaces --help
    > ktrouble get node --help
    > ktrouble get nodelabels --help
    > ktrouble get running --help
    > ktrouble get serviceaccount --help
    > ktrouble get sizes --help
    > ktrouble get templates --help
    > ktrouble get utilities --help

EXAMPLE:
  Get a list of PODs that are currently running on the current context kubernetes
  cluster that were created with the ktrouble utility.  If the 'enableBashLinks'
  config.yaml setting is 'true', a '<bash: ... >' command will be displayed,
  otherwise the SHELL path will be displayed.

    > ktrouble get pods

      NAME                NAMESPACE       STATUS   EXEC
      basic-tools-e1df2f  common-tooling  Running  <bash:kubectl -n common-tooling exec -it basic-tools-e1df2f -- /bin/bash>

      NAME                NAMESPACE       STATUS   SHELL
      basic-tools-e1df2f  common-tooling  Running  /bin/bash

Usage:
  ktrouble get [flags]
  ktrouble get [command]

Available Commands:
  configs        Get a list of configs
  namespace      Get a list of namespaces
  node           Get a list of node labels
  nodelabels     Get a list of defined node labels in config.yaml
  running        Get a list of running pods
  serviceaccount Get a list of K8s ServiceAccount(s) in a Namespace
  sizes          Get a list of defined sizes
  templates      Get a list of templates
  utilities      Get a list of supported utility container images

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")

Use "ktrouble get [command] --help" for more information about a command.
```

[TOC](#TOC)

## get configs

```plaintext
EXAMPLE:
  The ktrouble utility can support multiple config files, either with the
  '--config <config path>' or by setting the environment variable
  'KTROUBLE_CONFIG' to just the name of the config file, which will need to
  reside in the '$HOME/.config/ktrouble' directory

    > ktrouble get configs

      CONFIG
      alteryx-config.yaml
      cmaahs-config.yaml
      config.yaml

    > export KTROUBLE_CONFIG=cmaahs-config.yaml

Usage:
  ktrouble get configs [flags]

Aliases:
  configs, size, requests, request, limit, limits

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## get namespace

```plaintext
EXAMPLE:
  Return a list of kubernetes namespaces for the current context cluster

    > ktrouble get ns

Usage:
  ktrouble get namespace [flags]

Aliases:
  namespace, namespaces, ns

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## get node

```plaintext
EXAMPLE:
  Get a list of nodes for the current context cluster

    > ktrouble get node

Usage:
  ktrouble get node [flags]

Aliases:
  node, nodes

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## get nodelabels

```plaintext
EXAMPLE:
  Show the list of node labels in the configuration file

    > ktrouble get nodelabels

Usage:
  ktrouble get nodelabels [flags]

Aliases:
  nodelabels, nodelabel, nl, labels

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## get running

```plaintext
EXAMPLE:
  Get a list of PODs that are currently running on the current context kubernetes
  cluster that were created with the ktrouble utility.  If the 'enableBashLinks'
  config.yaml setting is 'true', a '<bash: ... >' command will be displayed,
  otherwise the SHELL path will be displayed.

    > ktrouble get running

      NAME                NAMESPACE       STATUS   EXEC
      basic-tools-e1df2f  common-tooling  Running  <bash:kubectl -n common-tooling exec -it basic-tools-e1df2f -- /bin/bash>

      NAME                NAMESPACE       STATUS   SHELL
      basic-tools-e1df2f  common-tooling  Running  /bin/bash

EXAMPLE:
  You can use the subcommand 'pods' in place of 'running'

    > ktrouble get pods

Usage:
  ktrouble get running [flags]

Aliases:
  running, pods, pod

Flags:
  -a, --all   List running PODs from ALL users

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## get serviceaccount

```plaintext
EXAMPLE:
  Return a list of kubernetes service accounts for a namespace

    > ktrouble get serviceaccount -n myspace

EXAMPLE:
  If you do not specify a namespace with '-n <namespace>', you will be prompted
  to select one

    > ktrouble get sa

Usage:
  ktrouble get serviceaccount [flags]

Aliases:
  serviceaccount, serviceaccounts, sa

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## get sizes

```plaintext
EXAMPLE:
  Display a list of POD size options from the configuration file

    > ktrouble get sizes

Usage:
  ktrouble get sizes [flags]

Aliases:
  sizes, size, requests, request, limit, limits

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## get templates

```plaintext
EXAMPLE:
  The 'get templates' command will output a list of templates in the templates
  directory

    > ktrouble get templates

Usage:
  ktrouble get templates [flags]

Aliases:
  templates, size, requests, request, limit, limits

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## get utilities

```plaintext
EXAMPLE:
  Display a list of utilities defined in the configuration file

    > ktrouble get utilities

Usage:
  ktrouble get utilities [flags]

Aliases:
  utilities, utility, utils, util, container, containers, image, images

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## launch

```plaintext
EXAMPLE:
  Just running ktrouble launch will prompt for all the things required to run

    > ktrouble launch

EXAMPLE:
  As an option, ktrouble can prompt with a list of kubernetes secrets from the
  namespace selected.  The chosen secrets will all be mounted under '/secrets/'
  directory, with each key as a file inside a directory named for the secret.

    > ktrouble launch --secrets

EXAMPLE:
  TODO: add command line parameters that can be used to set all the options
  for launching a POD

    > ktrouble launch (...)

Usage:
  ktrouble launch [flags]

Aliases:
  launch, create, apply, pod, l

Flags:
      --configs   Use this switch to prompt to mount configmaps in the POD
      --secrets   Use this switch to prompt to mount secrets in the POD
      --volumes   Use this switch to prompt to mount volumes in the POD

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## pull

```plaintext
EXAMPLE:
  The 'pull' command will prompt to choose from a list of utilities that are
  missing from your local 'config.yaml' utility defintions.

    > ktrouble pull

EXAMPLE:
  Items that you have that are local, but have different setting, can be pulled,
  and overwritten by adding a '-a' switch to the command.

    > ktrouble pull -a

Usage:
  ktrouble pull [flags]

Flags:
  -a, --all   Specify --all to list locally modified definitions as pull selections

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## push

```plaintext
EXAMPLE:
  The 'push' command allows you to push your local utility definitions into a
  common repository in 'futurama/farnsworth/tools/ktrouble-utils'.  The command
  will prompt you to choose a list of utilities to push to the repository.
  Utilities marked 'exclude from push' will not appear on the selection list.

    > ktrouble push

Usage:
  ktrouble push [flags]

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## remove

```plaintext
EXAMPLE:
    > ktrouble remove utility --help

Usage:
  ktrouble remove [flags]
  ktrouble remove [command]

Available Commands:
  utility     Remove a utility from the config file, or HIDE it if it is an upstream definition

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")

Use "ktrouble remove [command] --help" for more information about a command.
```

[TOC](#TOC)

## remove utility

```plaintext
EXAMPLE:
  The 'remove utility' command will prompt to select a utility definition to
  remove from your local 'config.yaml' file

    > ktrouble remove utility

Usage:
  ktrouble remove utility [flags]

Flags:
  -u, --name string   Unique name for your utility pod

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## set

```plaintext
EXAMPLE:
    > ktrouble set config --help

Usage:
  ktrouble set [flags]
  ktrouble set [command]

Available Commands:
  config      Set configuration options for ktrouble

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")

Use "ktrouble set [command] --help" for more information about a command.
```

[TOC](#TOC)

## set config

```plaintext
EXAMPLE:
  If you store your git personal access token in an ENV variable, you can specify
  the variable name.

    > ktrouble set config --user christopher.maahs --tokenvar GLA_TOKEN

EXAMPLE:
  If you don't store your personal access token in an ENV variable, it can be
  stored directly in the config.yaml file.  Don't forgot to add a 'space' in
  front of running this next command so the token doesn't end up in your
  history file, if you have that option set in your shell

    > ktrouble set config --user christopher.maahs --token <your token>

EXAMPLE:
  If you want to point 'ktrouble' to a different git repository for upstream
  storage of utility pod definitions

    > ktrouble set config --giturl "https://github.com/cmaahs/ktrouble-utils.git"

EXAMPLE:
  If you would like 'ktrouble launch' to prompt for secret selection on every
  run, rather than just when you use the '--secrets' switch

    > ktrouble set config --secrets

EXAMPLE:
  If you would like 'ktrouble launch' to prompt for configmap selection on every
  run, rather than just when you use the '--configs' switch

    > ktrouble set config --configs

EXAMPLE:
  If you use dynamic hyperlinking in your terminal software, you can enable output
  with '<bash: >' decorations

    > ktrouble set config --bashlinks

Usage:
  ktrouble set config [flags]

Flags:
      --bashlinks         Toggle the use of Bash Links for iTerm2
      --configs           Toggle the Prompt for ConfigMaps default
      --giturl string     Set the URL for the repository for upstream utils
      --secrets           Toggle the Prompt for Secrets default
      --token string      Set your git personal token
      --tokenvar string   Set the name of the ENV VAR that contains your git personal token
  -u, --user string       Set your git username

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## status

```plaintext
EXAMPLE:
  The 'status' command will list the disposition of your local 'config.yaml'
  file 'utilities' definitions against the 'futurama/farnsworth/tools/ktrouble-utils'
  repostory.

    > ktrouble status

Usage:
  ktrouble status [flags]

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## update

```plaintext
EXAMPLE:

    > ktrouble update utility --help

Usage:
  ktrouble update [flags]
  ktrouble update [command]

Aliases:
  update, modify

Available Commands:
  utility     Update an existing utility pod definition in the local config.yaml file

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")

Use "ktrouble update [command] --help" for more information about a command.
```

[TOC](#TOC)

## update utility

```plaintext
EXAMPLE:
  Toggle the 'exclude from push' flag for a utility definition.

    > ktrouble update utility -u helm-kubectl311 --toggle-exclude

EXAMPLE:
  Toggle the 'hidden' flag for an existing utility pod definition

    > ktrouble update utility -u alpine3 --toggle-hidden

EXAMPLE:
  Change the 'command' the utility will run

    > ktrouble update utility -u helm-kubectl311 -c '/bin/sh'

Usage:
  ktrouble update utility [flags]

Flags:
  -c, --cmd string          Default shell/command to use when 'exec'ing into the POD
  -u, --name string         Unique name for your utility pod
  -r, --repository string   Repository and tag for your utility container, eg: cmaahs/basic-tools:latest
  -e, --toggle-exclude      Switch the current 'excludeFromShare' flag for the utility definition
      --toggle-hidden       Switch the current 'hidden' flag for the utility definition

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)

## version

```plaintext
EXAMPLE: 
    > ktrouble version

Usage:
  ktrouble version [flags]

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
  -t, --template string    Specify the template file to use to render the POD manifest (default "default")
```

[TOC](#TOC)
