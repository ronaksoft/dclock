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
	"github.com/valyala/fasthttp"
	"net/http"
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

		// Initialize KV infrastructure
		kv.MustInit(kv.Config{
			DirPath:             config.GetString("dataPath"),
			ConflictRetries:     100,
			ConflictMaxInterval: time.Millisecond,
		})

		rony.SetLogLevel(-1)

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
		service.RegisterClock(service.NewClock(edgeServer), edgeServer, edge.NewHandlerOptions().SetPreHandlers(Authorize, Log))
		if edgeServer.Cluster().ReplicaSet() == 1 {
			service.RegisterRoster(service.NewRoster(edgeServer), edgeServer, edge.NewHandlerOptions().SetPreHandlers(Log))
		}

		// Start the edge server components
		edgeServer.Start()

		// Join the cluster
		seedAddress := config.GetString("join")
		if len(seedAddress) > 0 {
			_, _ = edgeServer.JoinCluster(seedAddress)
		}

		// Initialize and Start Executor
		httpClient := &fasthttp.Client{
			Name: "dClock",
		}
		e := NewExecutor(runtime.NumCPU()*10, func(h *model.Hook) {
			h, err = model.ReadHook(h.GetClientID(), h.GetID(), h)
			if err != nil {
				fmt.Println(err)
				return
			}
			req := fasthttp.AcquireRequest()
			res := fasthttp.AcquireResponse()
			req.Header.SetMethod(http.MethodPost)
			req.Header.SetContentType("application/json")
			req.SetRequestURIBytes(h.GetCallbackUrl())
			req.SetBody(h.GetJsonData())
			err := httpClient.Do(req, res)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fasthttp.ReleaseRequest(req)
			fasthttp.ReleaseResponse(res)
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
	fmt.Println("Received Request", ctx.ReqID(), ctx.Kind().String())
}
