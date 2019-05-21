package cmd

import (
	"github.com/imdario/mergo"
	"github.com/spf13/cobra"
)

func makeAlertInfraConditionsCmd(dst cobra.Command) *cobra.Command {
	src := cobra.Command{
		Use:     "infra-conditions",
		Aliases: []string{"infra-condition", "infra-cond"},
	}

	if err := mergo.Merge(&dst, src); err != nil {
		panic(err)
	}

	return &dst
}

var getAlertInfraConditionsCmd = makeAlertInfraConditionsCmd(cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newInfraClient(cmd)
		if err != nil {
			return err
		}
		policyID, err := cmd.Flags().GetInt("policy-id")
		if err != nil {
			return err
		}
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			return err
		}
		condition, err := client.GetAlertInfraCondition(policyID, id)

		if err != nil {
			return err
		}

		return outputJSON(cmd.OutOrStdout(), condition)
	},
})

func init() {
	getCmd.AddCommand(getAlertInfraConditionsCmd)
	getAlertInfraConditionsCmd.Flags().IntP("policy-id", "", 0, "ID of policy for which to get conditions")
	getAlertInfraConditionsCmd.Flags().IntP("id", "", 0, "ID of the condition")
}
