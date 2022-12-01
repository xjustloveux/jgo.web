// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jgopostgresqlmod

// SysPro system program
type SysPro struct {
	// ProId program id
	ProId string `json:"pro_id" structs:"pro_id"`
	// ProUrl program url
	ProUrl string `json:"pro_url" structs:"pro_url"`
	// ProName program name
	ProName string `json:"pro_name" structs:"pro_name"`
	// ParentProId parent program id
	ParentProId string `json:"parent_pro_id" structs:"parent_pro_id"`
	// Sort sort
	Sort int `json:"sort" structs:"sort"`
}
