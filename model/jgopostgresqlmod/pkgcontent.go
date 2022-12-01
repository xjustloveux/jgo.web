// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jgopostgresqlmod

// PkgContent package web view content
type PkgContent struct {
	// Seq sequence
	Seq int `json:"seq" structs:"seq"`
	// ProId program id
	ProId string `json:"pro_id" structs:"pro_id"`
	// Title title
	Title string `json:"title" structs:"title"`
	// Description description
	Description string `json:"description" structs:"description"`
	// Code code
	Code string `json:"code" structs:"code"`
	// Note note
	Note string `json:"note" structs:"note"`
	// Sort sort
	Sort int `json:"sort" structs:"sort"`
}
