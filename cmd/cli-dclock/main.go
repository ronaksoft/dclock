package main

import (
	"fmt"
	"github.com/ronaksoft/rony"
	"github.com/ronaksoft/rony/config"
	"github.com/ronaksoft/rony/edge"
	"github.com/spf13/cobra"
	"time"
)

var (
	edgeServer *edge.Server
)

func main() {
	// Initialize the config package
	err := config.Init("dclock")
	if err != nil {
		fmt.Println("config initialization had error:", err)
	}
	rony.SetLogLevel(-1)

	// Set the flags as config parameters
	config.SetPersistentFlags(ServerCmd,
		config.StringFlag("serverID", "ServerID", ""),
		config.StringFlag("gatewayListen", "0.0.0.0:80", ""),
		config.StringSliceFlag("gatewayAdvertiseUrl", nil, ""),
		config.StringFlag("tunnelListen", "0.0.0.0:81", ""),
		config.StringSliceFlag("tunnelAdvertiseUrl", nil, ""),
		config.DurationFlag("idleTime", time.Minute, ""),
		config.IntFlag("raftPort", 7080, ""),
		config.UInt64Flag("replicaSet", 1, ""),
		config.IntFlag("gossipPort", 7081, ""),
		config.StringFlag("dataPath", "./_hdd", ""),
		config.BoolFlag("bootstrap", false, ""),
	)

	config.SetPersistentFlags(ClientCmd,
		config.StringFlag("clientID", "", ""),
		config.StringFlag("host", "127.0.0.1", ""),
		config.IntFlag("port", 80, ""),
		config.StringFlag("dataPath", "./_hdd", ""),
		config.StringFlag("hookID", "", ""),
		config.StringFlag("url", "", ""),
		config.Int64Flag("delay", 30, ""),
	)

	// Execute the cli command
	RootCmd.AddCommand(ServerCmd, ClientCmd)
	ClientCmd.AddCommand(HookSetCmd, HookDeleteCmd)
	err = RootCmd.Execute()
	if err != nil {
		fmt.Println("we got error:", err)
	}
}

var RootCmd = &cobra.Command{
	Use: "dclock",
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}
