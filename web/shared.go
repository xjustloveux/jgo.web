// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package web

import (
	"github.com/gin-gonic/gin"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jlog"
	"github.com/xjustloveux/jgo/jslice"
	"github.com/xjustloveux/jgo/jsql"
	"net/http"
)

type shared struct {
}

func (c *shared) initRouter() {
	router.GET("/Shared/QueryMenu", c.queryMenu)
}

func (c *shared) queryMenu(ctx *gin.Context) {
	ta := jsql.TableAgent{Table: "SYS_PRO", SelStr: "PRO_ID, PRO_URL, PRO_NAME, PARENT_PRO_ID", OrdStr: "SORT"}
	if res, err := ta.Query(); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"success": false})
		jlog.Error(err)
		return
	} else {
		var list []map[string]interface{}
		var l1 []interface{}
		if l1, err = jslice.Filter(res.Rows(), func(i interface{}) bool {
			return jcast.String(i.(map[string]interface{})["PARENT_PRO_ID"]) == ""
		}); err != nil {
			ctx.JSON(http.StatusOK, gin.H{"success": false})
			jlog.Error(err)
			return
		}
		for _, tl1 := range l1 {
			item1 := tl1.(map[string]interface{})
			var l2 []interface{}
			if l2, err = jslice.Filter(res.Rows(), func(i interface{}) bool {
				return jcast.String(i.(map[string]interface{})["PARENT_PRO_ID"]) == jcast.String(item1["PRO_ID"])
			}); err != nil {
				ctx.JSON(http.StatusOK, gin.H{"success": false})
				jlog.Error(err)
				return
			}
			delete(item1, "PRO_ID")
			delete(item1, "PARENT_PRO_ID")
			list2 := make([]map[string]interface{}, 0)
			for _, tl2 := range l2 {
				item2 := tl2.(map[string]interface{})
				delete(item2, "PRO_ID")
				delete(item2, "PARENT_PRO_ID")
				list2 = append(list2, item2)
			}
			item1["LIST"] = list2
			list = append(list, item1)
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "menu": list})
	}
}
