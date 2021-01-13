package main

import (
	"fmt"
	"github.com/ronaksoft/dclock/service"
	"github.com/ronaksoft/rony/config"
	"github.com/ronaksoft/rony/edgec"
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

	wsC := edgec.NewWebsocket(edgec.WebsocketConfig{
		SeedHostPort: fmt.Sprintf("%s:%d", config.GetString("host"), config.GetInt("port")),
		Header: map[string]string{
			"APIKEY": "",
		},
		Router: nil,
		Secure: false,
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
			UniqueID:  config.GetString("hookID"),
			Timestamp: tools.TimeUnix() + config.GetInt64("delay"),
			HookUrl:   config.GetString("url"),
		}
		res, err := cli.HookSet(req)
		if err != nil {
			return err
		}
		cmd.Println("Response:", res.Successful)
		return nil
	},
}

var HookDeleteCmd = &cobra.Command{
	Use: "HookDelete",
	RunE: func(cmd *cobra.Command, args []string) error {
		cli, err := prepareCmd(cmd)
		if err != nil {
			return err
		}

		res, err := cli.HookDelete(&service.HookDeleteRequest{
			UniqueID: config.GetString("hookID"),
		})
		if err != nil {
			return err
		}
		cmd.Println("Response:", res.Successful)
		return nil
	},
}
