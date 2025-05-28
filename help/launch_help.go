package help

import (
	"fmt"

	"github.com/fatih/color"
)

type LaunchCmd struct {
}

func (l *LaunchCmd) Short() string {
	return "Launch a kubernetes troubleshooting pod"
}

func (l *LaunchCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Just running ktrouble launch will prompt for all the things required to run
`

	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble launch`))

	longText += `EXAMPLE:
  As an option, ktrouble can prompt with a list of kubernetes secrets from the
  namespace selected.  The chosen secrets will all be mounted under '/secrets/'
  directory, with each key as a file inside a directory named for the secret.
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble launch --prompt-secrets`))

	longText += `EXAMPLE:
  Launch a container that has nginx setup to consume two environment variables;
  'APPLICATION_BASE_PATH' and 'LISTEN_PORT'.  The 'APPLICATION_BASE_PATH' is the
  path that the application is served from, and 'LISTEN_PORT' is the port that
  the application listens on.  This will also create a service and ingress for
  the POD.  The host and path for the ingress can be specified with the --host
  and --path flags.  The port that the POD listens on can be specified with the
  --port flag.
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble launch --port 8080 --host myservice.example.com --path service-myservice --ingress`))

	longText += `EXAMPLE:
  All of the above examples prompt for all the missing parameters.  You can also specify ALL of the
  parameters on the command line, and optionally just return the POD name.

  All of these parameters, except or node-selector, need to be set if you want to suppress the prompts.

  Parameters:
    - --utility/-u <name>           : The name of the utility to launch, must match the utility name
                                    : be sure to specify the "environment" name if the utility
                                    : has multiple environments, eg: --utility 'uppers/dns-tools'
    - --namespace/-n <name>         : The namespace to use
    - --service-account <name>      : The name of the service account to use
    - --node-selector <label/value> : The node selector to use
                                    : The label/value pair must be inside single quotes, eg:
                                    : --node-selector '"kubernetes.io/arch": "amd64"'
                                    : MUST specify '-none-' to suppress the prompt
    - --secrets '<name>,<name>'     : The secret names to mount, comma separated
    - --configmaps '<name>,<name>'  : The configmap names to mount, comma separated
    - --size <name>                 : The size of the POD to use, must match a size name, ktrouble get sizes
    - --output-name                 : Use this boolean switch to just return the name of the POD
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble launch (...)`))

	return longText
}
