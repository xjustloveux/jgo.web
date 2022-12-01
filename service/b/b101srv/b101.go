// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package b101srv

import (
	"github.com/xjustloveux/jgo.web/model/globalmod"
	"github.com/xjustloveux/jgo.web/service/globalsrv"
)

func QueryContent() ([]globalmod.PkgContent, error) {

	return globalsrv.QueryContent("B101")
}
