package cmd

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"taskx-server/pb"
	"taskx/tool"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "List all of a user's tasks",
	Long:  ``,
	RunE:  getTasks,
}

func init() {
	clientCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getTasks(cmd *cobra.Command, args []string) error {
	conn, err := tool.GetClientConnection()
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pb.NewTaskXServiceClient(conn)
	resp, err := client.GetTasks(context.Background(), &emptypb.Empty{})
	if err != nil {
		return err
	}

	printTasks(resp)

	return nil
}

func printTasks(tasksResp *pb.GetTasksResp) {

	fmt.Println("Your Tasks:")
	fmt.Println("============")

	if len(tasksResp.Tasks) == 0 {
		fmt.Println("No tasks found")
		return
	} else {
		fmt.Println("Tasks:")
		for _, t := range tasksResp.Tasks {
			fmt.Printf("ID: %d, Name: %s, Description: %t\n", t.GetID(), t.GetDescription(), t.GetCompleted())
		}
	}

}
