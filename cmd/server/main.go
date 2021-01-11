package main

import (
	"github.com/ronaksoft/rony/config"
	"github.com/ronaksoft/rony/edge"
	"github.com/spf13/cobra"
	"os"
)

var (
	edgeServer *edge.Server
)

func main() {
	config.MustInit("dclock")
	config.SetCmdPersistentFlags(RootCmd, config.StringFlag("serverID", "", ""))
	_ = RootCmd.Execute()
}

var RootCmd = &cobra.Command{
	Use: "dclock",
	Run: func(cmd *cobra.Command, args []string) {
		edgeServer = edge.NewServer(
			config.GetString("serverID"),
			edge.WithTcpGateway(edge.TcpGatewayConfig{}),
			edge.WithGossipCluster(edge.GossipClusterConfig{}),
			edge.WithUdpTunnel(edge.UdpTunnelConfig{}),
		)

		// Start the edge server components
		edgeServer.Start()

		// Wait until a shutdown signal received.
		edgeServer.ShutdownWithSignal(os.Kill, os.Interrupt)
	},
}
