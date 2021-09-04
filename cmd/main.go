package cmd

import (
	"github.com/clh021/crud-api/api"
	"github.com/spf13/cobra"
)

// mainCmd represents the main command
var mainCmd = &cobra.Command{
	Use:   "main",
	Short: "server for web api",
	Long:  `crud-api is a tool for translate database opera as some web api. Just conf for all.`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiServer()
	},
}

func init() {
	rootCmd.AddCommand(mainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ApiServer() {
	api.Main()
}
