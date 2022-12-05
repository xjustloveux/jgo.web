// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package job

import (
	"github.com/xjustloveux/jgo.web/model/jgopostgresqlmod"
	"github.com/xjustloveux/jgo.web/service/jgopostgresqlsrv"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jlog"
	"time"
)

type job002 struct{}

func (j *job002) Name() string {

	return "job002"
}

func (j *job002) Run(m map[string]interface{}) {

	jlog.Info(j.Name(), " start")
	now := time.Now()
	if loc, err := time.LoadLocation("Asia/Taipei"); err != nil {

		jlog.Error(err)
	} else {

		now = time.Now().In(loc)
	}
	if _, err := jgopostgresqlsrv.JobLog.Create(jgopostgresqlmod.JobLog{
		Name:   j.Name(),
		Id:     jcast.String(m["id"]),
		CtDate: now,
	}); err != nil {

		jlog.Error(err)
	}
	jlog.Info(j.Name(), " end")
}
