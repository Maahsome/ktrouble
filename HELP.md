# ktrouble help for all commands

## TOC

- [_main_](#ktrouble)
- [add](#add)
- [add utility](#add-utility)
- [delete](#delete)
- [fields](#fields)
- [get](#get)
- [get namespace](#get-namespace)
- [get node](#get-node)
- [get nodelabels](#get-nodelabels)
- [get running](#get-running)
- [get serviceaccount](#get-serviceaccount)
- [get sizes](#get-sizes)
- [get utilities](#get-utilities)
- [launch](#launch)
- [pull](#pull)
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
  add         
  completion  Generate the autocompletion script for the specified shell
  delete      Delete PODs that have been created by ktrouble
  fields      Display a list of valid fields to use with the --fields/-f parameter for each command
  genhelp     Output help from all the sub commands
  get         Get various internal configuration and kubernetes resource listings
  help        Help about any command
  launch      launch a kubernetes troubleshooting pod
  pull        Pull utility definitions from git
  remove      
  set         
  status      Get a comparison of the local utility definitions with the upstream one
  update      
  version     Express the 'version' of ktrouble.

Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'

Use "ktrouble [command] --help" for more information about a command.
```

[TOC](#TOC)

## add

```plaintext
EXAMPLES
	ktrouble add utility

Usage:
  ktrouble add [flags]
  ktrouble add [command]

Available Commands:
  utility     

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'

Use "ktrouble add [command] --help" for more information about a command.
```

[TOC](#TOC)

## add utility

```plaintext
Usage:
  ktrouble add utility [flags]

Flags:
  -c, --cmd string          Default shell/command to use when 'exec'ing into the POD (default "/bin/sh")
  -e, --exclude             Exclude from 'push' to central repository
  -u, --name string         Unique name for your utility pod
  -r, --repository string   Repository and tag for your utility container, eg: cmaahs/basic-tools:latest

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
```

[TOC](#TOC)

## delete

```plaintext
EXAMPLE:
  Delete a running POD.  This will prompt with a list of PODs that are running
  and were launched using ktrouble

  > ktrouble delete

Usage:
  ktrouble delete [flags]

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
```

[TOC](#TOC)

## fields

```plaintext
Display a list of valid fields to use with the --fields/-f parameter for each command

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
```

[TOC](#TOC)

## get

```plaintext
EXAMPLE:
  These are mostly utility commands to review things important to running ktrouble.
  Allowing a display of various items stored in config.yaml and listing various
  kubernetes resources.

  > ktrouble get namespaces
  > ktrouble get utilities

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
  namespace      Get a list of namespaces
  node           Get a list of node labels
  nodelabels     Get a list of defined node labels in config.yaml
  running        Get a list of running pods
  serviceaccount Get a list of K8s ServiceAccount(s) in a Namespace
  sizes          Get a list of defined sizes
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

Use "ktrouble get [command] --help" for more information about a command.
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

Usage:
  ktrouble get running [flags]

Aliases:
  running, pods, pod

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
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
```

[TOC](#TOC)

## launch

```plaintext
EXAMPLE:
  Just running kubectl launch will prompt for all the things required to run

  > kubectl launch

EXAMPLE:
  TODO: add command line parameters that can be used to set all the options
  for launching a POD

  > kubectl launch (...)

Usage:
  ktrouble launch [flags]

Aliases:
  launch, create, apply, pod, l

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
```

[TOC](#TOC)

## pull

```plaintext
EXAMPLE:
  > ktrouble pull

Usage:
  ktrouble pull [flags]

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'
```

[TOC](#TOC)

## remove

```plaintext
EXAMPLES
	ktrouble remove utility

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

Use "ktrouble remove [command] --help" for more information about a command.
```

[TOC](#TOC)

## remove utility

```plaintext
Remove a utility from the config file, or HIDE it if it is an upstream definition

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
```

[TOC](#TOC)

## set

```plaintext
EXAMPLES
	ktrouble set gituser

Usage:
  ktrouble set [flags]
  ktrouble set [command]

Available Commands:
  config      Set git configuration options

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'

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

Usage:
  ktrouble set config [flags]

Flags:
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
```

[TOC](#TOC)

## status

```plaintext
EXAMPLE:
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
```

[TOC](#TOC)

## update

```plaintext
EXAMPLES
	ktrouble update utility

Usage:
  ktrouble update [flags]
  ktrouble update [command]

Available Commands:
  utility     

Global Flags:
      --config string      config file (default is $HOME/.splicectl/config.yml)
  -f, --fields strings     Specify an array of field names: eg, --fields 'NAME,REPOSITORY'
      --log-file string    Set the logging level: trace,debug,info,warning,error,fatal
  -v, --log-level string   Set the logging level: trace,debug,info,warning,error,fatal
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
      --no-headers         Suppress header output in Text output
  -o, --output string      output types: json, text, yaml, gron, raw
  -s, --show-hidden        Show entries with the 'hidden' property set to 'true'

Use "ktrouble update [command] --help" for more information about a command.
```

[TOC](#TOC)

## update utility

```plaintext
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
```

[TOC](#TOC)

## version

```plaintext
Express the 'version' of ktrouble.

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
```

[TOC](#TOC)

