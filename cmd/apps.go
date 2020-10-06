package cmd

import (
	"github.com/spf13/cobra"
)

var fakeApps = []string{"hello-kalm", "myblog"}

// appsCmd represents the apps command
var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "List your apps",
}

func init() {
	rootCmd.AddCommand(appsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// appsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// appsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
