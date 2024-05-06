package cmd

import (
	"context"
	"fmt"
	"strconv"
	"taskx-server/pb"
	"taskx/tool"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task as completed",
	Long:  ``,
	RunE:  completeTask,
}

func init() {
	clientCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func completeTask(cmd *cobra.Command, args []string) error {

	taskID, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return err
	}

	conn, err := tool.GetClientConnection()
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pb.NewTaskXServiceClient(conn)
	_, err = client.CompleteTask(context.Background(), &pb.CompleteTaskReq{
		ID: taskID,
	})
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Task with ID:%s marked as complete", args[0]))

	return nil
}
