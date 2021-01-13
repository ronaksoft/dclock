package main

import (
	"fmt"
	"github.com/ronaksoft/dclock/model"
	"github.com/ronaksoft/dclock/service"
	"github.com/ronaksoft/rony"
	"github.com/ronaksoft/rony/cluster"
	"github.com/ronaksoft/rony/config"
	"github.com/ronaksoft/rony/edge"
	"github.com/ronaksoft/rony/gateway"
	"github.com/ronaksoft/rony/repo/kv"
	"github.com/ronaksoft/rony/tools"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"runtime"
	"strings"
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

		// Initialize KV infrastructure
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

		// Set middlewares for logging, authorizing etc.
		edgeServer.SetPreHandlers(Authorize, Log)

		// Start the edge server components
		edgeServer.Start()

		// Initialize and Start Executor
		e := NewExecutor(runtime.NumCPU()*10, func(h *model.Hook) {
			err = model.ReadHook(h)
			if err != nil {
				fmt.Println(err)
			}
			_, err := http.DefaultClient.Post(h.GetCallbackUrl(), "application/json", strings.NewReader(h.GetJsonData()))
			if err != nil {
				fmt.Println("Error:", err)
			}
		})
		e.Start()

		// Wait until a shutdown signal received.
		edgeServer.ShutdownWithSignal(os.Kill, os.Interrupt)
		return nil
	},
}

func Authorize(ctx *edge.RequestCtx, in *rony.MessageEnvelope) {
	for _, x := range in.Header {
		switch x.Key {
		case "ClientID":
			ctx.Set("ClientID", x.Value)
		}
	}
}

func Log(ctx *edge.RequestCtx, in *rony.MessageEnvelope) {
	fmt.Println("Received Request", ctx.ReqID(), ctx.Conn().ClientIP(), ctx.Kind().String())
}
