// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package shared

import (
	"github.com/xjustloveux/jgo.web/controller/shared/menuctr"
	"github.com/xjustloveux/jgo.web/global"
)

func Init() {

	shared := global.Router.Group("shared")
	{
		menu := shared.Group("menu")
		{
			menu.GET("", menuctr.QueryMenu)
		}
	}
}
