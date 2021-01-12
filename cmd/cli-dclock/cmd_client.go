package main

import (
	"fmt"
	"github.com/ronaksoft/dclock/service"
	"github.com/ronaksoft/rony/config"
	"github.com/ronaksoft/rony/edgec"
	"github.com/ronaksoft/rony/repo/kv"
	"github.com/ronaksoft/rony/tools"
	"github.com/spf13/cobra"
)

/*
   Creation Time: 2021 - Jan - 12
   Created by:  (ehsan)
   Maintainers:
      1.  Ehsan N. Moosa (E2)
   Auditor: Ehsan N. Moosa (E2)
   Copyright Ronak Software Group 2020
*/

var ClientCmd = &cobra.Command{
	Use: "client",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func prepareCmd(cmd *cobra.Command) (*service.ClockClient, error) {
	// Bind the current flags to registered flags in config package
	err := config.BindCmdFlags(cmd)
	if err != nil {
		return nil, err
	}

	cmd.Println(config.GetString("dataPath"), config.GetString("host"), config.GetInt("port"))
	kv.MustInit(kv.Config{
		DirPath:             config.GetString("dataPath"),
		ConflictRetries:     0,
		ConflictMaxInterval: 0,
	})

	wsC := edgec.NewWebsocket(edgec.WebsocketConfig{
		SeedHostPort: fmt.Sprintf("%s:%d", config.GetString("host"), config.GetInt("port")),
		Header: map[string]string{
			"APIKEY": "",
		},
		Router:  nil,
		Secure:  false,
	})

	err = wsC.Start()
	if err != nil {
		return nil, err
	}

	cli := service.NewClockClient(wsC)
	return cli, nil
}

var HookSetCmd = &cobra.Command{
	Use: "HookSet",
	RunE: func(cmd *cobra.Command, args []string) error {
		cli, err := prepareCmd(cmd)
		if err != nil {
			return err
		}
		req := &service.HookSetRequest{
			UniqueID:  tools.RandomID(32),
			Timestamp: tools.TimeUnix() + 60,
			HookUrl:   "https://webhook.site/776f9805-40b9-4147-93fb-40c92a6711d3",
		}
		res, err := cli.HookSet(req)
		if err != nil {
			return err
		}
		cmd.Println("Response:", res.Successful)
		return nil
	},
}
