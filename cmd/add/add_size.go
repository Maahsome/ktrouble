package add

import (
	"fmt"
	"ktrouble/common"
	"ktrouble/defaults"
	"ktrouble/objects"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// sizeParam is used to store command line parameters
var sizeParam objects.ResourceSize

// sizeCmd represents the size command
var sizeCmd = &cobra.Command{
	Use:     "size",
	Aliases: defaults.GetSizesAliases,
	Short:   addSizeHelp.Short(),
	Long:    addSizeHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		if checkAddSizeParams() {
			s, err := addSize()
			if err != nil {
				common.Logger.WithError(err).Error("Failed to add the size definition")
			}
			if !c.FormatOverridden {
				c.OutputFormat = "text"
			}
			added := objects.ResourceSizeList{}
			added = append(added, s)
			c.OutputData(&added, objects.TextOptions{
				NoHeaders:     c.NoHeaders,
				Fields:        c.Fields,
				DefaultFields: c.OutputFieldsMap["size"],
			})
		} else {
			common.Logger.Error("Parameters passed in have failed checks.  Please review the warnings above")
		}
	},
}

func checkAddSizeParams() bool {
	allParamsSet := true
	if len(sizeParam.Name) == 0 {
		allParamsSet = false
		common.Logger.Warn("The --name/-n parameter must be set")
	}
	if len(sizeParam.LimitsCPU) == 0 {
		allParamsSet = false
		common.Logger.Warn("The --limitscpu parameter must be set")
	}
	if len(sizeParam.LimitsMEM) == 0 {
		allParamsSet = false
		common.Logger.Warn("The --limitsmem parameter must be set")
	}
	if len(sizeParam.RequestCPU) == 0 {
		allParamsSet = false
		common.Logger.Warn("The --requestcpu parameter must be set")
	}
	if len(sizeParam.RequestMEM) == 0 {
		allParamsSet = false
		common.Logger.Warn("The --requestmem parameter must be set")
	}

	if allParamsSet {
		for _, v := range c.SizeDefs {
			if sizeParam.Name == v.Name {
				allParamsSet = false
				s := objects.ResourceSizeList{}
				s = append(s, v)
				common.Logger.Warn("The --name/-n size name clashes with an existing size name, please choose another, or use 'update size' to update the existing size")
				if !c.FormatOverridden {
					c.OutputFormat = "text"
				}
				c.OutputData(&s, objects.TextOptions{
					NoHeaders:     c.NoHeaders,
					Fields:        c.Fields,
					DefaultFields: c.OutputFieldsMap["size"],
				})
				break
			}
		}
	}

	return allParamsSet
}

func addSize() (objects.ResourceSize, error) {
	newSize := objects.ResourceSize{
		Name:       sizeParam.Name,
		LimitsCPU:  sizeParam.LimitsCPU,
		LimitsMEM:  sizeParam.LimitsMEM,
		RequestCPU: sizeParam.RequestCPU,
		RequestMEM: sizeParam.RequestMEM,
	}

	c.SizeDefs = append(c.SizeDefs, newSize)
	viper.Set("resourceSizing", c.SizeDefs)
	verr := viper.WriteConfig()
	if verr != nil {
		common.Logger.WithError(verr).Info("Failed to write config")
		return newSize, verr
	}

	c.SizeMap[newSize.Name] = newSize
	return newSize, nil
}

func init() {
	addCmd.AddCommand(sizeCmd)

	sizeCmd.Flags().StringVar(&sizeParam.Name, "name", "", "Unique name for your size definition")
	sizeCmd.Flags().StringVar(&sizeParam.LimitsCPU, "limitscpu", "", "CPU limit for the size definition, eg: 250m")
	sizeCmd.Flags().StringVar(&sizeParam.LimitsMEM, "limitsmem", "", "Memory limit for the size definition, eg: 2Gi")
	sizeCmd.Flags().StringVar(&sizeParam.RequestCPU, "requestcpu", "", "CPU request for the size definition, eg: 100m")
	sizeCmd.Flags().StringVar(&sizeParam.RequestMEM, "requestmem", "", "Memory request for the size definition, eg: 512Mi")

	_ = sizeCmd.MarkFlagRequired("name")
	_ = sizeCmd.MarkFlagRequired("limitscpu")
	_ = sizeCmd.MarkFlagRequired("limitsmem")
	_ = sizeCmd.MarkFlagRequired("requestcpu")
	_ = sizeCmd.MarkFlagRequired("requestmem")

	sizeCmd.Example = fmt.Sprintf("%s\n%s", "ktrouble add size --name small --limitscpu 250m --limitsmem 2Gi --requestcpu 100m --requestmem 512Mi", "ktrouble add size -n medium --limitscpu 500m --limitsmem 4Gi --requestcpu 200m --requestmem 1Gi")
}
