package main

import (
	"github.com/ronaksoft/dclock/service"
	"github.com/ronaksoft/rony/cluster"
	"github.com/ronaksoft/rony/config"
	"github.com/ronaksoft/rony/edge"
	"github.com/ronaksoft/rony/gateway"
	"github.com/ronaksoft/rony/repo/kv"
	"github.com/ronaksoft/rony/tools"
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"time"
)

/*
   Creation Time: 2021 - Jan - 12
   Created by:  (ehsan)
   Maintainers:
      1.  Ehsan N. Moosa (E2)
   Auditor: Ehsan N. Moosa (E2)
   Copyright Ronak Software Group 2020
*/

var ServerCmd = &cobra.Command{
	Use: "server",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Bind the current flags to registered flags in config package
		err := config.BindCmdFlags(cmd)
		if err != nil {

			return err
		}

		kv.MustInit(kv.Config{
			DirPath:             config.GetString("dataPath"),
			ConflictRetries:     100,
			ConflictMaxInterval: time.Millisecond,
		})

		// Instantiate the edge server
		edgeServer = edge.NewServer(
			config.GetString("serverID"),
			edge.WithTcpGateway(edge.TcpGatewayConfig{
				Concurrency:   runtime.NumCPU() * 100,
				ListenAddress: config.GetString("gatewayListen"),
				MaxIdleTime:   config.GetDuration("idleTime"),
				Protocol:      gateway.Websocket,
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
