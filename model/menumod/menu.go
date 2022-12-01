// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package menumod

import "github.com/xjustloveux/jgo.web/model/jgopostgresqlmod"

type ResSysPro struct {
	Rows []jgopostgresqlmod.SysPro
}

type ResMenu struct {
	Url   string    `json:"url"`
	Name  string    `json:"name"`
	Child []ResMenu `json:"child"`
}
