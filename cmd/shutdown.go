/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"taskx/tool"

	"github.com/spf13/cobra"
)

// shutdownCmd represents the shutdown command
var shutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "A brief description of your command",
	Long:  ``,
	RunE:  shutdownServer,
}

func init() {
	serverCmd.AddCommand(shutdownCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shutdownCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shutdownCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func shutdownServer(*cobra.Command, []string) error {

	fmt.Println("Stopping Server...")

	shellCommand := exec.Command("docker-compose", "down")
	shellCommand.Env = tool.EnsureEnv()
	shellCommand.Dir = "server"
	err := shellCommand.Run()
	if err != nil {
		return err
	}

	fmt.Println("Server Stopped")

	return nil
}
