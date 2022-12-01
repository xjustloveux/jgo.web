// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package c

import (
	"github.com/xjustloveux/jgo.web/controller/c/c101ctr"
	"github.com/xjustloveux/jgo.web/global"
)

func Init() {

	c := global.Router.Group("c")
	{
		c101 := c.Group("c101")
		{
			c101.GET("QueryVer", c101ctr.QueryVer)
			c101.GET("QueryVerUpdate", c101ctr.QueryVerUpdate)
		}
	}
}
