package cmd

import (
	"github.com/imdario/mergo"
	"github.com/spf13/cobra"
)

func makeChannelsCmd(dst cobra.Command) *cobra.Command {
	src := cobra.Command{
		Use:     "channels",
		Aliases: []string{"channel", "ch"},
	}

	if err := mergo.Merge(&dst, src); err != nil {
		panic(err)
	}

	return &dst
}

var getAlertChannelsCmd = makeChannelsCmd(cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newApiClient(cmd)
		if err != nil {
			return err
		}
		resources, err := client.ListAlertChannels()
		if err != nil {
			return err
		}

		return outputTable(cmd, resources)
	},
})

func init() {
	GetCmd.AddCommand(getAlertChannelsCmd)
}
