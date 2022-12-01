// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package globalmod

type PkgContent struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Note        string `json:"note"`
}

type ResPkgContent struct {
	Rows []PkgContent
}
