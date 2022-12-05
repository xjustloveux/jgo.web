// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package a101srv

import (
	"github.com/xjustloveux/jgo.web/model/a/a101mod"
	"github.com/xjustloveux/jgo.web/model/jgomssqlmod"
	"github.com/xjustloveux/jgo.web/model/jgomysqlmod"
	"github.com/xjustloveux/jgo.web/model/jgooraclemod"
	"github.com/xjustloveux/jgo.web/model/jgopostgresqlmod"
	"github.com/xjustloveux/jgo.web/service/jgomssqlsrv"
	"github.com/xjustloveux/jgo.web/service/jgomysqlsrv"
	"github.com/xjustloveux/jgo.web/service/jgooraclesrv"
	"github.com/xjustloveux/jgo.web/service/jgopostgresqlsrv"
	"github.com/xjustloveux/jgo/jsql"
	"time"
)

func QueryPage(req a101mod.ReqQueryPage) (*a101mod.ResMsg, error) {

	if t, err := jsql.ParseDBType(req.DbType); err != nil {

		return nil, err
	} else {

		ds := ""
		switch t {
		case jsql.MySql:
			ds = "JGoMySql"
		case jsql.MSSql:
			ds = "JGoMSSql"
		case jsql.Oracle:
			ds = "JGoOracle"
		case jsql.PostgreSql:
			ds = "JGoPostgreSql"
		}
		var agent *jsql.Agent
		if agent, err = jsql.GetAgent(ds); err != nil {

			return nil, err
		}
		var v a101mod.ResMsg
		s := 1 + ((req.Page - 1) * req.Size)
		e := req.Page * req.Size
		if _, err = agent.QueryPage("A.QueryPage", s, e, &v); err != nil {

			return nil, err
		}
		return &v, nil
	}
}

func CreateMsg(req a101mod.ReqCreateMsg, ip string) error {

	if t, err := jsql.ParseDBType(req.DbType); err != nil {

		return err
	} else {

		var loc *time.Location
		if loc, err = time.LoadLocation("Asia/Taipei"); err != nil {

			return err
		}
		now := time.Now().In(loc)
		switch t {
		case jsql.MySql:
			if _, err = jgomysqlsrv.Message.Create(jgomysqlmod.Message{
				Ip:      ip,
				Content: req.Content,
				CtDate:  now,
			}); err != nil {

				return err
			}
		case jsql.MSSql:
			if _, err = jgomssqlsrv.Message.Create(jgomssqlmod.Message{
				Ip:      ip,
				Content: req.Content,
				CtDate:  now,
			}); err != nil {

				return err
			}
		case jsql.Oracle:
			if _, err = jgooraclesrv.Message.Create(jgooraclemod.Message{
				Ip:      ip,
				Content: req.Content,
				CtDate:  now,
			}); err != nil {

				return err
			}
		case jsql.PostgreSql:
			if _, err = jgopostgresqlsrv.Message.Create(jgopostgresqlmod.Message{
				Ip:      ip,
				Content: req.Content,
				CtDate:  now,
			}); err != nil {

				return err
			}
		}
	}
	return nil
}
