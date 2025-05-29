package main

import (
	"fmt"
	tr "github.com/ocelot-cloud/task-runner"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var (
	projectDir = getParentDir()
)

func getParentDir() string {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		tr.CleanupAndExitWithError()
	}
	return filepath.Dir(currentDir)
}

func main() {
	go tr.HandleSignals()

	rootCmd := &cobra.Command{
		Use:   "ci-runner",
		Short: "ci-runner is a service that runs CI jobs",
	}

	rootCmd.AddCommand(deployCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	err := rootCmd.Execute()
	if err != nil {
		tr.CleanupAndExitWithError()
	}
}

const sshConfigHostName = "ocelot"

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy current version to server",
	Run: func(cmd *cobra.Command, args []string) {
		var remoteHomeDir = "/home/user"

		tr.Execute("hugo")
		executeOnServer("systemctl stop website")
		executeOnServer("mkdir -p %s/website", remoteHomeDir)
		rsyncCmd := fmt.Sprintf("rsync -avz --delete ./public/ %s:%s/website/", sshConfigHostName, remoteHomeDir)
		tr.ExecuteInDir(projectDir, rsyncCmd)
		executeOnServer("chown -R user:user %s/website", remoteHomeDir)
		executeOnServer("find /home/user/website -type d -exec chmod 755 {} +")
		executeOnServer("find /home/user/website -type f -exec chmod 644 {} +")
	},
}

func executeOnServer(command string, args ...string) {
	var cmd string
	if len(args) == 0 {
		cmd = command
	} else {
		cmd = fmt.Sprintf(command, args[0])
	}
	println("executing: " + cmd)
	tr.ExecuteInDir(".", "ssh ocelot \""+cmd+"\"")
}
