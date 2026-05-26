package network

import (
	"errors"
	"fmt"
	"strings"

	"github.com/moby/swarmkit/swarmd/cmd/swarmctl/common"
	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/cobra"
)

var (
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a network",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				return errors.New("create command takes no arguments")
			}

			flags := cmd.Flags()
			if !flags.Changed("name") {
				return errors.New("--name is required")
			}

			name, err := flags.GetString("name")
			if err != nil {
				return err
			}

			// Process driver configurations
			var driver *api.Driver
			if flags.Changed("driver") {
				driver = new(api.Driver)

				driverName, err := flags.GetString("driver")
				if err != nil {
					return err
				}

				driver.Name = driverName

				opts, err := cmd.Flags().GetStringSlice("opts")
				if err != nil {
					return err
				}

				driver.Options = map[string]string{}
				for _, opt := range opts {
					optPair := strings.Split(opt, "=")
					if len(optPair) != 2 {
						return fmt.Errorf("Malformed opts: %s", opt)
					}
					driver.Options[optPair[0]] = optPair[1]
				}
			}

			ipamOpts, err := processIPAMOptions(cmd)
			if err != nil {
				return err
			}

			spec := &api.NetworkSpec{
				Annotations: api.Annotations{
					Name: name,
				},
				DriverConfig: driver,
				IPAM:         ipamOpts,
			}

			c, err := common.Dial(cmd)
			if err != nil {
				return err
			}
			r, err := c.CreateNetwork(common.Context(cmd), &api.CreateNetworkRequest{Spec: spec})
			if err != nil {
				return err
			}
			fmt.Println(r.Network.ID)
			return nil
		},
	}
)

func processIPAMOptions(cmd *cobra.Command) (*api.IPAMOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func init() {
	createCmd.Flags().String("name", "", "Network name")
	createCmd.Flags().String("driver", "", "Network driver")
	createCmd.Flags().String("ipam-driver", "", "IPAM driver")
	createCmd.Flags().StringSlice("subnet", []string{}, "Subnets in CIDR format that represents a network segments")
	createCmd.Flags().StringSlice("gateway", []string{}, "Gateway IP addresses for network segments")
	createCmd.Flags().StringSlice("ip-range", []string{}, "IP ranges to allocate from within the subnets")
	createCmd.Flags().StringSlice("opts", []string{}, "Network driver options")
}
