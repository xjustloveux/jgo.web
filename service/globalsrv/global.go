// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package globalsrv

import (
	"github.com/xjustloveux/jgo.web/model/globalmod"
	"github.com/xjustloveux/jgo/jsql"
	"strings"
)

func QueryContent(proId string) ([]globalmod.PkgContent, error) {

	if agent, err := jsql.GetAgent(); err != nil {

		return nil, err
	} else {
		var v globalmod.ResPkgContent
		if _, err = agent.Query("B.QueryContent", map[string]interface{}{"PRO_ID": proId}, &v); err != nil {

			return nil, err
		}
		for i, _ := range v.Rows {
			v.Rows[i].Code = strings.ReplaceAll(v.Rows[i].Code, "\r\n", "<br>")
			v.Rows[i].Code = strings.ReplaceAll(v.Rows[i].Code, "\r", "<br>")
			v.Rows[i].Code = strings.ReplaceAll(v.Rows[i].Code, "\n", "<br>")
			v.Rows[i].Code = strings.ReplaceAll(v.Rows[i].Code, "\t", "&emsp;")
		}
		return v.Rows, nil
	}
}
