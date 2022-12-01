// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jgopostgresqlmod

// VerUpdate version update
type VerUpdate struct {
	// Seq sequence
	Seq int `json:"seq" structs:"seq"`
	// VerS version start
	VerS int `json:"ver_s" structs:"ver_s"`
	// VerE version end
	VerE int `json:"ver_e" structs:"ver_e"`
	// Content update content
	Content string `json:"content" structs:"content"`
}
