package service

import (
	"github.com/ronaksoft/rony/edge"
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
	panic("implement me")
}

func (c *Clock) HookDelete(ctx *edge.RequestCtx, req *HookDeleteRequest, res *HookDeleteResponse) {
	panic("implement me")
}
