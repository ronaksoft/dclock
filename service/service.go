package service

import (
	"fmt"
	"github.com/ronaksoft/dclock/model"
	"github.com/ronaksoft/rony"
	"github.com/ronaksoft/rony/edge"
	"github.com/ronaksoft/rony/tools"
	"net/http"
	"time"
)

/*
   Creation Time: 2021 - Jan - 11
   Created by:  (ehsan)
   Maintainers:
      1.  Ehsan N. Moosa (E2)
   Auditor: Ehsan N. Moosa (E2)
   Copyright Ronak Software Group 2020
*/

//go:generate protoc -I=. -I=../vendor --go_out=paths=source_relative:. service.proto
//go:generate protoc -I=. -I=../vendor --gorony_out=paths=source_relative:. service.proto
func init() {

}

type Clock struct{}

func (c *Clock) HookSet(ctx *edge.RequestCtx, req *HookSetRequest, res *HookSetResponse) {
	h := &model.Hook{
		ClientID:  "",
		ID:        req.GetUniqueID(),
		Timestamp: req.GetTimestamp(),
		HookUrl:   req.GetHookUrl(),
		Fired:     false,
		Success:   false,
	}
	err := model.SaveHook(h)
	if err != nil {
		ctx.PushError(rony.ErrCodeInternal, err.Error())
		return
	}

	waitTime := time.Duration(req.GetTimestamp()-tools.TimeUnix()) * time.Second
	go func(hookID string, waitTime time.Duration) {
		time.Sleep(waitTime)
		h := &model.Hook{
			ClientID: "",
			ID:       hookID,
		}
		err = model.ReadHook(h)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Hook", h.ID, h.HookUrl, h.ClientID, h.Fired)
		_, err := http.DefaultClient.Post(h.HookUrl, "application/json", nil)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}(req.GetUniqueID(), waitTime)
	res.Successful = true
}

func (c *Clock) HookDelete(ctx *edge.RequestCtx, req *HookDeleteRequest, res *HookDeleteResponse) {
	h := &model.Hook{
		ID: req.GetUniqueID(),
	}
	err := model.ReadHook(h)
	if err != nil {
		ctx.PushError(rony.ErrCodeInternal, err.Error())
		return
	}

	err = model.DeleteHook(h)
	if err != nil {
		ctx.PushError(rony.ErrCodeInternal, err.Error())
		return
	}

	res.Successful = true
}
