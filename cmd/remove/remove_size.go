package remove

import (
	"errors"
	"ktrouble/ask"
	"ktrouble/common"
	"ktrouble/defaults"
	"ktrouble/objects"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// sizeParam is used to store command line parameters
var sizeParam objects.ResourceSize

// sizeCmd represents the size command
var sizeCmd = &cobra.Command{
	Use:     "size",
	Aliases: defaults.GetSizesAliases,
	Short:   removeSizeHelp.Short(),
	Long:    removeSizeHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		err := removeSize()
		if err != nil {
			logrus.WithError(err).Error("Failed to remove the size definition")
		}
		if !c.FormatOverridden {
			c.OutputFormat = "text"
		}
		c.OutputData(&c.SizeDefs, objects.TextOptions{
			NoHeaders:     c.NoHeaders,
			Fields:        c.Fields,
			DefaultFields: c.OutputFieldsMap["size"],
		})
	},
}

func removeSize() error {
	if len(sizeParam.Name) == 0 {
		sizeParam.Name = ask.PromptForResourceSize(c.SizeDefs)
	}

	if len(sizeParam.Name) == 0 {
		return errors.New("no size selected")
	}

	if len(c.SizeDefs) == 1 {
		return errors.New("cannot remove the last size definition; at least one size must remain")
	}

	removeIndex := -1
	for i, v := range c.SizeDefs {
		if sizeParam.Name == v.Name {
			removeIndex = i
			break
		}
	}

	if removeIndex == -1 {
		return errors.New("size name was not found in the existing size definitions")
	}

	c.SizeDefs = append(c.SizeDefs[:removeIndex], c.SizeDefs[removeIndex+1:]...)
	viper.Set("resourceSizing", c.SizeDefs)
	verr := viper.WriteConfig()
	if verr != nil {
		common.Logger.WithError(verr).Info("Failed to write config")
		return verr
	}

	delete(c.SizeMap, sizeParam.Name)
	return nil
}

func init() {
	removeCmd.AddCommand(sizeCmd)
	sizeCmd.Flags().StringVar(&sizeParam.Name, "name", "", "Name of the size definition to remove")
}
