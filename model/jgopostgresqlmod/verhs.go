// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jgopostgresqlmod

// VerHs version history
type VerHs struct {
	// Ver version
	Ver int `json:"ver" structs:"ver"`
	// Name version name
	Name string `json:"name" structs:"name"`
}
