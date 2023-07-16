package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/falo2/ma/input"
	"github.com/falo2/ma/output"

	"time"
)

var moore = &cobra.Command{
	Use:   "moore",
	Short: "Execute the moore algorithm on the given graph file",
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

		fmt.Println("\n== Moore ==")

		var sourceNode int64
		var targetNode int64

		switch file {
		case "Wege1.txt", "Wege2.txt":
			sourceNode = 2
			targetNode = 0
		default:
			sourceNode = 0
			targetNode = 1
		}

		if viper.GetString("source") != "" {
			sourceNode = viper.GetInt64("source")
		}

		if viper.GetString("target") != "" {
			targetNode = viper.GetInt64("target")
		}

		path, pathWeight, negativeCycle := graph.MooreBellmanFordLegacy(sourceNode, targetNode, viper.GetBool("verbose"))

		if len(negativeCycle) > 0 {
 			fmt.Println("! Negative cycle detected !")
			fmt.Println(negativeCycle)
		} else {
			fmt.Println("Node", sourceNode, "-> Node", targetNode, "(", pathWeight, ")")

			if viper.GetBool("verbose") {
				fmt.Println("\n= Tree =")
				output.Print(path, true, false)
			}
		}

		totalTime := time.Now()

		fmt.Println("\n== Time ==")
		fmt.Printf("Read  -> %vs (%s)\n", readTime.Sub(startTime).Seconds(), readTime.String())
		fmt.Printf("Moore -> %vs (%s)\n", totalTime.Sub(readTime).Seconds(), totalTime.String())
		fmt.Printf("Total -> %vs (%s)\n", totalTime.Sub(startTime).Seconds(), totalTime.String())
	},
}
