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
	x.ClientID = x.ClientID[:0]
	x.ID = x.ID[:0]
	x.Timestamp = 0
	x.CallbackUrl = x.CallbackUrl[:0]
	x.JsonData = x.JsonData[:0]
	x.Fired = false
	x.Success = false
	p.pool.Put(x)
}

var PoolHook = poolHook{}

func (x *Hook) DeepCopy(z *Hook) {
	z.ClientID = append(z.ClientID[:0], x.ClientID...)
	z.ID = append(z.ID[:0], x.ID...)
	z.Timestamp = x.Timestamp
	z.CallbackUrl = append(z.CallbackUrl[:0], x.CallbackUrl...)
	z.JsonData = append(z.JsonData[:0], x.JsonData...)
	z.Fired = x.Fired
	z.Success = x.Success
}

func (x *Hook) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *Hook) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *Hook) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_Hook, x)
}

const C_HookHolder int64 = 226559863

type poolHookHolder struct {
	pool sync.Pool
}

func (p *poolHookHolder) Get() *HookHolder {
	x, ok := p.pool.Get().(*HookHolder)
	if !ok {
		return &HookHolder{}
	}
	return x
}

func (p *poolHookHolder) Put(x *HookHolder) {
	x.ClientID = x.ClientID[:0]
	x.ID = x.ID[:0]
	if x.Hook != nil {
		PoolHook.Put(x.Hook)
		x.Hook = nil
	}
	p.pool.Put(x)
}

var PoolHookHolder = poolHookHolder{}

func (x *HookHolder) DeepCopy(z *HookHolder) {
	z.ClientID = append(z.ClientID[:0], x.ClientID...)
	z.ID = append(z.ID[:0], x.ID...)
	if x.Hook != nil {
		z.Hook = PoolHook.Get()
		x.Hook.DeepCopy(z.Hook)
	}
}

func (x *HookHolder) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *HookHolder) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *HookHolder) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_HookHolder, x)
}

func init() {
	registry.RegisterConstructor(74116203, "Hook")
	registry.RegisterConstructor(226559863, "HookHolder")
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

		err = txn.Set(alloc.GenKey(C_Hook, 3583556648, m.CallbackUrl, m.ID), b)
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

func ReadHookByCallbackUrlAndID(m *Hook) error {
	alloc := kv.NewAllocator()
	defer alloc.ReleaseAll()
	return kv.View(func(txn *badger.Txn) error {
		item, err := txn.Get(alloc.GenKey(C_Hook, 3583556648, m.CallbackUrl, m.ID))
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

		err = txn.Delete(alloc.GenKey(C_Hook, 3583556648, m.CallbackUrl, m.ID))
		if err != nil {
			return err
		}

		return nil
	})
}

func SaveHookHolder(m *HookHolder) error {
	alloc := kv.NewAllocator()
	defer alloc.ReleaseAll()
	return kv.Update(func(txn *badger.Txn) error {
		b := alloc.GenValue(m)
		err := txn.Set(alloc.GenKey(C_HookHolder, m.ClientID, m.ID), b)
		if err != nil {
			return err
		}

		return nil
	})
}

func ReadHookHolder(m *HookHolder) error {
	alloc := kv.NewAllocator()
	defer alloc.ReleaseAll()
	return kv.View(func(txn *badger.Txn) error {
		item, err := txn.Get(alloc.GenKey(C_HookHolder, m.ClientID, m.ID))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return m.Unmarshal(val)
		})
	})
}

func DeleteHookHolder(m *HookHolder) error {
	alloc := kv.NewAllocator()
	defer alloc.ReleaseAll()
	return kv.Update(func(txn *badger.Txn) error {
		err := txn.Delete(alloc.GenKey(C_HookHolder, m.ClientID, m.ID))
		if err != nil {
			return err
		}

		return nil
	})
}
