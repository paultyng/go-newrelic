# cmd
--
    import "github.com/paultyng/go-newrelic/internal/cmd"


## Usage

```go
var RootCmd = &cobra.Command{
	Use:   "newrelic",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}
```
RootCmd represents the base command when called without any subcommands

#### func  Execute

```go
func Execute()
```
Execute adds all child commands to the root command sets flags appropriately.
This is called by main.main(). It only needs to happen once to the rootCmd.
