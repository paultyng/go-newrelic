package cmd

import (
	"fmt"

	"github.com/imdario/mergo"
	"github.com/spf13/cobra"

	"github.com/paultyng/go-newrelic/v4/api"
)

func makeDashboardsCmd(dst cobra.Command) *cobra.Command {
	src := cobra.Command{
		Use:     "dashboards",
		Aliases: []string{"dashboards", "db"},
	}

	if err := mergo.Merge(&dst, src); err != nil {
		panic(err)
	}

	return &dst
}

var getDashboardsCmd = makeDashboardsCmd(cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newAPIClient(cmd)
		if err != nil {
			return err
		}

		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			return err
		}

		var resources []api.Dashboard

		if id != 0 {
			var resource *api.Dashboard

			resource, err = client.GetDashboard(id)
			if err != nil {
				return err
			}

			resources = []api.Dashboard{*resource}
		} else {
			resources, err = client.ListDashboards()
			if err != nil {
				return err
			}
		}

		return outputList(cmd, resources)
	},
})

var deleteDashboardCmd = makeDashboardsCmd(cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newAPIClient(cmd)
		if err != nil {
			return err
		}

		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			return err
		}

		err = client.DeleteDashboard(id)
		if err != nil {
			return err
		}

		fmt.Printf("Dashboard '%v' deleted.\n", id)

		return nil
	},
})

func init() {
	getCmd.AddCommand(getDashboardsCmd)
	getDashboardsCmd.Flags().Int("id", 0, "ID of the dashboard to get")

	deleteCmd.AddCommand(deleteDashboardCmd)
	deleteCmd.Flags().Int("id", 0, "ID of the dashboard to delete")
}
