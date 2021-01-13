package main

import (
	"encoding/binary"
	"fmt"
	"github.com/dgraph-io/badger/v2"
	"github.com/ronaksoft/dclock/model"
	"github.com/ronaksoft/rony/pools"
	"github.com/ronaksoft/rony/repo/kv"
	"github.com/ronaksoft/rony/tools"
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

var (
	CheckPointKey    = []byte("CPK")
	CheckPointPrefix = []byte("CPP")
)

type Executor struct {
	rateLimitChan chan struct{}
	checkPoint    int64
	workerFunc    func(h *model.Hook)
}

func NewExecutor(concurrency int, workerFunc func(h *model.Hook)) *Executor {
	e := &Executor{
		rateLimitChan: make(chan struct{}, concurrency),
		workerFunc:    workerFunc,
	}
	e.loadCheckPoint()
	if e.checkPoint == 0 {
		e.checkPoint = tools.TimeUnix()
		e.saveCheckPoint()
	}
	return e
}

func (e *Executor) Start() {
	go e.Run()
}

func (e *Executor) Run() {
	e.watcher()
}

func (e *Executor) watcher() {
	for {
		if e.checkPoint >= tools.TimeUnix() {
			time.Sleep(time.Second)
			continue
		}
		e.runCheckPoint()
		e.checkPoint += 1
		e.saveCheckPoint()
	}
}

func (e *Executor) saveCheckPoint() {
	var b [8]byte
	_ = kv.Update(func(txn *badger.Txn) error {
		binary.BigEndian.PutUint64(b[:], uint64(e.checkPoint))
		return txn.Set(CheckPointKey, b[:])
	})
}

func (e *Executor) loadCheckPoint() {
	_ = kv.View(func(txn *badger.Txn) error {
		item, err := txn.Get(CheckPointKey)
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			e.checkPoint = int64(binary.BigEndian.Uint64(val))
			return nil
		})
	})
	return
}

func (e *Executor) runCheckPoint() {
	var checkPointPrefix [11]byte
	copy(CheckPointPrefix[:3], CheckPointPrefix)
	binary.BigEndian.PutUint64(CheckPointPrefix[3:], uint64(e.checkPoint))
	err := kv.View(func(txn *badger.Txn) error {
		opt := badger.DefaultIteratorOptions
		opt.Prefix = CheckPointPrefix
		opt.PrefetchValues = true
		iter := txn.NewIterator(opt)
		for iter.Seek(checkPointPrefix[:]); iter.Valid(); iter.Next() {
			h := &model.Hook{
				ID: string(iter.Item().Key()[11:]),
			}
			err := model.ReadHook(h)
			if err != nil {
				fmt.Println("err on reading hook:", err)
				continue
			}
			e.rateLimitChan <- struct{}{}
			pools.Go(func() {
				e.workerFunc(h)
				<-e.rateLimitChan
			})
		}

		return nil
	})
	if err != nil {
		fmt.Println("error on running checkpoint:", err)
	}
}
