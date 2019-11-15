package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/paultyng/go-newrelic/v4/api"
)

var cfgFile string
var debug bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "newrelic",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	var err error

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.newrelic.yaml)")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Run in debug mode")
	RootCmd.PersistentFlags().String("api-key", "", "New Relic API key")

	if err = viper.BindPFlag("api-key", RootCmd.PersistentFlags().Lookup("api-key")); err != nil {
		fmt.Println(err)
	}

	RootCmd.PersistentFlags().String("api-url", "", "Base URL for the New Relic API")

	if err = viper.BindPFlag("api-url", RootCmd.PersistentFlags().Lookup("api-url")); err != nil {
		fmt.Println(err)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".newrelic") // name of config file (without extension)
	viper.AddConfigPath("$HOME")     // adding home directory as first search path

	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("newrelic")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}

func newAPIClient(cmd *cobra.Command) (api.Client, error) {
	apiKey := viper.GetString("api-key")
	baseURL := viper.GetString("api-url")

	config := api.Config{
		APIKey:  apiKey,
		BaseURL: baseURL,
		Debug:   debug,
	}

	client := api.New(config)

	return client, nil
}

func newInfraClient(cmd *cobra.Command) (api.InfraClient, error) {
	apiKey := viper.GetString("api-key")
	baseURL := viper.GetString("api-url")

	config := api.Config{
		APIKey:  apiKey,
		BaseURL: baseURL,
		Debug:   debug,
	}

	client := api.NewInfraClient(config)

	return client, nil
}
