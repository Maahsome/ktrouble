package cmd

import "os"

func determineNamespace() string {

	namespace := ""
	if len(os.Getenv("NAMESPACE")) > 0 {
		namespace = os.Getenv("NAMESPACE")
		if len(c.Namespace) > 0 {
			namespace = c.Namespace
		}
	} else {
		if len(c.Namespace) > 0 {
			namespace = c.Namespace
		}
	}

	if namespace == "" {
		nssList := getNamespaces()
		namespace = askForNamespace(nssList)
	}
	return namespace
}
