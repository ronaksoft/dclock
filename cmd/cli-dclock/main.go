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
	rony.SetLogLevel(0)

	// Set the flags as config parameters
	config.SetPersistentFlags(ServerCmd,
		config.StringFlag("join", "", ""),
		config.StringFlag("serverID", "ServerID", ""),
		config.StringFlag("gatewayListen", "0.0.0.0:80", ""),
		config.StringSliceFlag("gatewayAdvertiseUrl", nil, ""),
		config.StringFlag("tunnelListen", "0.0.0.0:81", ""),
		config.StringSliceFlag("tunnelAdvertiseUrl", nil, ""),
		config.DurationFlag("idleTime", time.Minute, ""),
		config.IntFlag("raftPort", 7080, ""),
		config.Uint64Flag("replicaSet", 1, ""),
		config.IntFlag("gossipPort", 7081, ""),
		config.StringFlag("dataPath", "./_hdd", ""),
		config.BoolFlag("bootstrap", false, ""),
	)

	config.SetPersistentFlags(ClientCmd,
		config.StringFlag("hostPort", "127.0.0.1:81", ""),
		config.StringFlag("clientID", "989121228718", ""),
	)

	// Execute the cli command
	RootCmd.AddCommand(ServerCmd, ClientCmd)

	err = RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}

var RootCmd = &cobra.Command{
	Use: "dclock",
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}
