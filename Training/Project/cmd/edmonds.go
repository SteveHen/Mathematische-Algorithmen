package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/falo2/ma/input"
	"github.com/falo2/ma/output"

	"time"
)

var edmonds = &cobra.Command{
	Use:   "edmonds",
	Short: "Execute the edmonds karp algorithm on the given graph file",
	Long:  "",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return viper.BindPFlags(cmd.PersistentFlags())
	},
	Run: func(cmd *cobra.Command, args []string) {
		file := viper.GetString("file")
		fmt.Printf("=== %s ===\n", file)
		startTime := time.Now()

		fmt.Println("\n== Read ==")

		isDirected := viper.GetBool("directed")

		graph, _, _ := input.Read("graphs/"+file, isDirected, false)

		if graph == nil {
			fmt.Println("The file does not exist.")
			return
		}
		readTime := time.Now()

		fmt.Println("\n== Edmonds ==")

		var sourceNode int64
		var targetNode int64

		sourceNode = 0
		targetNode = 7

		if viper.GetString("source") != "" {
			sourceNode = viper.GetInt64("source")
		}

		if viper.GetString("target") != "" {
			targetNode = viper.GetInt64("target")
		}

		if viper.GetBool("verbose") {
			fmt.Println("= Original graph =")
			output.Print(graph, true, false)
			fmt.Println("")
		}

		flow, result, _ := graph.EdmondsKarpLegacy(sourceNode, targetNode)

		if viper.GetBool("verbose") {
			fmt.Println("= Resulting graph =")
			output.Print(result, true, false)
			fmt.Println("")
		}

		fmt.Println("Max flow:", flow)

		totalTime := time.Now()

		fmt.Println("\n== Time ==")
		fmt.Printf("Read    -> %vs (%s)\n", readTime.Sub(startTime).Seconds(), readTime.String())
		fmt.Printf("Edmonds -> %vs (%s)\n", totalTime.Sub(readTime).Seconds(), totalTime.String())
		fmt.Printf("Total   -> %vs (%s)\n", totalTime.Sub(startTime).Seconds(), totalTime.String())
	},
}
