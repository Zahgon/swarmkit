package task

import (
	"errors"
	"fmt"
	"io"

	"github.com/moby/swarmkit/swarmd/cmd/swarmctl/common"
	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/cobra"
)

func printTaskStatus(w io.Writer, t *api.Task) { _ = "STUB: not implemented"; return }

func printTaskSummary(task *api.Task, res *common.Resolver) { _ = "STUB: not implemented"; return }

var (
	inspectCmd = &cobra.Command{
		Use:   "inspect <task ID>",
		Short: "Inspect a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("task ID missing")
			}

			if len(args) > 1 {
				return errors.New("inspect command takes exactly 1 argument")
			}

			c, err := common.Dial(cmd)
			if err != nil {
				return err
			}

			t, err := c.GetTask(common.Context(cmd), &api.GetTaskRequest{TaskID: args[0]})
			if err != nil {
				return err
			}
			task := t.Task

			r, err := c.ListTasks(common.Context(cmd),
				&api.ListTasksRequest{
					Filters: &api.ListTasksRequest_Filters{
						ServiceIDs: []string{task.ServiceID},
					},
				})
			if err != nil {
				return err
			}
			previous := []*api.Task{}
			for _, t := range r.Tasks {
				if t.Slot == task.Slot {
					previous = append(previous, t)
				}
			}

			res := common.NewResolver(cmd, c)

			printTaskSummary(task, res)
			if len(previous) > 0 {
				fmt.Println("\n===> Task Parents")
				Print(previous, true, res)
			}

			return nil
		},
	}
)
