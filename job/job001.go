// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package job

import (
	"fmt"
	"github.com/xjustloveux/jgo.web/model/jgopostgresqlmod"
	"github.com/xjustloveux/jgo.web/model/job/job001mod"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jlog"
	"github.com/xjustloveux/jgo/jsql"
	"time"
)

type job001 struct{}

func (j *job001) Name() string {

	return "job001"
}

func (j *job001) Run(map[string]interface{}) {

	jlog.Info(j.Name(), " start")
	now := time.Now()
	if loc, err := time.LoadLocation("Asia/Taipei"); err != nil {

		jlog.Error(err)
	} else {

		now = time.Now().In(loc)
	}
	var delParam job001mod.DelParam
	if hrStr, err := jcast.TimeString(now); err != nil {

		jlog.Error(err)
		return
	} else if delParam.Time, err = jcast.TimeLoc(fmt.Sprint(hrStr[:14], "00:00"), time.UTC); err != nil {

		jlog.Error(err)
		return
	}
	if agent, err := jsql.GetAgent(); err != nil {

		jlog.Error(err)
	} else {

		if err = agent.UseTx(func() error {

			if _, e := agent.DeleteTx("Job001.DeleteLog", delParam); e != nil {

				return e
			}
			if _, e := agent.InsertTx("Job001.InsertLog", jgopostgresqlmod.JobLog{
				Name:   j.Name(),
				Id:     "SYSTEM",
				CtDate: now,
			}); e != nil {

				return e
			}
			return nil
		}); err != nil {

			jlog.Error(err)
		}
	}
	jlog.Info(j.Name(), " end")
}
