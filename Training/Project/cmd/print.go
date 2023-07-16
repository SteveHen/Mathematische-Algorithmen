package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/falo2/ma/input"
	"github.com/falo2/ma/output"
)

var print = &cobra.Command{
	Use:   "print",
	Short: "Prints a textual representation of the given graph file",
	Long:  "",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return viper.BindPFlags(cmd.PersistentFlags())
	},
	Run: func(cmd *cobra.Command, args []string) {
		file := viper.GetString("file")
		isDirected := viper.GetBool("directed")
		hasBalance := viper.GetBool("balanced")

		graph, _, isWeighted := input.Read("graphs/" + file, isDirected, hasBalance)
		if graph == nil {
			fmt.Println("The file does not exist.")
			return
		}

		output.Print(graph, isWeighted, hasBalance)
	},
}
