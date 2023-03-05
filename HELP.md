# ktrouble help for all commands

## TOC

- [_main_](#ktrouble)
- [delete](#delete)
- [get](#get)
- [get namespace](#get-namespace)
- [get node](#get-node)
- [get running](#get-running)
- [get serviceaccount](#get-serviceaccount)
- [get sizes](#get-sizes)
- [get utilities](#get-utilities)
- [launch](#launch)
- [version](#version)

## ktrouble

```plaintext
EXAMPLE:

  TODO: add description

  > ktrouble

Usage:
  ktrouble [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  delete      Delete PODs that have been created by ktrouble
  genhelp     Output help from all the sub commands
  get         Get various resource lists
  help        Help about any command
  launch      launch a kubernetes troubleshooting pod
  version     Express the 'version' of ktrouble.

Flags:
      --config string      config file (default is $HOME/.ktrouble.yaml)
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
  -o, --output string      Set an output format: json, text, yaml, gron, md

Use "ktrouble [command] --help" for more information about a command.
```

[TOC](#TOC)

## delete

```plaintext
EXAMPLE:
	> ktrouble delete

Usage:
  ktrouble delete [flags]

Global Flags:
      --config string      config file (default is $HOME/.ktrouble.yaml)
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
  -o, --output string      Set an output format: json, text, yaml, gron, md
```

[TOC](#TOC)

## get

```plaintext
EXAMPLE:

Usage:
  ktrouble get [command]

Available Commands:
  namespace      Get a list of namespaces
  node           Get a list of node labels
  running        Get a list of running pods
  serviceaccount Get a list of K8s ServiceAccount(s) in a Namespace
  sizes          Get a list of defined sizes
  utilities      Get a list of supported utility container images

Global Flags:
      --config string      config file (default is $HOME/.ktrouble.yaml)
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
  -o, --output string      Set an output format: json, text, yaml, gron, md

Use "ktrouble get [command] --help" for more information about a command.
```

[TOC](#TOC)

## get namespace

```plaintext
EXAMPLE:

Usage:
  ktrouble get namespace [flags]

Aliases:
  namespace, namespaces, ns

Global Flags:
      --config string      config file (default is $HOME/.ktrouble.yaml)
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
  -o, --output string      Set an output format: json, text, yaml, gron, md
```

[TOC](#TOC)

## get node

```plaintext
EXAMPLE:
	> ktrouble get node

Usage:
  ktrouble get node [flags]

Aliases:
  node, nodes

Global Flags:
      --config string      config file (default is $HOME/.ktrouble.yaml)
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
  -o, --output string      Set an output format: json, text, yaml, gron, md
```

[TOC](#TOC)

## get running

```plaintext
EXAMPLE:
	> ktrouble get running

Usage:
  ktrouble get running [flags]

Aliases:
  running, pods, pod

Global Flags:
      --config string      config file (default is $HOME/.ktrouble.yaml)
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
  -o, --output string      Set an output format: json, text, yaml, gron, md
```

[TOC](#TOC)

## get serviceaccount

```plaintext
EXAMPLE:
	> ktrouble get serviceaccount -n myspace

Usage:
  ktrouble get serviceaccount [flags]

Aliases:
  serviceaccount, serviceaccounts, sa

Global Flags:
      --config string      config file (default is $HOME/.ktrouble.yaml)
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
  -o, --output string      Set an output format: json, text, yaml, gron, md
```

[TOC](#TOC)

## get sizes

```plaintext
EXAMPLE:
	> ktrouble get sizes

Usage:
  ktrouble get sizes [flags]

Aliases:
  sizes, size, requests, request, limit, limits

Global Flags:
      --config string      config file (default is $HOME/.ktrouble.yaml)
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
  -o, --output string      Set an output format: json, text, yaml, gron, md
```

[TOC](#TOC)

## get utilities

```plaintext
EXAMPLE:
	> ktrouble get utilities

Usage:
  ktrouble get utilities [flags]

Aliases:
  utilities, utility, util, container, containers, image, images

Global Flags:
      --config string      config file (default is $HOME/.ktrouble.yaml)
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
  -o, --output string      Set an output format: json, text, yaml, gron, md
```

[TOC](#TOC)

## launch

```plaintext
EXAMPLE:

Usage:
  ktrouble launch [flags]

Global Flags:
      --config string      config file (default is $HOME/.ktrouble.yaml)
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
  -o, --output string      Set an output format: json, text, yaml, gron, md
```

[TOC](#TOC)

## version

```plaintext
Express the 'version' of ktrouble.

Usage:
  ktrouble version [flags]

Global Flags:
      --config string      config file (default is $HOME/.ktrouble.yaml)
  -n, --namespace string   Specify the namespace to run in, ENV NAMESPACE then -n for preference
  -o, --output string      Set an output format: json, text, yaml, gron, md
```

[TOC](#TOC)

