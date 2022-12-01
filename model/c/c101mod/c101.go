// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package c101mod

import "github.com/xjustloveux/jgo.web/model/jgopostgresqlmod"

type ReqQueryVer struct {
	Ver *int `json:"ver" form:"ver"`
}

type ReqQueryVerUpdate struct {
	VerS int `json:"verS" form:"verS" validate:"required"`
	VerE int `json:"verE" form:"verE" validate:"required"`
}

type ResVerHs struct {
	Rows []jgopostgresqlmod.VerHs
}

type ResVerUpdate struct {
	Rows []jgopostgresqlmod.VerUpdate
}
