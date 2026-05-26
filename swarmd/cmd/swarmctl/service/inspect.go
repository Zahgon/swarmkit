package service

import (
	"errors"
	"fmt"

	"github.com/moby/swarmkit/swarmd/cmd/swarmctl/common"
	"github.com/moby/swarmkit/swarmd/cmd/swarmctl/task"
	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/cobra"
)

func printServiceSummary(service *api.Service, running int) { _ = "STUB: not implemented"; return }

var (
	inspectCmd = &cobra.Command{
		Use:   "inspect <service ID>",
		Short: "Inspect a service",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("service ID missing")
			}

			if len(args) > 1 {
				return errors.New("inspect command takes exactly 1 argument")
			}

			flags := cmd.Flags()

			all, err := flags.GetBool("all")
			if err != nil {
				return err
			}

			c, err := common.Dial(cmd)
			if err != nil {
				return err
			}

			res := common.NewResolver(cmd, c)

			service, err := getService(common.Context(cmd), c, args[0])
			if err != nil {
				return err
			}

			r, err := c.ListTasks(common.Context(cmd),
				&api.ListTasksRequest{
					Filters: &api.ListTasksRequest_Filters{
						ServiceIDs: []string{service.ID},
					},
				})
			if err != nil {
				return err
			}
			var running int
			for _, t := range r.Tasks {
				if t.Status.State == api.TaskStateRunning {
					running++
				}
			}

			printServiceSummary(service, running)
			if len(r.Tasks) > 0 {
				fmt.Println()
				task.Print(r.Tasks, all, res)
			}

			return nil
		},
	}
)

func init() {
	inspectCmd.Flags().BoolP("all", "a", false, "Show all tasks (default shows just running)")
}
