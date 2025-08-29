package set

import (
	"fmt"
	"ktrouble/defaults"
	"sort"
	"strings"

	"github.com/fatih/color"
)

type SetOutputFieldsCmd struct {
}

func (s *SetOutputFieldsCmd) Short() string {
	return "Set default output fields for ktrouble commands"
}

func (s *SetOutputFieldsCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Set the default output field for 'environment' commands to 'name,image,hidden'.
`

	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble set output-fields --output-name environments --fields 'name,image,excluded,hidden'`))

	longText += `EXAMPLE:
  Set the default output for 'pod' commands to 'name,launched_by'
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble set output-fields --output-name pod --fields 'name,launched_by'`))

	longText += `EXAMPLE:
  Reset to the default output fields for 'pod' commands.
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble set output-fields --output-name pod --fields ''`))

	longText += `EXAMPLE:
  The possible --output-name and --field combinations are:

  Output Names/Fields:

`
	validOutputFields := defaults.ValidOutputFields()
	outputNames := make([]string, 0, len(validOutputFields))
	for k := range validOutputFields {
		outputNames = append(outputNames, k)
	}
	sort.Strings(outputNames)
	// I would like to output the validOutputFields 5 at a time, with the first output
	// containing the 'v' of outputNames and all the rest would be indented 4 spaces plus the
	// lenggth of the 'v' string.
	// for _, v := range outputNames {
	// 	longText += fmt.Sprintf("    - %s: %s\n", v, validOutputFields[v])
	// }
	maxKeyLen := 0
	for _, v := range outputNames {
		if len(v) > maxKeyLen {
			maxKeyLen = len(v)
		}
	}

	for _, v := range outputNames {
		items := strings.Split(validOutputFields[v], ",")
		for i := 0; i < len(items); i += 5 {
			end := i + 5
			if end > len(items) {
				end = len(items)
			}
			chunk := items[i:end]
			line := strings.Join(chunk, ", ")

			if i == 0 {
				// Pad the key so all colons align
				paddedKey := fmt.Sprintf("%-*s", maxKeyLen, v)
				longText += fmt.Sprintf("    - %s: %s\n", paddedKey, line)
			} else {
				// Continuation lines with matching indentation
				indent := strings.Repeat(" ", 7+maxKeyLen)
				longText += fmt.Sprintf("%s%s\n", indent, line)
			}
		}
	}

	return longText
}
