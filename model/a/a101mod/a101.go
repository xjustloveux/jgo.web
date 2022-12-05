// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package a101mod

import "time"

type ReqQueryPage struct {
	DbType string `json:"dbType" form:"dbType" validate:"required"`
	Page   int64  `json:"page" form:"page" validate:"required"`
	Size   int64  `json:"size" form:"size" validate:"required"`
}

type ReqCreateMsg struct {
	DbType  string `json:"dbType" form:"dbType" validate:"required"`
	Content string `json:"content" form:"content" validate:"required"`
}

type Msg struct {
	Content string    `json:"content"`
	CtDate  time.Time `json:"ct_date"`
}

type ResMsg struct {
	Rows        []Msg `json:"rows"`
	RowStart    int   `json:"rowStart"`
	RowEnd      int   `json:"rowEnd"`
	TotalRecord int   `json:"totalRecord"`
}
