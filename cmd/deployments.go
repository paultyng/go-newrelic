package cmd

import (
	"fmt"

	"github.com/imdario/mergo"
	"github.com/spf13/cobra"

	"github.com/paultyng/go-newrelic/v4/api"
)

func makeDeploymentsCmd(dst cobra.Command) *cobra.Command {
	src := cobra.Command{
		Use:     "deployments",
		Aliases: []string{"deployment", "deploy"},
	}

	if err := mergo.Merge(&dst, src); err != nil {
		panic(err)
	}

	return &dst
}

var getDeploymentsCmd = makeDeploymentsCmd(cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newAPIClient(cmd)
		if err != nil {
			return err
		}
		applicationID, err := cmd.Flags().GetInt("application-id")
		if err != nil {
			return err
		}
		resources, err := client.ListDeployments(applicationID)
		if err != nil {
			return err
		}

		return outputList(cmd, resources)
	},
})

var createDeploymentCmd = makeDeploymentsCmd(cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newAPIClient(cmd)
		if err != nil {
			return err
		}

		applicationID, err := cmd.Flags().GetInt("application-id")
		if err != nil {
			return err
		}

		revision, err := cmd.Flags().GetString("revision")
		if err != nil {
			return err
		}

		user, err := cmd.Flags().GetString("user")
		if err != nil {
			return err
		}

		changelog, err := cmd.Flags().GetString("changelog")
		if err != nil {
			return err
		}

		description, err := cmd.Flags().GetString("description")
		if err != nil {
			return err
		}

		_, err = client.CreateDeployment(applicationID, api.Deployment{
			Revision:    revision,
			User:        user,
			Changelog:   changelog,
			Description: description,
		})
		if err != nil {
			return err
		}

		fmt.Printf("Deployment created for application ID '%v'.\n", applicationID)

		return nil
	},
})

func init() {
	getCmd.AddCommand(getDeploymentsCmd)
	getDeploymentsCmd.Flags().IntP("application-id", "a", 0, "application ID of the deployments to get")
	getDeploymentsCmd.MarkFlagRequired("application-id")

	createCmd.AddCommand(createDeploymentCmd)
	createDeploymentCmd.Flags().IntP("application-id", "a", 0, "application ID of the deployment to create")
	createDeploymentCmd.Flags().StringP("revision", "r", "", "Revision of the deployment")
	createDeploymentCmd.Flags().StringP("user", "u", "", "User posting the deployment")
	createDeploymentCmd.Flags().StringP("changelog", "c", "", "Changelog of the deployment")
	createDeploymentCmd.Flags().StringP("description", "d", "", "Description of the deployment")
	createDeploymentCmd.MarkFlagRequired("application-id")
	createDeploymentCmd.MarkFlagRequired("revision")
}
