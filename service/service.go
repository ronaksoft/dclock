package service

import (
	"encoding/binary"
	"github.com/dgraph-io/badger/v2"
	"github.com/ronaksoft/dclock/model"
	"github.com/ronaksoft/rony"
	"github.com/ronaksoft/rony/edge"
	"github.com/ronaksoft/rony/repo/kv"
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
		ClientID:    "",
		ID:          req.GetUniqueID(),
		Timestamp:   req.GetTimestamp(),
		CallbackUrl: req.GetHookUrl(),
		JsonData:    req.GetHookJsonData(),
		Fired:       false,
		Success:     false,
	}
	err := model.SaveHook(h)
	if err != nil {
		ctx.PushError(rony.ErrCodeInternal, err.Error())
		return
	}

	err = kv.Update(func(txn *badger.Txn) error {
		key := make([]byte, 11+len(h.ID))
		copy(key[:3], "CPP")
		binary.BigEndian.PutUint64(key[3:11], uint64(req.Timestamp))
		copy(key[11:], h.ID)
		return txn.Set(key, []byte("OK"))
	})
	if err != nil {
		ctx.PushError(rony.ErrCodeInternal, err.Error())
		return
	}
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
