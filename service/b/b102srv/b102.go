// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package b102srv

import (
	"github.com/fatih/structs"
	"github.com/xjustloveux/jgo.web/model/b/b102mod"
	"github.com/xjustloveux/jgo.web/model/globalmod"
	"github.com/xjustloveux/jgo.web/model/jgopostgresqlmod"
	"github.com/xjustloveux/jgo.web/service/globalsrv"
	"github.com/xjustloveux/jgo/jcron"
	"github.com/xjustloveux/jgo/jsql"
)

func QueryContent() ([]globalmod.PkgContent, error) {

	return globalsrv.QueryContent("B102")
}

func QueryLog(req b102mod.ReqQueryLog) ([]jgopostgresqlmod.JobLog, error) {

	if agent, err := jsql.GetAgent(); err != nil {

		return nil, err
	} else {

		var v b102mod.ResJobLog
		if _, err = agent.Query("B102.QueryLog", req, &v); err != nil {

			return nil, err
		}
		return v.Rows, nil
	}
}

func Trigger(req b102mod.ReqTrigger) error {

	if sch, err := jcron.GetSchedule("Sch002"); err != nil {

		return err
	} else {

		if err = sch.Trigger(structs.Map(req)); err != nil {

			return err
		}
	}
	return nil
}
