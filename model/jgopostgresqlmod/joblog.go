// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jgopostgresqlmod

import "time"

// JobLog job log
type JobLog struct {
	// Name job name
	Name string `json:"name" structs:"name"`
	// Id id
	Id string `json:"id" structs:"id"`
	// CtDate create datetime
	CtDate time.Time `json:"ct_date" structs:"ct_date"`
}
