package main

import (
	"fmt"
	"github.com/ronaksoft/dclock/service"
	"github.com/ronaksoft/rony/cluster"
	"github.com/ronaksoft/rony/config"
	"github.com/ronaksoft/rony/edge"
	"github.com/ronaksoft/rony/gateway"
	"github.com/ronaksoft/rony/tools"
	"github.com/spf13/cobra"
	"os"
	"runtime"
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

	// Set the flags as config parameters
	config.SetCmdPersistentFlags(RootCmd,
		config.StringFlag("serverID", tools.RandomID(12), ""),
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

	// Execute the cli command
	err = RootCmd.Execute()
	if err != nil {
		fmt.Println("we got error:", err)
	}
}

var RootCmd = &cobra.Command{
	Use: "dclock",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Bind the current flags to registered flags in config package
		err := config.BindCmdFlags(cmd)
		if err != nil {

			return err
		}

		// Instantiate the edge server
		edgeServer = edge.NewServer(
			config.GetString("serverID"),
			edge.WithTcpGateway(edge.TcpGatewayConfig{
				Concurrency:   runtime.NumCPU() * 100,
				ListenAddress: config.GetString("gatewayListen"),
				MaxIdleTime:   config.GetDuration("idleTime"),
				Protocol:      gateway.Http,
				ExternalAddrs: config.GetStringSlice("gatewayAdvertiseUrl"),
			}),
			edge.WithGossipCluster(edge.GossipClusterConfig{
				ServerID:   tools.StrToByte(config.GetString("serverID")),
				Bootstrap:  config.GetBool("bootstrap"),
				RaftPort:   config.GetInt("raftPort"),
				ReplicaSet: config.GetUint64("replicaSet"),
				Mode:       cluster.MultiReplica,
				GossipPort: config.GetInt("gossipPort"),
				DataPath:   config.GetString("dataPath"),
			}),
			edge.WithUdpTunnel(edge.UdpTunnelConfig{
				ServerID:      config.GetString("serverID"),
				Concurrency:   runtime.NumCPU() * 100,
				ListenAddress: config.GetString("tunnelListen"),
				ExternalAddrs: config.GetStringSlice("tunnelAdvertiseUrl"),
			}),
		)

		// Register the service into the edge server
		service.RegisterClock(&service.Clock{}, edgeServer)

		// Start the edge server components
		edgeServer.Start()

		// Wait until a shutdown signal received.
		edgeServer.ShutdownWithSignal(os.Kill, os.Interrupt)
		return nil
	},
}
