// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jgopostgresqlsrv

import (
	"github.com/fatih/structs"
	"github.com/xjustloveux/jgo.web/model/jgopostgresqlmod"
	"github.com/xjustloveux/jgo/jsql"
)

// PkgContent package web view content
var PkgContent = &pkgContent{
	ds:    "JGoPostgreSql",
	table: "pkg_content",
}

type pkgContent struct {
	ds    string
	table string
}

// Create insert data
func (srv *pkgContent) Create(data jgopostgresqlmod.PkgContent) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey: srv.ds,
		Table: srv.table,
		Col:   structs.Map(data),
	}).Insert()
}

// CreateTx insert data for tx
func (srv *pkgContent) CreateTx(agent *jsql.Agent, data jgopostgresqlmod.PkgContent) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent: agent,
		DSKey: srv.ds,
		Table: srv.table,
		Col:   structs.Map(data),
	}).Insert()
}

// FindAll query all data
func (srv *pkgContent) FindAll(param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).Query()
}

// FindAllTx query all data for tx
func (srv *pkgContent) FindAllTx(agent *jsql.Agent, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).QueryTx()
}

// FindFirst query first data
func (srv *pkgContent) FindFirst(param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).QueryRow()
}

// FindFirstTx query first data for tx
func (srv *pkgContent) FindFirstTx(agent *jsql.Agent, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).QueryRowTx()
}

// Update update data
func (srv *pkgContent) Update(data map[string]interface{}, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Col:    data,
		Params: param,
	}).Update()
}

// UpdateTx update data for tx
func (srv *pkgContent) UpdateTx(agent *jsql.Agent, data map[string]interface{}, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Col:    data,
		Params: param,
	}).UpdateTx()
}

// Delete delete data
func (srv *pkgContent) Delete(param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).Delete()
}

// DeleteTx delete data for tx
func (srv *pkgContent) DeleteTx(agent *jsql.Agent, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).DeleteTx()
}
