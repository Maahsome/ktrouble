package get

import (
	"ktrouble/common"
	"ktrouble/defaults"
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

// serviceaccountCmd represents the serviceaccount command
var serviceaccountCmd = &cobra.Command{
	Use:     "serviceaccount",
	Aliases: defaults.GetServiceAccountsAliases,
	Short:   getServiceAccountHelp.Short(),
	Long:    getServiceAccountHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		if c.Client != nil {
			namespace := c.Client.DetermineNamespace(c.Namespace)

			sasList := c.Client.GetServiceAccounts(namespace)

			saData := objects.ServiceAccountList{}
			for _, v := range sasList.Items {
				saData.ServiceAccount = append(saData.ServiceAccount, v.Name)
			}

			c.OutputData(&saData, objects.TextOptions{
				NoHeaders:     c.NoHeaders,
				Fields:        c.Fields,
				DefaultFields: c.OutputFieldsMap["service_account"],
			})
		} else {
			common.Logger.Warn("Cannot fetch service accounts, no valid kubernetes context")
		}

	},
}

func init() {
	getCmd.AddCommand(serviceaccountCmd)
}
