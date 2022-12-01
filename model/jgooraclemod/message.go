// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jgooraclemod

import "time"

// Message message, comment or note
type Message struct {
	// Ip ip address
	Ip string `json:"IP" structs:"IP"`
	// Content you want to say
	Content string `json:"CONTENT" structs:"CONTENT"`
	// CtDate create datetime
	CtDate time.Time `json:"CT_DATE" structs:"CT_DATE"`
}
