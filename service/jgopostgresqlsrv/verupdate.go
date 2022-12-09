// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jgopostgresqlsrv

import (
	"github.com/fatih/structs"
	"github.com/xjustloveux/jgo.web/model/jgopostgresqlmod"
	"github.com/xjustloveux/jgo/jsql"
)

// VerUpdate version update
var VerUpdate = &verUpdate{
	ds:    "JGoPostgreSql",
	table: "ver_update",
}

type verUpdate struct {
	ds    string
	table string
}

// Ds return ds
func (srv *verUpdate) Ds() string {

	return srv.ds
}

// Table return table name
func (srv *verUpdate) Table() string {

	return srv.table
}

// Create insert data
func (srv *verUpdate) Create(data jgopostgresqlmod.VerUpdate) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey: srv.ds,
		Table: srv.table,
		Col:   structs.Map(data),
	}).Insert()
}

// CreateTx insert data for tx
func (srv *verUpdate) CreateTx(agent *jsql.Agent, data jgopostgresqlmod.VerUpdate) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent: agent,
		DSKey: srv.ds,
		Table: srv.table,
		Col:   structs.Map(data),
	}).Insert()
}

// FindAll query all data
func (srv *verUpdate) FindAll(param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).Query()
}

// FindAllTx query all data for tx
func (srv *verUpdate) FindAllTx(agent *jsql.Agent, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).QueryTx()
}

// FindFirst query first data
func (srv *verUpdate) FindFirst(param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).QueryRow()
}

// FindFirstTx query first data for tx
func (srv *verUpdate) FindFirstTx(agent *jsql.Agent, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).QueryRowTx()
}

// Update update data
func (srv *verUpdate) Update(data map[string]interface{}, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Col:    data,
		Params: param,
	}).Update()
}

// UpdateTx update data for tx
func (srv *verUpdate) UpdateTx(agent *jsql.Agent, data map[string]interface{}, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Col:    data,
		Params: param,
	}).UpdateTx()
}

// Delete delete data
func (srv *verUpdate) Delete(param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).Delete()
}

// DeleteTx delete data for tx
func (srv *verUpdate) DeleteTx(agent *jsql.Agent, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).DeleteTx()
}
