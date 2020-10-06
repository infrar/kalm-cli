package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// routesListCmd represents the routesList command
var routesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List certificates",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("routesList called")
	},
}

func init() {
	routesCmd.AddCommand(routesListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// routesListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// routesListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
