package edit

import (
	"bytes"
	"fmt"
	"ktrouble/common"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/cmd/util/editor"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Edit the default config, or specified in KTROUBLE_CONFIG",
	Long: `EXAMPLE
  > ktrouble edit config
`,
	Run: func(cmd *cobra.Command, args []string) {
		editConfig()
	},
}

func editConfig() {
	edit := editor.NewDefaultEditor([]string{
		"KTROUBLE_EDITOR",
		"EDITOR",
	})
	home, herr := os.UserHomeDir()
	if herr != nil {
		common.Logger.WithError(herr).Error("failed to determine the HOME directory")
	}
	confDir := fmt.Sprintf("%s/.config/ktrouble", home)
	fileToOpen := fmt.Sprintf("%s/config.yaml", confDir)
	envCfgFile := os.Getenv("KTROUBLE_CONFIG")
	if envCfgFile != "" {
		fileToOpen = fmt.Sprintf("%s/%s", confDir, envCfgFile)
	}

	_, buffer := openFile(fileToOpen)

	original := buffer.Bytes()

	edited, _, err := edit.LaunchTempFile("ktrouble-config", ".yaml", buffer)
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

func init() {
	editCmd.AddCommand(configCmd)
}
