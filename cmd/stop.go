package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/idio1981/go-pn-public/v1/logger"

	"github.com/spf13/cobra"
)

func init() {
	var stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Shutdown server.",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			file, _ := exec.LookPath(os.Args[0])
			logger.Info("Process File: %s", file)

			err := os.Chdir(filepath.Dir(file))
			if err != nil {
				panic(err)
			}
			logger.Info("Chdir : %s", filepath.Dir(file))

			filename := filepath.Base(file)
			strb, _ := os.ReadFile(filename + ".lock")

			command := exec.Command("kill", string(strb))
			command.Start()

			fmt.Printf("Server shutdown, [PID] %s\n", string(strb))
			os.Exit(0)
		},
	}

	rootCmd.AddCommand(stopCmd)
}
