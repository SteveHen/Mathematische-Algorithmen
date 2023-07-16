package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/falo2/ma/input"
	"github.com/falo2/ma/output"

	"time"
)

var cc = &cobra.Command{
	Use:   "cc",
	Short: "Execute the cycle canceling algorithm on the given graph file",
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

		graph, balance, _ := input.Read("graphs/"+file, isDirected, true)

		if graph == nil {
			fmt.Println("The file does not exist.")
			return
		}

		if balance == nil {
			fmt.Println("The file does not contain a balance.")
			return
		}

		readTime := time.Now()

		fmt.Println("\n== Cycle Canceling ==")

		if viper.GetBool("verbose") {
			fmt.Println("= Balance =")
			balanceList := *balance
			for i := 0; i < len(balanceList); i++ {
				fmt.Printf("%d: %f\n", i, balanceList[int64(i)])
			}

			fmt.Println("")
			fmt.Println("= Original graph =")
			output.Print(graph, true, true)
			fmt.Println("")
		}

		result, flow := graph.CycleCanceling(*balance, viper.GetBool("verbose"))

		if result == nil {
			fmt.Println("The graph is not balanced.")
		} else {
			if viper.GetBool("verbose") {
				fmt.Println("")
				fmt.Println("= Resulting graph =")
				output.Print(result, true, true)
				fmt.Println("")
			}

			fmt.Println("Cost min flow:", flow)
		}

		totalTime := time.Now()

		fmt.Println("\n== Time ==")
		fmt.Printf("Read  -> %vs (%s)\n", readTime.Sub(startTime).Seconds(), readTime.String())
		fmt.Printf("CC    -> %vs (%s)\n", totalTime.Sub(readTime).Seconds(), totalTime.String())
		fmt.Printf("Total -> %vs (%s)\n", totalTime.Sub(startTime).Seconds(), totalTime.String())
	},
}
