package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/falo2/ma/input"
	"github.com/falo2/ma/output"
)

var prim = &cobra.Command{
	Use:   "prim",
	Short: "Execute the prim algorithm on the given graph file",
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

		fmt.Println("\n== Prim ==")
		mst, totalWeight := graph.Prim()
		output.Print(mst, true, false)
		fmt.Println(totalWeight)
		totalTime := time.Now()

		fmt.Println("\n== Time ==")
		fmt.Printf("Read   -> %vs (%s)\n", readTime.Sub(startTime).Seconds(), readTime.String())
		fmt.Printf("Prim   -> %vs (%s)\n", totalTime.Sub(readTime).Seconds(), totalTime.String())
		fmt.Printf("Total  -> %vs (%s)\n", totalTime.Sub(startTime).Seconds(),totalTime.String())
	},
}
