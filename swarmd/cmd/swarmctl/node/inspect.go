package node

import (
	"errors"
	"fmt"

	"github.com/moby/swarmkit/swarmd/cmd/swarmctl/common"
	"github.com/moby/swarmkit/swarmd/cmd/swarmctl/task"
	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/cobra"
)

func printNodeSummary(node *api.Node) { _ = "STUB: not implemented"; return }

// Ignore flushing errors - there's nothing we can do.

// sort label output for readability

// append to pluginTypes only if not done previously

// ensure stable output

var (
	inspectCmd = &cobra.Command{
		Use:   "inspect <node ID>",
		Short: "Inspect a node",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("node ID missing")
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

			node, err := getNode(common.Context(cmd), c, args[0])
			if err != nil {
				return err
			}

			r, err := c.ListTasks(common.Context(cmd),
				&api.ListTasksRequest{
					Filters: &api.ListTasksRequest_Filters{
						NodeIDs: []string{node.ID},
					},
				})
			if err != nil {
				return err
			}

			printNodeSummary(node)
			if len(r.Tasks) > 0 {
				fmt.Println()
				task.Print(r.Tasks, all, common.NewResolver(cmd, c))
			}

			return nil
		},
	}
)

func init() {
	inspectCmd.Flags().BoolP("all", "a", false, "Show all tasks (default shows just running)")
}
