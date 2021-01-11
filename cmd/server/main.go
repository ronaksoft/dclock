package main

import (
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
	config.MustInit("dclock")
	config.SetCmdPersistentFlags(RootCmd, config.StringFlag("serverID", "", ""))
	config.SetCmdPersistentFlags(RootCmd, config.StringFlag("gatewayListen", "0.0.0.0:80", ""))
	config.SetCmdPersistentFlags(RootCmd, config.StringSliceFlag("gatewayAdvertiseUrl", nil, ""))
	config.SetCmdPersistentFlags(RootCmd, config.StringFlag("tunnelListen", "0.0.0.0:81", ""))
	config.SetCmdPersistentFlags(RootCmd, config.StringSliceFlag("tunnelAdvertiseUrl", nil, ""))
	config.SetCmdPersistentFlags(RootCmd, config.DurationFlag("idleTime", time.Minute, ""))
	config.SetCmdPersistentFlags(RootCmd, config.IntFlag("raftPort", 7080, ""))
	config.SetCmdPersistentFlags(RootCmd, config.UInt64Flag("replicaSet", 1, ""))
	config.SetCmdPersistentFlags(RootCmd, config.IntFlag("gossipPort", 7081, ""))
	config.SetCmdPersistentFlags(RootCmd, config.StringFlag("dataPath", "./_hdd", ""))
	config.SetCmdPersistentFlags(RootCmd, config.BoolFlag("bootstrap", false, ""))
	_ = RootCmd.Execute()
}

var RootCmd = &cobra.Command{
	Use: "dclock",
	Run: func(cmd *cobra.Command, args []string) {
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

		// Start the edge server components
		edgeServer.Start()

		// Wait until a shutdown signal received.
		edgeServer.ShutdownWithSignal(os.Kill, os.Interrupt)
	},
}
