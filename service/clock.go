package service

import (
	"encoding/binary"
	"github.com/dgraph-io/badger/v3"
	"github.com/ronaksoft/dclock/model"
	"github.com/ronaksoft/rony"
	"github.com/ronaksoft/rony/edge"
	"github.com/ronaksoft/rony/store"
	"hash/crc32"
)

/*
   Creation Time: 2021 - Jan - 11
   Created by:  (ehsan)
   Maintainers:
      1.  Ehsan N. Moosa (E2)
   Auditor: Ehsan N. Moosa (E2)
   Copyright Ronak Software Group 2020
*/

//go:generate protoc -I=. -I=.. -I=../vendor --go_out=paths=source_relative:. clock.proto
//go:generate protoc -I=. -I=.. -I=../vendor --gorony_out=paths=source_relative:. clock.proto
func init() {

}

type Clock struct {
	es *edge.Server
}

func NewClock(es *edge.Server) *Clock {
	return &Clock{
		es: es,
	}
}

func (c *Clock) HookSet(ctx *edge.RequestCtx, req *HookSetRequest, res *HookSetResponse) {
	h := &model.Hook{
		ClientID:    []byte(""),
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

	tReq := &PageGetRequest{
		Page:       crc32.ChecksumIEEE(req.GetUniqueID()),
		ReplicaSet: c.es.Cluster().ReplicaSet(),
		CreateNew:  true,
	}
	tRes := &model.Page{}
	err = ExecuteRemotePageGet(ctx, 1, tReq, tRes)
	if err != nil {
		ctx.PushError(rony.ErrCodeInternal, err.Error())
		return
	}

	// ctxReplica := uint64(crc32.ChecksumIEEE(tools.StrToByte(ctx.GetString("ClientID", "")))%uint32(ctx.Cluster().TotalReplicas()) + 1)
	if tRes.GetReplicaSet() != ctx.Cluster().ReplicaSet() {
		err = ExecuteRemoteHookSet(ctx, tRes.GetReplicaSet(), req, res,
			&rony.KeyValue{
				Key:   "ClientID",
				Value: ctx.GetString("ClientID", ""),
			},
		)
		if err != nil {
			ctx.PushError(rony.ErrCodeInternal, err.Error())
			return
		}
	} else {
		err = store.Update(func(txn *badger.Txn) error {
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
	}

	res.Successful = true
}

func (c *Clock) HookDelete(ctx *edge.RequestCtx, req *HookDeleteRequest, res *HookDeleteResponse) {
	h := &model.Hook{}
	h, err := model.ReadHook(nil, req.GetUniqueID(), h)
	if err != nil {
		ctx.PushError(rony.ErrCodeInternal, err.Error())
		return
	}

	// tReq := &PageGetRequest{
	// 	Page:     crc32.ChecksumIEEE(req.GetUniqueID()),
	// 	ReplicaSet: c.es.Cluster().ReplicaSet(),
	// 	CreateNew: true,
	// }
	// tRes := &model.Page{}
	// err = ExecuteRemotePageGet(ctx, 1, tReq, tRes)
	// if err != nil {
	// 	ctx.PushError(rony.ErrCodeInternal, err.Error())
	// 	return
	// }

	err = model.DeleteHook(nil, req.GetUniqueID())
	if err != nil {
		ctx.PushError(rony.ErrCodeInternal, err.Error())
		return
	}

	res.Successful = true
}
