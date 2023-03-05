package get

import (
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

// serviceaccountCmd represents the serviceaccount command
var serviceaccountCmd = &cobra.Command{
	Use:     "serviceaccount",
	Aliases: []string{"serviceaccounts", "sa"},
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
		namespace := c.Client.DetermineNamespace(c.Namespace)

		sasList := c.Client.GetServiceAccounts(namespace)

		saData := objects.ServiceAccountList{}
		for _, v := range sasList.Items {
			saData.ServiceAccount = append(saData.ServiceAccount, v.Name)
		}

		c.OutputData(&saData, objects.TextOptions{NoHeaders: c.NoHeaders})

	},
}

func init() {
	getCmd.AddCommand(serviceaccountCmd)
}
