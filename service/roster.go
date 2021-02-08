package service

import (
	"github.com/ronaksoft/rony/edge"
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

func (r *Roster) PageGet(ctx *edge.RequestCtx, req *PageGetRequest, res *PageSetResponse) {
	panic("implement me")
}

func (r *Roster) PageSet(ctx *edge.RequestCtx, req *PageSetRequest, res *PageSetResponse) {
	panic("implement me")
}
