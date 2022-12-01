// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package c101srv

import (
	"github.com/xjustloveux/jgo.web/model/c/c101mod"
	"github.com/xjustloveux/jgo.web/model/jgopostgresqlmod"
	"github.com/xjustloveux/jgo/jsql"
)

func QueryVer(req c101mod.ReqQueryVer) ([]jgopostgresqlmod.VerHs, error) {

	if agent, err := jsql.GetAgent(); err != nil {

		return nil, err
	} else {
		var v c101mod.ResVerHs
		if _, err = agent.Query("C101.QueryVer", req, &v); err != nil {

			return nil, err
		}
		return v.Rows, nil
	}
}

func QueryVerUpdate(req c101mod.ReqQueryVerUpdate) ([]jgopostgresqlmod.VerUpdate, error) {

	if agent, err := jsql.GetAgent(); err != nil {

		return nil, err
	} else {

		var v c101mod.ResVerUpdate
		if _, err = agent.Query("C101.QueryVerUpdate", req, &v); err != nil {

			return nil, err
		}
		return v.Rows, nil
	}
}
