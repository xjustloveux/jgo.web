// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jgopostgresqlsrv

import (
	"github.com/fatih/structs"
	"github.com/xjustloveux/jgo.web/model/jgopostgresqlmod"
	"github.com/xjustloveux/jgo/jsql"
)

// SysPro system program
var SysPro = &sysPro{
	ds:    "JGoPostgreSql",
	table: "sys_pro",
}

type sysPro struct {
	ds    string
	table string
}

// Ds return ds
func (srv *sysPro) Ds() string {

	return srv.ds
}

// Table return table name
func (srv *sysPro) Table() string {

	return srv.table
}

// Create insert data
func (srv *sysPro) Create(data jgopostgresqlmod.SysPro) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey: srv.ds,
		Table: srv.table,
		Col:   structs.Map(data),
	}).Insert()
}

// CreateTx insert data for tx
func (srv *sysPro) CreateTx(agent *jsql.Agent, data jgopostgresqlmod.SysPro) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent: agent,
		DSKey: srv.ds,
		Table: srv.table,
		Col:   structs.Map(data),
	}).Insert()
}

// FindAll query all data
func (srv *sysPro) FindAll(param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).Query()
}

// FindAllTx query all data for tx
func (srv *sysPro) FindAllTx(agent *jsql.Agent, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).QueryTx()
}

// FindFirst query first data
func (srv *sysPro) FindFirst(param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).QueryRow()
}

// FindFirstTx query first data for tx
func (srv *sysPro) FindFirstTx(agent *jsql.Agent, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).QueryRowTx()
}

// Update update data
func (srv *sysPro) Update(data map[string]interface{}, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Col:    data,
		Params: param,
	}).Update()
}

// UpdateTx update data for tx
func (srv *sysPro) UpdateTx(agent *jsql.Agent, data map[string]interface{}, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Col:    data,
		Params: param,
	}).UpdateTx()
}

// Delete delete data
func (srv *sysPro) Delete(param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).Delete()
}

// DeleteTx delete data for tx
func (srv *sysPro) DeleteTx(agent *jsql.Agent, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).DeleteTx()
}
