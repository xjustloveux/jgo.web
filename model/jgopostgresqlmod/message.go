// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jgopostgresqlmod

import "time"

// Message message, comment or note
type Message struct {
	// Ip ip address
	Ip string `json:"ip" structs:"ip"`
	// Content you want to say
	Content string `json:"content" structs:"content"`
	// CtDate create datetime
	CtDate time.Time `json:"ct_date" structs:"ct_date"`
}
