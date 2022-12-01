// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package b102mod

import "github.com/xjustloveux/jgo.web/model/jgopostgresqlmod"

type ReqQueryLog struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type ReqTrigger struct {
	Id string `json:"id" form:"id" validate:"required"`
}

type ResJobLog struct {
	Rows []jgopostgresqlmod.JobLog
}
