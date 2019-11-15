package cmd

import (
	"github.com/imdario/mergo"
	"github.com/spf13/cobra"

	"github.com/paultyng/go-newrelic/v4/api"
)

func makeKeyTransactionsCmd(dst cobra.Command) *cobra.Command {
	src := cobra.Command{
		Use:     "key-transactions",
		Aliases: []string{"transactions", "kt"},
	}

	if err := mergo.Merge(&dst, src); err != nil {
		panic(err)
	}

	return &dst
}

var getKeyTransactionsCmd = makeKeyTransactionsCmd(cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newAPIClient(cmd)
		if err != nil {
			return err
		}

		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			return err
		}

		var resources []api.KeyTransaction

		if id != 0 {
			var resource *api.KeyTransaction

			resource, err = client.GetKeyTransaction(id)
			if err != nil {
				return err
			}

			resources = []api.KeyTransaction{*resource}
		} else {
			resources, err = client.ListKeyTransactions()
			if err != nil {
				return err
			}
		}

		return outputList(cmd, resources)
	},
})

func init() {
	getCmd.AddCommand(getKeyTransactionsCmd)
	getKeyTransactionsCmd.Flags().Int("id", 0, "ID of the key transaction to get")
}
