package cmd

import (
	"fmt"

	"github.com/imdario/mergo"
	"github.com/spf13/cobra"
)

func makeApplicationsCmd(dst cobra.Command) *cobra.Command {
	src := cobra.Command{
		Use:     "applications",
		Aliases: []string{"application", "apps", "app"},
	}

	if err := mergo.Merge(&dst, src); err != nil {
		panic(err)
	}

	return &dst
}

var getApplicationsCmd = makeApplicationsCmd(cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newAPIClient(cmd)
		if err != nil {
			return err
		}

		resources, err := client.ListApplications()
		if err != nil {
			return err
		}

		return outputList(cmd, resources)
	},
})

var deleteApplicationCmd = makeApplicationsCmd(cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newAPIClient(cmd)
		if err != nil {
			return err
		}

		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			return err
		}

		err = client.DeleteApplication(id)
		if err != nil {
			return err
		}

		fmt.Printf("Application '%v' deleted.\n", id)

		return nil
	},
})

func init() {
	getCmd.AddCommand(getApplicationsCmd)

	deleteCmd.AddCommand(deleteApplicationCmd)
	deleteApplicationCmd.Flags().Int("id", 0, "ID of the application to delete")
}
