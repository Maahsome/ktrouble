package objects

import "strings"

func fieldsToUpper(fields []string) []string {
	upperFields := make([]string, len(fields))
	for i, field := range fields {
		upperFields[i] = strings.ToUpper(field)
	}
	return upperFields
}
