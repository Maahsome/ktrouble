package update

import (
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
	Short:   updateSizeHelp.Short(),
	Long:    updateSizeHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		if checkUpdateSizeParams() {
			s, err := updateSize()
			if err != nil {
				logrus.WithError(err).Error("Failed to update the size definition")
			}
			if !c.FormatOverridden {
				c.OutputFormat = "text"
			}
			updated := objects.ResourceSizeList{}
			updated = append(updated, s)
			c.OutputData(&updated, objects.TextOptions{
				NoHeaders:     c.NoHeaders,
				Fields:        c.Fields,
				DefaultFields: c.OutputFieldsMap["size"],
			})
		} else {
			common.Logger.Error("Parameters passed in have failed checks.  Please review the warnings above")
		}
	},
}

func checkUpdateSizeParams() bool {
	allParamsSet := true

	if len(sizeParam.Name) == 0 {
		sizeParam.Name = ask.PromptForResourceSize(c.SizeDefs)
	}

	if len(sizeParam.Name) == 0 {
		allParamsSet = false
		common.Logger.Warn("No size selected")
	}

	if len(sizeParam.LimitsCPU) == 0 && len(sizeParam.LimitsMEM) == 0 && len(sizeParam.RequestCPU) == 0 && len(sizeParam.RequestMEM) == 0 {
		allParamsSet = false
		common.Logger.Warn("At least one value must be provided to update: --limitscpu, --limitsmem, --requestcpu, or --requestmem")
	}

	if allParamsSet {
		found := false
		for _, v := range c.SizeDefs {
			if sizeParam.Name == v.Name {
				found = true
				break
			}
		}
		if !found {
			allParamsSet = false
			common.Logger.Warnf("The --name size name does not match an existing size: %s", sizeParam.Name)
		}
	}

	return allParamsSet
}

func updateSize() (objects.ResourceSize, error) {
	updatedSize := false
	updatedDef := objects.ResourceSize{}

	for i, v := range c.SizeDefs {
		if sizeParam.Name == v.Name {
			if len(sizeParam.LimitsCPU) > 0 {
				c.SizeDefs[i].LimitsCPU = sizeParam.LimitsCPU
				updatedSize = true
			}
			if len(sizeParam.LimitsMEM) > 0 {
				c.SizeDefs[i].LimitsMEM = sizeParam.LimitsMEM
				updatedSize = true
			}
			if len(sizeParam.RequestCPU) > 0 {
				c.SizeDefs[i].RequestCPU = sizeParam.RequestCPU
				updatedSize = true
			}
			if len(sizeParam.RequestMEM) > 0 {
				c.SizeDefs[i].RequestMEM = sizeParam.RequestMEM
				updatedSize = true
			}
			updatedDef = c.SizeDefs[i]
			break
		}
	}

	if updatedSize {
		viper.Set("resourceSizing", c.SizeDefs)
		verr := viper.WriteConfig()
		if verr != nil {
			common.Logger.WithError(verr).Info("Failed to write config")
			return updatedDef, verr
		}
		c.SizeMap[updatedDef.Name] = updatedDef
	} else {
		common.Logger.Warnf("The size, %s, was not updated, perhaps a mis-matched name", sizeParam.Name)
	}

	return updatedDef, nil
}

func init() {
	updateCmd.AddCommand(sizeCmd)
	sizeCmd.Flags().StringVar(&sizeParam.Name, "name", "", "Name of the size definition to update")
	sizeCmd.Flags().StringVar(&sizeParam.LimitsCPU, "limitscpu", "", "CPU limit for the size definition, eg: 250m")
	sizeCmd.Flags().StringVar(&sizeParam.LimitsMEM, "limitsmem", "", "Memory limit for the size definition, eg: 2Gi")
	sizeCmd.Flags().StringVar(&sizeParam.RequestCPU, "requestcpu", "", "CPU request for the size definition, eg: 100m")
	sizeCmd.Flags().StringVar(&sizeParam.RequestMEM, "requestmem", "", "Memory request for the size definition, eg: 512Mi")
}
