package testenv

import (
	"github.com/ronaksoft/rony/store"
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
	kvc := store.DefaultConfig("./hdd")
	_ = os.MkdirAll(kvc.DirPath, os.ModePerm)
	store.MustInit(kvc)
}
