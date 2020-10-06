package cmd

import (
	"fmt"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

var (
	data = [][]interface{}{
		{"gke-kapp-live-pool-1-3188", "38.5%", "28d"},
		{"gke-kapp-live-pool-1-5bfd", "0%", "5m"},
		{"gke-kapp-live-pool-1-es32", "85.9%", "1d 13h"},
	}
)

// nodesCmd represents the nodes command
var nodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "View cluster node attributes and status",
	Run: func(cmd *cobra.Command, args []string) {
		showNodesTable()
	},
}

func showNodesTable() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "CPU"},
			{Align: simpletable.AlignCenter, Text: "Age"},
		},
	}

	for _, row := range data {
		r := []*simpletable.Cell{
			{Text: row[0].(string)},
			{Text: row[1].(string)},
			{Text: row[2].(string)},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleDefault)
	fmt.Println(table.String())
}

func init() {
	rootCmd.AddCommand(nodesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
