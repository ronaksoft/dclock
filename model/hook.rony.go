// Code generated by Rony's protoc plugin; DO NOT EDIT.

package model

import (
	edge "github.com/ronaksoft/rony/edge"
	registry "github.com/ronaksoft/rony/registry"
	store "github.com/ronaksoft/rony/store"
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
	if x == nil {
		return
	}
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
	if x == nil {
		return
	}
	x.ClientID = x.ClientID[:0]
	x.ID = x.ID[:0]
	PoolHook.Put(x.Hook)
	x.Hook = nil
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

func CreateHook(m *Hook) error {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()
	return store.Update(func(txn *store.Txn) error {
		return CreateHookWithTxn(txn, alloc, m)
	})
}

func CreateHookWithTxn(txn *store.Txn, alloc *store.Allocator, m *Hook) (err error) {
	if alloc == nil {
		alloc = store.NewAllocator()
		defer alloc.ReleaseAll()
	}

	if store.Exists(txn, alloc, 'M', C_Hook, 3973050528, m.ClientID, m.ID) {
		return store.ErrAlreadyExists
	}
	// save entry
	val := alloc.Marshal(m)
	err = store.Set(txn, alloc, val, 'M', C_Hook, 3973050528, m.ClientID, m.ID)
	if err != nil {
		return
	}

	// save views
	// save entry for view: [CallbackUrl ID]
	err = store.Set(txn, alloc, val, 'M', C_Hook, 3467894716, m.CallbackUrl, m.ID)
	if err != nil {
		return
	}

	return

}

func ReadHookWithTxn(txn *store.Txn, alloc *store.Allocator, clientID []byte, id []byte, m *Hook) (*Hook, error) {
	if alloc == nil {
		alloc = store.NewAllocator()
		defer alloc.ReleaseAll()
	}

	err := store.Unmarshal(txn, alloc, m, 'M', C_Hook, 3973050528, clientID, id)
	if err != nil {
		return nil, err
	}
	return m, err
}

func ReadHook(clientID []byte, id []byte, m *Hook) (*Hook, error) {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()

	if m == nil {
		m = &Hook{}
	}

	err := store.View(func(txn *store.Txn) (err error) {
		m, err = ReadHookWithTxn(txn, alloc, clientID, id, m)
		return err
	})
	return m, err
}

func ReadHookByCallbackUrlAndIDWithTxn(txn *store.Txn, alloc *store.Allocator, callbackUrl []byte, id []byte, m *Hook) (*Hook, error) {
	if alloc == nil {
		alloc = store.NewAllocator()
		defer alloc.ReleaseAll()
	}

	err := store.Unmarshal(txn, alloc, m, 'M', C_Hook, 3467894716, callbackUrl, id)
	if err != nil {
		return nil, err
	}
	return m, err
}

func ReadHookByCallbackUrlAndID(callbackUrl []byte, id []byte, m *Hook) (*Hook, error) {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()
	if m == nil {
		m = &Hook{}
	}
	err := store.View(func(txn *store.Txn) (err error) {
		m, err = ReadHookByCallbackUrlAndIDWithTxn(txn, alloc, callbackUrl, id, m)
		return err
	})
	return m, err
}

func UpdateHookWithTxn(txn *store.Txn, alloc *store.Allocator, m *Hook) error {
	if alloc == nil {
		alloc = store.NewAllocator()
		defer alloc.ReleaseAll()
	}

	om := &Hook{}
	err := store.Unmarshal(txn, alloc, om, 'M', C_Hook, 3973050528, m.ClientID, m.ID)
	if err != nil {
		return err
	}

	err = DeleteHookWithTxn(txn, alloc, om.ClientID, om.ID)
	if err != nil {
		return err
	}

	return CreateHookWithTxn(txn, alloc, m)
}

func UpdateHook(clientID []byte, id []byte, m *Hook) error {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()

	if m == nil {
		return store.ErrEmptyObject
	}

	err := store.View(func(txn *store.Txn) (err error) {
		return UpdateHookWithTxn(txn, alloc, m)
	})
	return err
}

func DeleteHookWithTxn(txn *store.Txn, alloc *store.Allocator, clientID []byte, id []byte) error {
	m := &Hook{}
	err := store.Unmarshal(txn, alloc, m, 'M', C_Hook, 3973050528, clientID, id)
	if err != nil {
		return err
	}
	err = store.Delete(txn, alloc, 'M', C_Hook, 3973050528, m.ClientID, m.ID)
	if err != nil {
		return err
	}

	err = store.Delete(txn, alloc, 'M', C_Hook, 3467894716, m.CallbackUrl, m.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteHook(clientID []byte, id []byte) error {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()

	return store.Update(func(txn *store.Txn) error {
		return DeleteHookWithTxn(txn, alloc, clientID, id)
	})
}

func SaveHookWithTxn(txn *store.Txn, alloc *store.Allocator, m *Hook) (err error) {
	if store.Exists(txn, alloc, 'M', C_Hook, 3973050528, m.ClientID, m.ID) {
		return UpdateHookWithTxn(txn, alloc, m)
	} else {
		return CreateHookWithTxn(txn, alloc, m)
	}
}

func SaveHook(m *Hook) error {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()
	return store.Update(func(txn *store.Txn) error {
		return SaveHookWithTxn(txn, alloc, m)
	})
}

func IterHooks(txn *store.Txn, alloc *store.Allocator, cb func(m *Hook) bool) error {
	if alloc == nil {
		alloc = store.NewAllocator()
		defer alloc.ReleaseAll()
	}

	exitLoop := false
	iterOpt := store.DefaultIteratorOptions
	iterOpt.Prefix = alloc.Gen('M', C_Hook, 3973050528)
	iter := txn.NewIterator(iterOpt)
	for iter.Rewind(); iter.ValidForPrefix(iterOpt.Prefix); iter.Next() {
		_ = iter.Item().Value(func(val []byte) error {
			m := &Hook{}
			err := m.Unmarshal(val)
			if err != nil {
				return err
			}
			if !cb(m) {
				exitLoop = true
			}
			return nil
		})
		if exitLoop {
			break
		}
	}
	iter.Close()
	return nil
}

func ListHook(
	offsetClientID []byte, offsetID []byte, lo *store.ListOption, cond func(m *Hook) bool,
) ([]*Hook, error) {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()

	res := make([]*Hook, 0, lo.Limit())
	err := store.View(func(txn *store.Txn) error {
		opt := store.DefaultIteratorOptions
		opt.Prefix = alloc.Gen('M', C_Hook, 3973050528)
		opt.Reverse = lo.Backward()
		osk := alloc.Gen('M', C_Hook, 3973050528, offsetClientID)
		iter := txn.NewIterator(opt)
		offset := lo.Skip()
		limit := lo.Limit()
		for iter.Seek(osk); iter.ValidForPrefix(opt.Prefix); iter.Next() {
			if offset--; offset >= 0 {
				continue
			}
			if limit--; limit < 0 {
				break
			}
			_ = iter.Item().Value(func(val []byte) error {
				m := &Hook{}
				err := m.Unmarshal(val)
				if err != nil {
					return err
				}
				if cond == nil || cond(m) {
					res = append(res, m)
				}
				return nil
			})
		}
		iter.Close()
		return nil
	})
	return res, err
}

func IterHookByClientID(txn *store.Txn, alloc *store.Allocator, clientID []byte, cb func(m *Hook) bool) error {
	if alloc == nil {
		alloc = store.NewAllocator()
		defer alloc.ReleaseAll()
	}

	exitLoop := false
	opt := store.DefaultIteratorOptions
	opt.Prefix = alloc.Gen('M', C_Hook, 3973050528, clientID)
	iter := txn.NewIterator(opt)
	for iter.Rewind(); iter.ValidForPrefix(opt.Prefix); iter.Next() {
		_ = iter.Item().Value(func(val []byte) error {
			m := &Hook{}
			err := m.Unmarshal(val)
			if err != nil {
				return err
			}
			if !cb(m) {
				exitLoop = true
			}
			return nil
		})
		if exitLoop {
			break
		}
	}
	iter.Close()
	return nil
}

func IterHookByCallbackUrl(txn *store.Txn, alloc *store.Allocator, callbackUrl []byte, cb func(m *Hook) bool) error {
	if alloc == nil {
		alloc = store.NewAllocator()
		defer alloc.ReleaseAll()
	}

	exitLoop := false
	opt := store.DefaultIteratorOptions
	opt.Prefix = alloc.Gen('M', C_Hook, 3467894716, callbackUrl)
	iter := txn.NewIterator(opt)
	for iter.Rewind(); iter.ValidForPrefix(opt.Prefix); iter.Next() {
		_ = iter.Item().Value(func(val []byte) error {
			m := &Hook{}
			err := m.Unmarshal(val)
			if err != nil {
				return err
			}
			if !cb(m) {
				exitLoop = true
			}
			return nil
		})
		if exitLoop {
			break
		}
	}
	iter.Close()
	return nil
}

func ListHookByClientID(clientID []byte, offsetID []byte, lo *store.ListOption) ([]*Hook, error) {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()

	res := make([]*Hook, 0, lo.Limit())
	err := store.View(func(txn *store.Txn) error {
		opt := store.DefaultIteratorOptions
		opt.Prefix = alloc.Gen('M', C_Hook, 3973050528, clientID)
		opt.Reverse = lo.Backward()
		osk := alloc.Gen('M', C_Hook, 3973050528, clientID, offsetID)
		iter := txn.NewIterator(opt)
		offset := lo.Skip()
		limit := lo.Limit()
		for iter.Seek(osk); iter.ValidForPrefix(opt.Prefix); iter.Next() {
			if offset--; offset >= 0 {
				continue
			}
			if limit--; limit < 0 {
				break
			}
			_ = iter.Item().Value(func(val []byte) error {
				m := &Hook{}
				err := m.Unmarshal(val)
				if err != nil {
					return err
				}
				res = append(res, m)
				return nil
			})
		}
		iter.Close()
		return nil
	})
	return res, err
}

func ListHookByCallbackUrl(callbackUrl []byte, offsetID []byte, lo *store.ListOption) ([]*Hook, error) {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()

	res := make([]*Hook, 0, lo.Limit())
	err := store.View(func(txn *store.Txn) error {
		opt := store.DefaultIteratorOptions
		opt.Prefix = alloc.Gen('M', C_Hook, 3467894716, callbackUrl)
		opt.Reverse = lo.Backward()
		osk := alloc.Gen('M', C_Hook, 3467894716, callbackUrl, offsetID)
		iter := txn.NewIterator(opt)
		offset := lo.Skip()
		limit := lo.Limit()
		for iter.Seek(osk); iter.ValidForPrefix(opt.Prefix); iter.Next() {
			if offset--; offset >= 0 {
				continue
			}
			if limit--; limit < 0 {
				break
			}
			_ = iter.Item().Value(func(val []byte) error {
				m := &Hook{}
				err := m.Unmarshal(val)
				if err != nil {
					return err
				}
				res = append(res, m)
				return nil
			})
		}
		iter.Close()
		return nil
	})
	return res, err
}

func CreateHookHolder(m *HookHolder) error {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()
	return store.Update(func(txn *store.Txn) error {
		return CreateHookHolderWithTxn(txn, alloc, m)
	})
}

func CreateHookHolderWithTxn(txn *store.Txn, alloc *store.Allocator, m *HookHolder) (err error) {
	if alloc == nil {
		alloc = store.NewAllocator()
		defer alloc.ReleaseAll()
	}

	if store.Exists(txn, alloc, 'M', C_HookHolder, 3973050528, m.ClientID, m.ID) {
		return store.ErrAlreadyExists
	}
	// save entry
	val := alloc.Marshal(m)
	err = store.Set(txn, alloc, val, 'M', C_HookHolder, 3973050528, m.ClientID, m.ID)
	if err != nil {
		return
	}

	return

}

func ReadHookHolderWithTxn(txn *store.Txn, alloc *store.Allocator, clientID []byte, id []byte, m *HookHolder) (*HookHolder, error) {
	if alloc == nil {
		alloc = store.NewAllocator()
		defer alloc.ReleaseAll()
	}

	err := store.Unmarshal(txn, alloc, m, 'M', C_HookHolder, 3973050528, clientID, id)
	if err != nil {
		return nil, err
	}
	return m, err
}

func ReadHookHolder(clientID []byte, id []byte, m *HookHolder) (*HookHolder, error) {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()

	if m == nil {
		m = &HookHolder{}
	}

	err := store.View(func(txn *store.Txn) (err error) {
		m, err = ReadHookHolderWithTxn(txn, alloc, clientID, id, m)
		return err
	})
	return m, err
}

func UpdateHookHolderWithTxn(txn *store.Txn, alloc *store.Allocator, m *HookHolder) error {
	if alloc == nil {
		alloc = store.NewAllocator()
		defer alloc.ReleaseAll()
	}

	om := &HookHolder{}
	err := store.Unmarshal(txn, alloc, om, 'M', C_HookHolder, 3973050528, m.ClientID, m.ID)
	if err != nil {
		return err
	}

	err = DeleteHookHolderWithTxn(txn, alloc, om.ClientID, om.ID)
	if err != nil {
		return err
	}

	return CreateHookHolderWithTxn(txn, alloc, m)
}

func UpdateHookHolder(clientID []byte, id []byte, m *HookHolder) error {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()

	if m == nil {
		return store.ErrEmptyObject
	}

	err := store.View(func(txn *store.Txn) (err error) {
		return UpdateHookHolderWithTxn(txn, alloc, m)
	})
	return err
}

func DeleteHookHolderWithTxn(txn *store.Txn, alloc *store.Allocator, clientID []byte, id []byte) error {
	err := store.Delete(txn, alloc, 'M', C_HookHolder, 3973050528, clientID, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteHookHolder(clientID []byte, id []byte) error {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()

	return store.Update(func(txn *store.Txn) error {
		return DeleteHookHolderWithTxn(txn, alloc, clientID, id)
	})
}

func SaveHookHolderWithTxn(txn *store.Txn, alloc *store.Allocator, m *HookHolder) (err error) {
	if store.Exists(txn, alloc, 'M', C_HookHolder, 3973050528, m.ClientID, m.ID) {
		return UpdateHookHolderWithTxn(txn, alloc, m)
	} else {
		return CreateHookHolderWithTxn(txn, alloc, m)
	}
}

func SaveHookHolder(m *HookHolder) error {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()
	return store.Update(func(txn *store.Txn) error {
		return SaveHookHolderWithTxn(txn, alloc, m)
	})
}

func IterHookHolders(txn *store.Txn, alloc *store.Allocator, cb func(m *HookHolder) bool) error {
	if alloc == nil {
		alloc = store.NewAllocator()
		defer alloc.ReleaseAll()
	}

	exitLoop := false
	iterOpt := store.DefaultIteratorOptions
	iterOpt.Prefix = alloc.Gen('M', C_HookHolder, 3973050528)
	iter := txn.NewIterator(iterOpt)
	for iter.Rewind(); iter.ValidForPrefix(iterOpt.Prefix); iter.Next() {
		_ = iter.Item().Value(func(val []byte) error {
			m := &HookHolder{}
			err := m.Unmarshal(val)
			if err != nil {
				return err
			}
			if !cb(m) {
				exitLoop = true
			}
			return nil
		})
		if exitLoop {
			break
		}
	}
	iter.Close()
	return nil
}

func ListHookHolder(
	offsetClientID []byte, offsetID []byte, lo *store.ListOption, cond func(m *HookHolder) bool,
) ([]*HookHolder, error) {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()

	res := make([]*HookHolder, 0, lo.Limit())
	err := store.View(func(txn *store.Txn) error {
		opt := store.DefaultIteratorOptions
		opt.Prefix = alloc.Gen('M', C_HookHolder, 3973050528)
		opt.Reverse = lo.Backward()
		osk := alloc.Gen('M', C_HookHolder, 3973050528, offsetClientID)
		iter := txn.NewIterator(opt)
		offset := lo.Skip()
		limit := lo.Limit()
		for iter.Seek(osk); iter.ValidForPrefix(opt.Prefix); iter.Next() {
			if offset--; offset >= 0 {
				continue
			}
			if limit--; limit < 0 {
				break
			}
			_ = iter.Item().Value(func(val []byte) error {
				m := &HookHolder{}
				err := m.Unmarshal(val)
				if err != nil {
					return err
				}
				if cond == nil || cond(m) {
					res = append(res, m)
				}
				return nil
			})
		}
		iter.Close()
		return nil
	})
	return res, err
}

func IterHookHolderByClientID(txn *store.Txn, alloc *store.Allocator, clientID []byte, cb func(m *HookHolder) bool) error {
	if alloc == nil {
		alloc = store.NewAllocator()
		defer alloc.ReleaseAll()
	}

	exitLoop := false
	opt := store.DefaultIteratorOptions
	opt.Prefix = alloc.Gen('M', C_HookHolder, 3973050528, clientID)
	iter := txn.NewIterator(opt)
	for iter.Rewind(); iter.ValidForPrefix(opt.Prefix); iter.Next() {
		_ = iter.Item().Value(func(val []byte) error {
			m := &HookHolder{}
			err := m.Unmarshal(val)
			if err != nil {
				return err
			}
			if !cb(m) {
				exitLoop = true
			}
			return nil
		})
		if exitLoop {
			break
		}
	}
	iter.Close()
	return nil
}

func ListHookHolderByClientID(clientID []byte, offsetID []byte, lo *store.ListOption) ([]*HookHolder, error) {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()

	res := make([]*HookHolder, 0, lo.Limit())
	err := store.View(func(txn *store.Txn) error {
		opt := store.DefaultIteratorOptions
		opt.Prefix = alloc.Gen('M', C_HookHolder, 3973050528, clientID)
		opt.Reverse = lo.Backward()
		osk := alloc.Gen('M', C_HookHolder, 3973050528, clientID, offsetID)
		iter := txn.NewIterator(opt)
		offset := lo.Skip()
		limit := lo.Limit()
		for iter.Seek(osk); iter.ValidForPrefix(opt.Prefix); iter.Next() {
			if offset--; offset >= 0 {
				continue
			}
			if limit--; limit < 0 {
				break
			}
			_ = iter.Item().Value(func(val []byte) error {
				m := &HookHolder{}
				err := m.Unmarshal(val)
				if err != nil {
					return err
				}
				res = append(res, m)
				return nil
			})
		}
		iter.Close()
		return nil
	})
	return res, err
}
