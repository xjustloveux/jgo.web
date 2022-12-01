// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package menusrv

import (
	"github.com/xjustloveux/jgo.web/model/jgopostgresqlmod"
	"github.com/xjustloveux/jgo.web/model/menumod"
	"github.com/xjustloveux/jgo.web/service/jgopostgresqlsrv"
	"github.com/xjustloveux/jgo/jfile"
	"github.com/xjustloveux/jgo/jslice"
	"sort"
)

func QueryMenu() ([]menumod.ResMenu, error) {

	if res, err := jgopostgresqlsrv.SysPro.FindAll(); err != nil {

		return nil, err
	} else {

		var v menumod.ResSysPro

		if err = jfile.Convert(map[string]interface{}{"Rows": res.Rows()}, &v); err != nil {

			return nil, err
		}
		sort.Slice(v.Rows, func(i, j int) bool {

			return v.Rows[i].Sort < v.Rows[j].Sort
		})
		var l1 []interface{}
		if l1, err = jslice.Filter(v.Rows, func(i interface{}) bool {

			return i.(jgopostgresqlmod.SysPro).ParentProId == ""
		}); err != nil {

			return nil, err
		}
		menu := make([]menumod.ResMenu, len(l1))
		for i1, v1 := range l1 {

			val1 := v1.(jgopostgresqlmod.SysPro)
			var l2 []interface{}
			if l2, err = jslice.Filter(v.Rows, func(i interface{}) bool {

				return i.(jgopostgresqlmod.SysPro).ParentProId == val1.ProId
			}); err != nil {

				return nil, err
			}
			menu[i1].Name = val1.ProName
			menu[i1].Url = val1.ProUrl
			menu[i1].Child = make([]menumod.ResMenu, len(l2))
			for i2, v2 := range l2 {

				val2 := v2.(jgopostgresqlmod.SysPro)
				menu[i1].Child[i2].Name = val2.ProName
				menu[i1].Child[i2].Url = val2.ProUrl
			}
		}
		return menu, nil
	}
}
