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

const sshConfigHostName = "website"

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy current version to server",
	Run: func(cmd *cobra.Command, args []string) {
		executeOnServer("docker rm -f website || true")
		tr.ExecuteInDir(projectDir, "hugo")
		rsyncCmd := fmt.Sprintf("rsync -avz --delete docker-compose.yml ./public %s:", sshConfigHostName)
		tr.ExecuteInDir(projectDir, rsyncCmd)
		executeOnServer("docker compose up -d")
		executeOnServer("curl https://ocelot-cloud.org >> /dev/null")
	},
}

func executeOnServer(command string) {
	sshCommand := fmt.Sprintf("ssh %s %s", sshConfigHostName, command)
	tr.Execute(sshCommand)
}
