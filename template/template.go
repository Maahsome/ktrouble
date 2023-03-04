package template

import (
	"strings"
	"text/template"
)

// takes in some text and pads it with spaces even if it is multiline
func indent(spaces int, v string) string {
	pad := strings.Repeat(" ", spaces)

	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}

var applicationsTemplateFuncs = template.FuncMap{
	"indent":    indent,
	"hasPrefix": strings.HasPrefix,
}

var ApplicationsTemplate = template.Must(
	template.New("applications.yaml.tpl").Funcs(applicationsTemplateFuncs).Parse(
		`---
apiVersion: v1
kind: Pod
metadata:
  name: {{ $.Parameters.name }}
  namespace: {{ $.Parameters.namespace }}
  labels:
    app: ktrouble
spec:
  containers:
  - name: {{ $.Parameters.name}}
    image: {{ $.Parameters.registry }}
    command:
      - sleep
      - "86400"
    imagePullPolicy: Always
    resources:
      limits:
        cpu: {{ $.Parameters.limitsCpu }}
        memory: {{ $.Parameters.limitsMem }}
      requests:
        cpu: {{ $.Parameters.requestCpu }}
        memory: {{ $.Parameters.requestMem }}
  {{- if eq $.Parameters.hasSelector "true" }}
  nodeSelector:
    {{ $.Parameters.selector }}
  {{- end }}
  serviceAccount: {{ $.Parameters.serviceAccount}}
  restartPolicy: Always
`))
