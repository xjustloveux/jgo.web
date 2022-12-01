// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package a

import (
	"github.com/xjustloveux/jgo.web/controller/a/a101ctr"
	"github.com/xjustloveux/jgo.web/global"
)

func Init() {

	a := global.Router.Group("a")
	{
		a101 := a.Group("a101")
		{
			a101.GET("", a101ctr.QueryPage)
			a101.POST("CreateMsg", a101ctr.CreateMsg)
		}
	}
}
