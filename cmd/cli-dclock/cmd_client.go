package main

import (
	"github.com/ronaksoft/dclock/service"
	"github.com/ronaksoft/rony"
	"github.com/ronaksoft/rony/config"
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
}

type ClockCli struct{}

func (c *ClockCli) HookSet(cli *service.ClockClient, cmd *cobra.Command, args []string) error {
	req := &service.HookSetRequest{
		UniqueID:  tools.StrToByte(config.GetString("hookID")),
		Timestamp: tools.TimeUnix() + config.GetInt64("delay"),
		HookUrl:   tools.StrToByte(config.GetString("url")),
	}
	res, err := cli.HookSet(
		req,
		&rony.KeyValue{
			Key:   "ClientID",
			Value: config.GetString("clientID"),
		},
	)
	if err != nil {
		return err
	}
	cmd.Println("Response:", res.Successful)
	return nil
}

func (c *ClockCli) HookDelete(cli *service.ClockClient, cmd *cobra.Command, args []string) error {
	req := &service.HookDeleteRequest{
		UniqueID: tools.StrToByte(config.GetString("hookID")),
	}
	res, err := cli.HookDelete(
		req,
		&rony.KeyValue{
			Key:   "ClientID",
			Value: config.GetString("clientID"),
		},
	)
	if err != nil {
		return err
	}
	cmd.Println("Response:", res.Successful)
	return nil
}
