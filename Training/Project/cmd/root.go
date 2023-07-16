package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "ma",
	Short: "Graph Library",
	Long:  "",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Init() {
	cobra.OnInitialize()

	dfs.PersistentFlags().String("file", "", "file containing the graph")
	if err := viper.BindPFlags(dfs.PersistentFlags()); err != nil {
		fmt.Printf("error binding flags: %s", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(dfs)

	print.PersistentFlags().String("file", "", "file containing the graph")
	if err := viper.BindPFlags(print.PersistentFlags()); err != nil {
		fmt.Printf("error binding flags: %s", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(print)

	prim.PersistentFlags().Bool("balanced", false, "balanced graph")
	prim.PersistentFlags().Bool("directed", false, "directed graph")
	prim.PersistentFlags().String("file", "", "file containing the graph")
	if err := viper.BindPFlags(prim.PersistentFlags()); err != nil {
		fmt.Printf("error binding flags: %s", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(prim)

	kruskal.PersistentFlags().String("file", "", "file containing the graph")
	if err := viper.BindPFlags(kruskal.PersistentFlags()); err != nil {
		fmt.Printf("error binding flags: %s", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(kruskal)

	nn.PersistentFlags().String("file", "", "file containing the graph")
	if err := viper.BindPFlags(nn.PersistentFlags()); err != nil {
		fmt.Printf("error binding flags: %s", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(nn)

	dt.PersistentFlags().String("file", "", "file containing the graph")
	if err := viper.BindPFlags(dt.PersistentFlags()); err != nil {
		fmt.Printf("error binding flags: %s", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(dt)

	bf.PersistentFlags().String("file", "", "file containing the graph")
	if err := viper.BindPFlags(bf.PersistentFlags()); err != nil {
		fmt.Printf("error binding flags: %s", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(bf)

	bb.PersistentFlags().String("file", "", "file containing the graph")
	if err := viper.BindPFlags(bb.PersistentFlags()); err != nil {
		fmt.Printf("error binding flags: %s", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(bb)

	dijkstra.PersistentFlags().Bool("directed", false, "directed graph")
	dijkstra.PersistentFlags().String("file", "", "file containing the graph")
	dijkstra.PersistentFlags().String("source", "", "start node")
	dijkstra.PersistentFlags().String("target", "", "target node")
	dijkstra.PersistentFlags().Bool("verbose", false, "verbose output")
	if err := viper.BindPFlags(dijkstra.PersistentFlags()); err != nil {
		fmt.Printf("error binding flags: %s", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(dijkstra)

	moore.PersistentFlags().Bool("directed", false, "directed graph")
	moore.PersistentFlags().String("file", "", "file containing the graph")
	moore.PersistentFlags().String("source", "", "start node")
	moore.PersistentFlags().String("target", "", "target node")
	moore.PersistentFlags().Bool("verbose", false, "verbose output")
	if err := viper.BindPFlags(moore.PersistentFlags()); err != nil {
		fmt.Printf("error binding flags: %s", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(moore)

	bfs.PersistentFlags().String("file", "", "file containing the graph")
	bfs.PersistentFlags().String("source", "", "start node")
	bfs.PersistentFlags().String("target", "", "target node")
	bfs.PersistentFlags().Bool("verbose", false, "verbose output")
	if err := viper.BindPFlags(bfs.PersistentFlags()); err != nil {
		fmt.Printf("error binding flags: %s", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(bfs)

	edmonds.PersistentFlags().Bool("directed", false, "directed graph")
	edmonds.PersistentFlags().String("file", "", "file containing the graph")
	edmonds.PersistentFlags().String("source", "", "start node")
	edmonds.PersistentFlags().String("target", "", "target node")
	edmonds.PersistentFlags().Bool("verbose", false, "verbose output")
	if err := viper.BindPFlags(edmonds.PersistentFlags()); err != nil {
		fmt.Printf("error binding flags: %s", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(edmonds)

	cc.PersistentFlags().Bool("directed", false, "directed graph")
	cc.PersistentFlags().String("file", "", "file containing the graph")
	cc.PersistentFlags().String("source", "", "start node")
	cc.PersistentFlags().String("target", "", "target node")
	cc.PersistentFlags().Bool("verbose", false, "verbose output")
	if err := viper.BindPFlags(cc.PersistentFlags()); err != nil {
		fmt.Printf("error binding flags: %s", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(cc)

	ssp.PersistentFlags().Bool("directed", false, "directed graph")
	ssp.PersistentFlags().String("file", "", "file containing the graph")
	ssp.PersistentFlags().String("source", "", "start node")
	ssp.PersistentFlags().String("target", "", "target node")
	ssp.PersistentFlags().Bool("verbose", false, "verbose output")
	if err := viper.BindPFlags(ssp.PersistentFlags()); err != nil {
		fmt.Printf("error binding flags: %s", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(ssp)
}
