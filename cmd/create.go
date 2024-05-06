package cmd

import (
	"context"
	"fmt"
	"taskx-server/pb"
	"taskx/tool"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new task",
	Long:  ``,
	RunE:  createTask,
}

func init() {
	clientCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("description", "d", "", "Description of the task")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createTask(cmd *cobra.Command, args []string) error {

	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return err
	}

	conn, err := tool.GetClientConnection()
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pb.NewTaskXServiceClient(conn)
	_, err = client.AddTask(context.Background(), &pb.AddTaskReq{
		Description: description,
	})
	if err != nil {
		return err
	}

	fmt.Println("Task created!")

	return nil
}
