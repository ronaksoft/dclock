package service

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/ronaksoft/dclock/model"
	"github.com/ronaksoft/rony"
	"github.com/ronaksoft/rony/edge"
	"github.com/ronaksoft/rony/store"
)

/*
   Creation Time: 2021 - Feb - 08
   Created by:  (ehsan)
   Maintainers:
      1.  Ehsan N. Moosa (E2)
   Auditor: Ehsan N. Moosa (E2)
   Copyright Ronak Software Group 2020
*/

//go:generate protoc -I=. -I=.. -I=../vendor --go_out=paths=source_relative:. roster.proto
//go:generate protoc -I=. -I=.. -I=../vendor --gorony_out=paths=source_relative:. roster.proto
func init() {

}

type Roster struct {
	es *edge.Server
}

func NewRoster(es *edge.Server) *Roster {
	return &Roster{
		es: es,
	}
}

func (r *Roster) PageList(ctx *edge.RequestCtx, req *PageListRequest, res *PagesMany) {
	if r.es.Cluster().ReplicaSet() != 1 {
		err := ExecuteRemotePageList(ctx, 1, req, res)
		if err != nil {
			ctx.PushError(rony.ErrCodeInternal, err.Error())
		}
		return
	}

	// This code will only run on Edge with replica set to 1
	pages, err := model.ListPageByReplicaSet(req.GetReplicaSet(), 0, store.NewListOption())
	if err != nil {
		ctx.PushError(rony.ErrCodeInternal, err.Error())
		return
	}

	for _, p := range pages {
		res.Pages = append(res.Pages, p)
	}
}

func (r *Roster) PageGet(ctx *edge.RequestCtx, req *PageGetRequest, res *model.Page) {
	if r.es.Cluster().ReplicaSet() != 1 {
		ctx.PushError(rony.ErrCodeUnavailable, rony.ErrItemRequest)
		return
	}
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()

	err := store.Update(func(txn *badger.Txn) (err error) {
		_, err = model.ReadPageWithTxn(txn, alloc, req.GetPage(), res)
		if err != nil && !req.GetCreateNew() {
			return err
		}
		res.ReplicaSet = req.GetReplicaSet()
		res.ID = req.GetPage()
		return model.SavePageWithTxn(txn, alloc, res)
	})
	if err != nil {
		ctx.PushError(rony.ErrCodeInternal, err.Error())
		return
	}
}

func (r *Roster) PageSet(ctx *edge.RequestCtx, req *PageSetRequest, res *model.Page) {
	if r.es.Cluster().ReplicaSet() != 1 {
		ctx.PushError(rony.ErrCodeUnavailable, rony.ErrItemRequest)
		return
	}

	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()

	err := store.Update(func(txn *badger.Txn) (err error) {
		_, err = model.ReadPageWithTxn(txn, alloc, req.GetPage(), res)
		if err == nil && !req.GetReplace() {
			return nil
		}
		res.ReplicaSet = req.GetReplicaSet()
		res.ID = req.GetPage()
		return model.SavePageWithTxn(txn, alloc, res)
	})
	if err != nil {
		ctx.PushError(rony.ErrCodeInternal, err.Error())
		return
	}
}
