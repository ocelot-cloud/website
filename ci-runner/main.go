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
	hugoDir    = projectDir + "/hugo"
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
		executeOnServer("systemctl stop website")
		executeOnServer("mkdir -p %s/website", remoteHomeDir)
		rsyncCmd := fmt.Sprintf("rsync -avz ./ %s:%s/website/", sshConfigHostName, remoteHomeDir)
		tr.ExecuteInDir(hugoDir, rsyncCmd)
		executeOnServer("chown -R user:user %s/website", remoteHomeDir)
		executeOnServer("chmod -R 700 %s/website", remoteHomeDir)
		executeOnServer("systemctl start website")
		executeOnServer("nmap -p 1313 localhost")
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
