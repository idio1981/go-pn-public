package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"go-pn-public/logger"

	"github.com/juju/fslock"
	"github.com/spf13/cobra"
)

func init() {
	var daemon bool
	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "Launch server.",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if daemon {
				file, _ := exec.LookPath(os.Args[0])
				logger.Info("Chdir : %s", filepath.Dir(file))

				logger.Info("Process Cmd: %s start -p %d >> %s", file, ARGV_LISTEN_PORT, file+"$(date +\"%Y.%m.%d\").log 2>&1")

				command := exec.Command("sh", "-c", fmt.Sprintf("%s start -p %d  >> %s", file, ARGV_LISTEN_PORT, file+"$(date +\"%Y.%m.%d\").log 2>&1"))
				command.Start()

				os.Exit(0)
			} else {
				file, _ := exec.LookPath(os.Args[0])
				logger.Info("Chdir : %s", filepath.Dir(file))
				filename := filepath.Base(file)

				lock := fslock.New(filename + ".lock")
				lockErr := lock.TryLock()
				if lockErr != nil {
					fmt.Println("falied to acquire lock > " + lockErr.Error())
					return
				}

				os.WriteFile(filename+".lock", []byte(fmt.Sprintf("%d", os.Getpid())), 0666)
				rootCmd.Run(cmd, args)
				lock.Unlock()

			}
		},
	}

	startCmd.Flags().BoolVarP(&daemon, "deamon", "d", false, "daemon mode")
	startCmd.Flags().IntVarP(&ARGV_LISTEN_PORT, "port", "p", 0, "listen port")
	startCmd.MarkFlagRequired("port")

	rootCmd.AddCommand(startCmd)
}
