package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/falo2/ma/input"
)

var bfs = &cobra.Command{
	Use:   "bfs",
	Short: "Execute breadth search on the given graph file",
	Long:  "",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return viper.BindPFlags(cmd.PersistentFlags())
	},
	Run: func(cmd *cobra.Command, args []string) {
		file := viper.GetString("file")
		fmt.Printf("=== %s ===\n", file)
		startTime := time.Now()

		fmt.Println("\n== Read ==")
		graph, _, _ := input.Read("graphs/" + file, false, false)
		if graph == nil {
			fmt.Println("The file does not exist.")
			return
		}

		readTime := time.Now()

		var sourceNode int64
		var targetNode int64

		if viper.GetString("source") != "" {
			sourceNode = viper.GetInt64("source")
		}

		if viper.GetString("target") != "" {
			targetNode = viper.GetInt64("target")
		}

		fmt.Println("\n== Search ==")
		graph.BFSLegacy(sourceNode, targetNode)
		searchTime := time.Now()

		fmt.Println("\n== Time ==")
		fmt.Println("Start  -> " + startTime.String())
		fmt.Println("Read   -> " + readTime.String())
		fmt.Println("Search -> " + searchTime.String())
	},
}
