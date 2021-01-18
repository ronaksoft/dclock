package testenv

import (
	"github.com/ronaksoft/rony/repo/kv"
	"os"
)

/*
   Creation Time: 2021 - Jan - 17
   Created by:  (ehsan)
   Maintainers:
      1.  Ehsan N. Moosa (E2)
   Auditor: Ehsan N. Moosa (E2)
   Copyright Ronak Software Group 2020
*/

func Init() {
	kvc := kv.DefaultConfig
	_ = os.MkdirAll(kvc.DirPath, os.ModePerm)
	kv.MustInit(kvc)
}
