package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shorten",
	Short: "Shorten is a simple App to provide ShortURLs",
	Long: `Shorten is a simple App to provide ShortURLs, simply use generate
to return a short URL version of the long URL. For example: 
$ shorten generate https://www.amazon.com`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {

}
