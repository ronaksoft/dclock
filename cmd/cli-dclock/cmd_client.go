package main

import (
	"github.com/c-bata/go-prompt"
	"github.com/ronaksoft/dclock/service"
	"github.com/ronaksoft/rony"
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
		// Bind the current flags to registered flags in config package
		err := config.BindCmdFlags(cmd)
		if err != nil {
			return err
		}

		// Register Client Commands
		c := edgec.NewHttp(edgec.HttpConfig{
			SeedHostPort: config.GetString("hostPort"),
		})
		err = c.Start()
		if err != nil {
			return err
		}

		cli := &ClockCli{}
		service.RegisterClockCli(cli, c, ClientInteractiveCmd)

		p := prompt.New(tools.PromptExecutor(ClientInteractiveCmd), tools.PromptCompleter(ClientInteractiveCmd))
		p.Run()
		return nil
	},
}

var ClientInteractiveCmd = &cobra.Command{
	Use: "interactive",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

type ClockCli struct{}

func (c *ClockCli) HookSet(cli *service.ClockClient, cmd *cobra.Command, args []string) error {
	req := &service.HookSetRequest{
		UniqueID:  tools.StrToByte(config.GetString("uniqueID")),
		Timestamp: tools.TimeUnix() + config.GetInt64("timestamp"),
		HookUrl:   tools.StrToByte(config.GetString("hookUrl")),
	}
	cmd.Println("Req:", req.String())
	res, err := cli.HookSet(
		req,
		&rony.KeyValue{
			Key:   "ClientID",
			Value: config.GetString("clientID"),
		},
	)
	if err != nil {
		cmd.Println("Error:", err)
		return nil
	}
	cmd.Println("Response:", res.Successful)
	return nil
}

func (c *ClockCli) HookDelete(cli *service.ClockClient, cmd *cobra.Command, args []string) error {
	req := &service.HookDeleteRequest{
		UniqueID: tools.StrToByte(config.GetString("hookID")),
	}
	cmd.Println("Req:", req.String())
	res, err := cli.HookDelete(
		req,
		&rony.KeyValue{
			Key:   "ClientID",
			Value: config.GetString("clientID"),
		},
	)
	if err != nil {
		cmd.Println("Error:", err)
		return nil
	}
	cmd.Println("Response:", res.Successful)
	return nil
}
