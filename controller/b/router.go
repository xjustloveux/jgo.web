// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package b

import (
	"github.com/xjustloveux/jgo.web/controller/b/b101ctr"
	"github.com/xjustloveux/jgo.web/controller/b/b102ctr"
	"github.com/xjustloveux/jgo.web/controller/b/b103ctr"
	"github.com/xjustloveux/jgo.web/global"
)

func Init() {

	b := global.Router.Group("b")
	{
		b101 := b.Group("b101")
		{
			b101.GET("", b101ctr.QueryContent)
		}
		b102 := b.Group("b102")
		{
			b102.GET("", b102ctr.QueryContent)
			b102.GET("QueryLog", b102ctr.QueryLog)
			b102.POST("Trigger", b102ctr.Trigger)
		}
		b103 := b.Group("b103")
		{
			b103.GET("", b103ctr.QueryContent)
			b103.GET("QueryLog", b103ctr.QueryLog)
		}
	}
}
