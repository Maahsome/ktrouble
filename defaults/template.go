package defaults

func DefaultTemplate() string {
	defaultTemplate := `---
apiVersion: v1
kind: Pod
metadata:
  name: {{ $.Parameters.name }}
  namespace: {{ $.Parameters.namespace }}
  labels:
    app: ktrouble
    app.kubernetes.io/name: {{ $.Parameters.name }}
    launchedby: {{ $.Parameters.launchedby }}
spec:
  containers:
  - name: {{ $.Parameters.name }}
    image: {{ $.Parameters.image }}:{{ $.Parameters.tag }}
    {{- if eq $.Parameters.ingressEnabled "false" }}
    command:
      - sleep
      - "86400"
    {{- end }}
    env:
    - name: APPLICATION_BASE_PATH
      value: {{ $.Parameters.path }}
    - name: LISTEN_PORT
      value: {{ $.Parameters.targetPort }}
    imagePullPolicy: Always
    resources:
      limits:
        cpu: {{ $.Parameters.limitsCpu }}
        memory: {{ $.Parameters.limitsMem }}
      requests:
        cpu: {{ $.Parameters.requestCpu }}
        memory: {{ $.Parameters.requestMem }}
    {{- if or $.Secrets $.ConfigMaps }}
    volumeMounts:
    {{- end }}
    {{- if $.Secrets }}
    {{- range $.Secrets }}
    - mountPath: "/secrets/{{ . }}"
      name: {{ printf "ktrouble-%.53s" . }}
      readOnly: true
    {{- end }}
    {{- end }}
    {{- if $.ConfigMaps }}
    {{- range $.ConfigMaps }}
    - mountPath: "/configmaps/{{ .}}"
      name: ktrouble-cm-{{ . }}
      readOnly: true
    {{- end }}
    {{- end }}
  {{- if eq $.Parameters.hasSelector "true" }}
  nodeSelector:
    {{ $.Parameters.selector }}
  {{- end }}
  serviceAccount: {{ $.Parameters.serviceAccount}}
  restartPolicy: Always
  {{- if or $.Secrets $.ConfigMaps }}
  volumes:
  {{- end }}
  {{- if $.Secrets }}
  {{- range $.Secrets }}
  - name: {{ printf "ktrouble-%.53s" . }}
    secret:
      defaultMode: 420
      secretName: {{ . }}
  {{- end }}
  {{- end }}
  {{- if $.ConfigMaps }}
  {{- range $.ConfigMaps }}
  - name: ktrouble-cm-{{ .}}
    configMap:
      defaultMode: 420
      name: {{ . }}
  {{- end }}
  {{- end }}
`

	return defaultTemplate
}
