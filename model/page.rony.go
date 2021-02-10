// Code generated by Rony's protoc plugin; DO NOT EDIT.

package model

import (
	badger "github.com/dgraph-io/badger/v3"
	edge "github.com/ronaksoft/rony/edge"
	registry "github.com/ronaksoft/rony/registry"
	store "github.com/ronaksoft/rony/store"
	proto "google.golang.org/protobuf/proto"
	sync "sync"
)

const C_Page int64 = 3023575326

type poolPage struct {
	pool sync.Pool
}

func (p *poolPage) Get() *Page {
	x, ok := p.pool.Get().(*Page)
	if !ok {
		return &Page{}
	}
	return x
}

func (p *poolPage) Put(x *Page) {
	x.ID = 0
	x.ReplicaSet = 0
	p.pool.Put(x)
}

var PoolPage = poolPage{}

func (x *Page) DeepCopy(z *Page) {
	z.ID = x.ID
	z.ReplicaSet = x.ReplicaSet
}

func (x *Page) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *Page) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *Page) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_Page, x)
}

func init() {
	registry.RegisterConstructor(3023575326, "Page")
}

func SavePageWithTxn(txn *badger.Txn, alloc *store.Allocator, m *Page) (err error) {
	if alloc == nil {
		alloc = store.NewAllocator()
		defer alloc.ReleaseAll()
	}

	// save entry
	b := alloc.GenValue(m)
	key := alloc.GenKey('M', C_Page, 299066170, m.ID)
	err = txn.Set(key, b)
	if err != nil {
		return
	}

	// save entry for view[ReplicaSet ID]
	err = txn.Set(alloc.GenKey('M', C_Page, 1040696757, m.ReplicaSet, m.ID), b)
	if err != nil {
		return
	}

	return

}

func SavePage(m *Page) error {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()
	return store.Update(func(txn *badger.Txn) error {
		return SavePageWithTxn(txn, alloc, m)
	})
}

func ReadPageWithTxn(txn *badger.Txn, alloc *store.Allocator, id uint32, m *Page) (*Page, error) {
	if alloc == nil {
		alloc = store.NewAllocator()
		defer alloc.ReleaseAll()
	}

	item, err := txn.Get(alloc.GenKey('M', C_Page, 299066170, id))
	if err != nil {
		return nil, err
	}
	err = item.Value(func(val []byte) error {
		return m.Unmarshal(val)
	})
	return m, err
}

func ReadPage(id uint32, m *Page) (*Page, error) {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()

	if m == nil {
		m = &Page{}
	}

	err := store.View(func(txn *badger.Txn) (err error) {
		m, err = ReadPageWithTxn(txn, alloc, id, m)
		return err
	})
	return m, err
}

func ReadPageByReplicaSetAndIDWithTxn(txn *badger.Txn, alloc *store.Allocator, replicaSet uint64, id uint32, m *Page) (*Page, error) {
	if alloc == nil {
		alloc = store.NewAllocator()
		defer alloc.ReleaseAll()
	}

	item, err := txn.Get(alloc.GenKey('M', C_Page, 1040696757, replicaSet, id))
	if err != nil {
		return nil, err
	}
	err = item.Value(func(val []byte) error {
		return m.Unmarshal(val)
	})
	return m, err
}

func ReadPageByReplicaSetAndID(replicaSet uint64, id uint32, m *Page) (*Page, error) {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()
	if m == nil {
		m = &Page{}
	}
	err := store.View(func(txn *badger.Txn) (err error) {
		m, err = ReadPageByReplicaSetAndIDWithTxn(txn, alloc, replicaSet, id, m)
		return err
	})
	return m, err
}

func DeletePageWithTxn(txn *badger.Txn, alloc *store.Allocator, id uint32) error {
	m := &Page{}
	item, err := txn.Get(alloc.GenKey('M', C_Page, 299066170, id))
	if err != nil {
		return err
	}
	err = item.Value(func(val []byte) error {
		return m.Unmarshal(val)
	})
	if err != nil {
		return err
	}
	err = txn.Delete(alloc.GenKey('M', C_Page, 299066170, m.ID))
	if err != nil {
		return err
	}

	err = txn.Delete(alloc.GenKey('M', C_Page, 1040696757, m.ReplicaSet, m.ID))
	if err != nil {
		return err
	}

	return nil
}

func DeletePage(id uint32) error {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()

	return store.Update(func(txn *badger.Txn) error {
		return DeletePageWithTxn(txn, alloc, id)
	})
}

func ListPage(
	offsetID uint32, lo *store.ListOption, cond func(m *Page) bool,
) ([]*Page, error) {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()

	res := make([]*Page, 0, lo.Limit())
	err := store.View(func(txn *badger.Txn) error {
		opt := badger.DefaultIteratorOptions
		opt.Prefix = alloc.GenKey(C_Page, 299066170)
		opt.Reverse = lo.Backward()
		osk := alloc.GenKey('M', C_Page, 299066170, offsetID)
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
				m := &Page{}
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

func IterPages(txn *badger.Txn, alloc *store.Allocator, cb func(m *Page) bool) error {
	if alloc == nil {
		alloc = store.NewAllocator()
		defer alloc.ReleaseAll()
	}

	exitLoop := false
	iterOpt := badger.DefaultIteratorOptions
	iterOpt.Prefix = alloc.GenKey(C_Page, 299066170)
	iter := txn.NewIterator(iterOpt)
	for iter.Rewind(); iter.ValidForPrefix(iterOpt.Prefix); iter.Next() {
		_ = iter.Item().Value(func(val []byte) error {
			m := &Page{}
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

func ListPageByReplicaSet(replicaSet uint64, offsetID uint32, lo *store.ListOption) ([]*Page, error) {
	alloc := store.NewAllocator()
	defer alloc.ReleaseAll()

	res := make([]*Page, 0, lo.Limit())
	err := store.View(func(txn *badger.Txn) error {
		opt := badger.DefaultIteratorOptions
		opt.Prefix = alloc.GenKey('M', C_Page, 1040696757, replicaSet)
		opt.Reverse = lo.Backward()
		osk := alloc.GenKey('M', C_Page, 1040696757, replicaSet, offsetID)
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
				m := &Page{}
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

func IterPageByReplicaSet(txn *badger.Txn, alloc *store.Allocator, replicaSet uint64, cb func(m *Page) bool) error {
	if alloc == nil {
		alloc = store.NewAllocator()
		defer alloc.ReleaseAll()
	}

	exitLoop := false
	opt := badger.DefaultIteratorOptions
	opt.Prefix = alloc.GenKey('M', C_Page, 1040696757, replicaSet)
	iter := txn.NewIterator(opt)
	for iter.Rewind(); iter.ValidForPrefix(opt.Prefix); iter.Next() {
		_ = iter.Item().Value(func(val []byte) error {
			m := &Page{}
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
