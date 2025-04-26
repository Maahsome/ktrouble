package defaults

func DefaultServiceTemplate() string {
	defaultTemplate := `---
apiVersion: v1
kind: Service
metadata:
  name: {{ $.Parameters.name }}
  namespace: {{ $.Parameters.namespace }}
  labels:
    app: ktrouble
    launchedby: {{ $.Parameters.launchedby }}
    associatedPod: {{ $.Parameters.associatedPod }}
  annotations:
spec:
  type: "ClusterIP"
  ports:
  - name: app
    port: 443
    protocol: TCP
    targetPort: {{ $.Parameters.targetPort }}
  selector:
    app.kubernetes.io/name: {{ $.Parameters.name }}
`

	return defaultTemplate
}
