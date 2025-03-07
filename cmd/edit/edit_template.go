package edit

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"ktrouble/common"
	"log"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/cmd/util/editor"
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: editTemplateHelp.Short(),
	Long: editTemplateHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		editTemplate()
	},
}

func editTemplate() {
	edit := editor.NewDefaultEditor([]string{
		"KTROUBLE_EDITOR",
		"EDITOR",
	})
	home, herr := os.UserHomeDir()
	if herr != nil {
		common.Logger.WithError(herr).Error("failed to determine the HOME directory")
	}
	tmplDir := fmt.Sprintf("%s/.config/ktrouble/templates", home)
	fileToOpen := fmt.Sprintf("%s/%s", tmplDir, c.TemplateFile)

	_, buffer := openFile(fileToOpen)

	original := buffer.Bytes()

	edited, _, err := edit.LaunchTempFile("ktrouble-template", ".tpl", buffer)
	if err != nil {
		common.Logger.WithError(err).Error("failed to exit the editor cleanly")
	}

	if bytes.Equal(edited, original) {
		common.Logger.Info("no changes detected")
	} else {
		err := os.WriteFile(fileToOpen, edited, 0644)
		if err != nil {
			common.Logger.WithError(err).Error("failed to write changes")
		} else {
			common.Logger.Info("changes saved")
		}
	}

}

func openFile(name string) (byteCount int, buffer *bytes.Buffer) {

	var (
		data  *os.File
		part  []byte
		err   error
		count int
	)

	data, err = os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	reader := bufio.NewReader(data)
	buffer = bytes.NewBuffer(make([]byte, 0))
	part = make([]byte, chunksize)

	for {
		if count, err = reader.Read(part); err != nil {
			break
		}
		buffer.Write(part[:count])
	}
	if err != io.EOF {
		common.Logger.WithError(err).Fatalf("Error Reading %s", name)
	} else {
		err = nil
	}

	byteCount = buffer.Len()
	return
}

func init() {
	editCmd.AddCommand(templateCmd)
}
