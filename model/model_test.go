package model_test

import (
	"github.com/ronaksoft/dclock/internal/testenv"
	"github.com/ronaksoft/dclock/model"
	"github.com/ronaksoft/rony/tools"
	. "github.com/smartystreets/goconvey/convey"
	"runtime"
	"testing"
)

/*
   Creation Time: 2021 - Jan - 17
   Created by:  (ehsan)
   Maintainers:
      1.  Ehsan N. Moosa (E2)
   Auditor: Ehsan N. Moosa (E2)
   Copyright Ronak Software Group 2020
*/

func init() {
	testenv.Init()
	runtime.MemProfileRate = 0
}

func BenchmarkSaveHook(b *testing.B) {
	clientID := "ClientID"
	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			h := model.PoolHook.Get()
			h.ID = tools.StrToByte(tools.RandomID(10))
			h.ClientID = tools.StrToByte(clientID)
			h.Timestamp = tools.TimeUnix()
			h.CallbackUrl = tools.StrToByte("http://callback")
			err := model.SaveHook(h)
			if err != nil {
				b.Fatal(err)
			}
			model.PoolHook.Put(h)
		}
	})
}

func TestHook(t *testing.T) {
	Convey("Hook", t, func(c C) {
		h := &model.Hook{
			ClientID:    tools.StrToByte("SomeClientID"),
			ID:          tools.StrToByte(tools.RandomID(32)),
			Timestamp:   tools.TimeUnix(),
			CallbackUrl: tools.StrToByte("https://random"),
			JsonData:    nil,
			Fired:       false,
			Success:     false,
		}
		err := model.SaveHook(h)
		c.So(err, ShouldBeNil)
		h2 := &model.Hook{}
		h2, err = model.ReadHook(h.GetClientID(), h.GetID(), h2)
		c.So(err, ShouldBeNil)
		c.So(h2.Timestamp, ShouldEqual, h.Timestamp)
		c.So(h2.CallbackUrl, ShouldEqual, h.CallbackUrl)

		h3 := &model.Hook{}
		h3, err = model.ReadHookByCallbackUrlAndID(h.GetCallbackUrl(), h.GetID(), h3)
		c.So(err, ShouldBeNil)
		c.So(h3.ClientID, ShouldEqual, h.ClientID)
		c.So(h3.Timestamp, ShouldEqual, h.Timestamp)

		err = model.DeleteHook(h.GetClientID(), h.GetID())
		c.So(err, ShouldBeNil)
		h.Reset()
		h2, err = model.ReadHook(h.GetClientID(), h.GetID(), h2)
		c.So(err, ShouldNotBeNil)

	})
}
