package defaults

import (
	"ktrouble/objects"
	"strings"
)

func MissingConfigOutputFieldDefinitions(confDefs map[string][]string) []string {
	missing := make([]string, 0)
	for k, v := range OutputNamesList() {
		if _, ok := confDefs[k]; !ok {
			missing = append(missing, k)
		} else {
			if len(v) == 0 {
				missing = append(missing, k)
			}
		}
	}
	return missing
}

func OutputNamesList() map[string]string {
	outputs := OutputFieldsList()
	names := make(map[string]string, len(outputs))
	for _, output := range outputs {
		names[output.Name] = output.Fields
	}
	return names
}

func ValidFieldNamesForOutputName(name string) (map[string]string, string) {
	validFields := ValidOutputFields()
	if fields, ok := validFields[name]; ok {
		fieldList := strings.Split(fields, ",")
		fieldMap := make(map[string]string, len(fieldList))
		for _, field := range fieldList {
			fieldMap[strings.TrimSpace(field)] = field
		}
		return fieldMap, fields
	}
	return nil, ""
}

func ValidOutputFields() map[string]string {
	return map[string]string{
		"environments":    "NAME, REPOSITORY, EXCLUDED, HIDDEN, REMOVE_UPSTREAM",
		"ephemeral_sleep": "NAME, SECONDS",
		"ingress":         "NAME, NAMESPACE, CLASS, HOSTS, ADDRESS, PORTS, LAUNCHED_BY",
		"namespace":       "NAMESPACE",
		"node_labels":     "LABEL",
		"node":            "NODE",
		"output_fields":   "NAME, FIELDS",
		"pod":             "NAME, NAMESPACE, STATUS, LAUNCHED_BY, UTILITY, SHELL/SERVICE",
		"service_account": "SERVICE_ACCOUNT",
		"service":         "NAME, NAMESPACE, TYPE, CLUSTER_IP, EXTERNAL_IP, PORTS, LAUNCHED_BY",
		"size":            "NAME, CPU_LIMIT, MEM_LIMIT, CPU_REQUEST, MEM_REQUEST",
		"status":          "NAME, STATUS, EXCLUDE",
		"utility":         "NAME, REPOSITORY, EXEC, HIDDEN, EXCLUDED, SOURCE, ENVIRONMENTS, REQUIRECONFIGMAPS, REQUIRESECRETS, HINT, REMOVE_UPSTREAM",
		"version":         "SEMVER, BUILD_DATE, GIT_COMMIT, GIT_REF",
	}
}

func OutputFieldsList() []objects.OutputFields {

	return []objects.OutputFields{
		{
			Name:   "environments",
			Fields: "NAME,REPOSITORY,EXCLUDED",
		},
		{
			Name:   "ephemeral_sleep",
			Fields: "NAME,SECONDS",
		},
		{
			Name:   "ingress",
			Fields: "NAME,NAMESPACE,CLASS,URL,ADDRESS,PORTS,LAUNCHED_BY",
		},
		{
			Name:   "namespace",
			Fields: "NAMESPACE",
		},
		{
			Name:   "node_labels",
			Fields: "LABEL",
		},
		{
			Name:   "node",
			Fields: "NODE",
		},
		{
			Name:   "output_fields",
			Fields: "NAME,FIELDS",
		},
		{
			Name:   "pod",
			Fields: "NAME,NAMESPACE,STATUS,LAUNCHED_BY,UTILITY,SHELL/SERVICE",
		},
		{
			Name:   "service_account",
			Fields: "SERVICE_ACCOUNT",
		},
		{
			Name:   "service",
			Fields: "NAME,NAMESPACE,TYPE,CLUSTER_IP,EXTERNAL_IP,PORTS,LAUNCHED_BY",
		},
		{
			Name:   "size",
			Fields: "NAME,CPU_LIMIT,MEM_LIMIT,CPU_REQUEST,MEM_REQUEST",
		},
		{
			Name:   "status",
			Fields: "NAME,STATUS,PUSH_EXCLUDE",
		},
		{
			Name:   "utility",
			Fields: "NAME,REPOSITORY,EXEC",
		},
		{
			Name:   "version",
			Fields: "SEMVER,BUILD_DATE,GIT_COMMIT,GIT_REF",
		},
	}
}
