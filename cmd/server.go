package cmd

import (
	"fmt"
	"os/exec"
	"taskx/tool"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the server and related services",
	Long:  ``,
	RunE:  startServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func startServer(*cobra.Command, []string) error {

	fmt.Println("Starting Server. Please wait...")

	shellCommand := exec.Command("docker-compose", "up", "--build", "-d")
	shellCommand.Env = tool.EnsureEnv()
	shellCommand.Dir = "server"
	err := shellCommand.Run()
	if err != nil {
		return err
	}

	fmt.Println("Started Server")

	return nil
}
