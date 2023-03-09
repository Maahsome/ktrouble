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
	Short:   "Get a list of K8s ServiceAccount(s) in a Namespace",
	Long: `EXAMPLE:
  Return a list of kubernetes service accounts for a namespace

  > ktrouble get serviceaccount -n myspace

EXAMPLE:
  If you do not specify a namespace with '-n <namespace>', you will be prompted
  to select one

  > ktrouble get sa
`,
	Run: func(cmd *cobra.Command, args []string) {
		if c.Client != nil {
			namespace := c.Client.DetermineNamespace(c.Namespace)

			sasList := c.Client.GetServiceAccounts(namespace)

			saData := objects.ServiceAccountList{}
			for _, v := range sasList.Items {
				saData.ServiceAccount = append(saData.ServiceAccount, v.Name)
			}

			c.OutputData(&saData, objects.TextOptions{NoHeaders: c.NoHeaders})
		} else {
			common.Logger.Warn("Cannot fetch service accounts, no valid kubernetes context")
		}

	},
}

func init() {
	getCmd.AddCommand(serviceaccountCmd)
}
