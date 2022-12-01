// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package b103mod

type ResLog struct {
	Name  string   `json:"name"`
	Dir   bool     `json:"dir"`
	Child []ResLog `json:"child"`
}
