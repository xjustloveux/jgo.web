// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package job

import (
	"fmt"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jlog"
	"github.com/xjustloveux/jgo/jsql"
	"reflect"
	"time"
)

type job001 struct {
	name string
}

func (j *job001) Init() {
	j.name = reflect.TypeOf(j).Elem().Name()
}

func (j *job001) Name() string {
	return j.name
}

func (j *job001) Run(map[string]interface{}) {
	jlog.Info(j.name, " Start")
	now := time.Now()
	var hrTime time.Time
	if hrStr, err := jcast.TimeString(now); err != nil {
		jlog.Error(err)
		return
	} else if hrTime, err = jcast.TimeLoc(fmt.Sprint(hrStr[:14], "00:00"), time.UTC); err != nil {
		jlog.Error(err)
		return
	}
	if agent, err := jsql.GetAgent(); err != nil {
		jlog.Error(err)
	} else {
		if _, err = agent.Begin(); err != nil {
			jlog.Error(err)
			return
		}
		defer func() {
			if err != nil {
				if e := agent.Rollback(); e != nil {
					jlog.Error(e)
				}
			}
		}()
		if _, err = agent.DeleteTx("Job001.DeleteLog", map[string]interface{}{"TIME": hrTime}); err != nil {
			jlog.Error(err)
			return
		}
		if _, err = agent.InsertTx("Job001.InsertLog", map[string]interface{}{"NAME": j.name, "ID": "SYSTEM", "CT_DATE": now}); err != nil {
			jlog.Error(err)
			return
		}
		if err = agent.Commit(); err != nil {
			jlog.Error(err)
			return
		}
	}
	jlog.Info(j.name, " End")
}
