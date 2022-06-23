// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package job

import (
	"github.com/xjustloveux/jgo/jlog"
	"github.com/xjustloveux/jgo/jsql"
	"reflect"
	"time"
)

type job002 struct {
	name string
}

func (j *job002) Init() {
	j.name = reflect.TypeOf(j).Elem().Name()
}

func (j *job002) Name() string {
	return j.name
}

func (j *job002) Run(m map[string]interface{}) {
	jlog.Info(j.name, " Start")
	agent := &jsql.TableAgent{Table: "JOB_LOG"}
	if err := agent.AddColumn("NAME", j.name, "ID", m["id"], "CT_DATE", time.Now()); err != nil {
		jlog.Error(err)
		return
	}
	if _, err := agent.Insert(); err != nil {
		jlog.Error(err)
		return
	}
	jlog.Info(j.name, " End")
}
