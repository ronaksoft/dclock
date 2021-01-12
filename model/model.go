package model

import (
	"github.com/ronaksoft/rony/repo/kv"
)

/*
   Creation Time: 2021 - Jan - 11
   Created by:  (ehsan)
   Maintainers:
      1.  Ehsan N. Moosa (E2)
   Auditor: Ehsan N. Moosa (E2)
   Copyright Ronak Software Group 2020
*/

//go:generate protoc -I=. -I=../vendor --go_out=paths=source_relative:. model.proto
//go:generate protoc -I=. -I=../vendor --gorony_out=paths=source_relative:. model.proto
func init() {
	kv.MustInit(kv.Config{
		DirPath:             "",
		ConflictRetries:     0,
		ConflictMaxInterval: 0,
	})
}
