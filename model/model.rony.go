package model

import (
	badger "github.com/dgraph-io/badger/v2"
	edge "github.com/ronaksoft/rony/edge"
	registry "github.com/ronaksoft/rony/registry"
	kv "github.com/ronaksoft/rony/repo/kv"
	proto "google.golang.org/protobuf/proto"
	sync "sync"
)

const C_Hook int64 = 74116203

type poolHook struct {
	pool sync.Pool
}

func (p *poolHook) Get() *Hook {
	x, ok := p.pool.Get().(*Hook)
	if !ok {
		return &Hook{}
	}
	return x
}

func (p *poolHook) Put(x *Hook) {
	x.ClientID = ""
	x.ID = ""
	x.Timestamp = 0
	x.HookUrl = ""
	x.Fired = false
	x.Success = false
	p.pool.Put(x)
}

var PoolHook = poolHook{}

func init() {
	registry.RegisterConstructor(74116203, "Hook")
}

func (x *Hook) DeepCopy(z *Hook) {
	z.ClientID = x.ClientID
	z.ID = x.ID
	z.Timestamp = x.Timestamp
	z.HookUrl = x.HookUrl
	z.Fired = x.Fired
	z.Success = x.Success
}

func (x *Hook) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_Hook, x)
}

func (x *Hook) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *Hook) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func SaveHook(m *Hook) error {
	alloc := kv.NewAllocator()
	defer alloc.ReleaseAll()
	return kv.Update(func(txn *badger.Txn) error {
		b := alloc.GenValue(m)
		err := txn.Set(alloc.GenKey(C_Hook, m.ClientID, m.ID), b)
		if err != nil {
			return err
		}

		err = txn.Set(alloc.GenKey(C_Hook, 507389681, m.HookUrl, m.ID), b)
		if err != nil {
			return err
		}

		return nil
	})
}

func ReadHook(m *Hook) error {
	alloc := kv.NewAllocator()
	defer alloc.ReleaseAll()
	return kv.View(func(txn *badger.Txn) error {
		item, err := txn.Get(alloc.GenKey(C_Hook, m.ClientID, m.ID))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return m.Unmarshal(val)
		})
	})
}

func ReadHookByHookUrlAndID(m *Hook) error {
	alloc := kv.NewAllocator()
	defer alloc.ReleaseAll()
	return kv.View(func(txn *badger.Txn) error {
		item, err := txn.Get(alloc.GenKey(C_Hook, 507389681, m.HookUrl, m.ID))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return m.Unmarshal(val)
		})
	})
}

func DeleteHook(m *Hook) error {
	alloc := kv.NewAllocator()
	defer alloc.ReleaseAll()
	return kv.Update(func(txn *badger.Txn) error {
		err := txn.Delete(alloc.GenKey(C_Hook, m.ClientID, m.ID))
		if err != nil {
			return err
		}

		err = txn.Delete(alloc.GenKey(C_Hook, 507389681, m.HookUrl, m.ID))
		if err != nil {
			return err
		}

		return nil
	})
}
