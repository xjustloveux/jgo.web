// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package menuctr

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xjustloveux/jgo.web/global/enum"
	"github.com/xjustloveux/jgo.web/middleware/ginc"
	"github.com/xjustloveux/jgo.web/service/sharedsrv/menusrv"
	"github.com/xjustloveux/jgo/jlog"
	"github.com/xjustloveux/jgo/jruntime"
)

func QueryMenu(ctx *gin.Context) {

	jlog.Info(fmt.Sprint(jruntime.GetFuncName(), "-", ctx.ClientIP()))
	if data, err := menusrv.QueryMenu(); err != nil {

		jlog.Error(err)
		ginc.Fail(ctx, enum.QueryFail)
		return
	} else {

		ginc.OkWithData(ctx, data)
	}
}
