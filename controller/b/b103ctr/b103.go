// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package b103ctr

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xjustloveux/jgo.web/enum/message"
	"github.com/xjustloveux/jgo.web/middleware/ginc"
	"github.com/xjustloveux/jgo.web/service/b/b103srv"
	"github.com/xjustloveux/jgo/jlog"
	"github.com/xjustloveux/jgo/jruntime"
)

func QueryContent(ctx *gin.Context) {

	jlog.Info(fmt.Sprint(jruntime.GetFuncName(), "-", ctx.RemoteIP()))
	if data, err := b103srv.QueryContent(); err != nil {

		jlog.Error(err)
		ginc.Fail(ctx, message.QueryFail)
		return
	} else {

		ginc.OkWithData(ctx, data)
	}
}

func QueryLog(ctx *gin.Context) {

	jlog.Info(fmt.Sprint(jruntime.GetFuncName(), "-", ctx.RemoteIP()))
	if data, err := b103srv.QueryLog(); err != nil {

		jlog.Error(err)
		ginc.Fail(ctx, message.QueryFail)
		return
	} else {

		ginc.OkWithData(ctx, data)
	}
}
