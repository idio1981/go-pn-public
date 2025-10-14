package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	ARGV_LISTEN_PORT int
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Start server",
	Long:  `Start server`,
}

func Execute(f func(cmd *cobra.Command, args []string)) {
	rootCmd.Run = f

	rootCmd.Flags().IntVarP(&ARGV_LISTEN_PORT, "port", "p", 0, "listen port")
	rootCmd.MarkFlagRequired("port")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
