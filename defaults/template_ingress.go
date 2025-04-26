package defaults

func DefaultIngressTemplate() string {
	defaultTemplate := `---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    ingress.kubernetes.io/service-upstream: "true"
    kubernetes.io/ingress.class: nginx
    nginx.ingress.kubernetes.io/proxy-read-timeout: "3600"
    nginx.ingress.kubernetes.io/proxy-send-timeout: "3600"
  labels:
    app: ktrouble
    launchedby: {{ $.Parameters.launchedby }}
    associatedPod: {{ $.Parameters.associatedPod }}
  name: {{ $.Parameters.name }}
  namespace: {{ $.Parameters.namespace }}
spec:
  rules:
  - host: {{ $.Parameters.host }}
    http:
      paths:
      - backend:
          service:
            name: {{ $.Parameters.name }}
            port:
              number: {{ $.Parameters.targetPort }}
        path: "/{{ $.Parameters.path }}/"
        pathType: Prefix
`

	return defaultTemplate
}
