package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// disksCmd represents the disks command
var disksCmd = &cobra.Command{
	Use:   "disks",
	Short: "View information related to disks and volumes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("disks called")
	},
}

func init() {
	rootCmd.AddCommand(disksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// disksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// disksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
