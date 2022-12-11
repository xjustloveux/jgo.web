// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package b102ctr

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xjustloveux/jgo.web/global/enum"
	"github.com/xjustloveux/jgo.web/middleware/ginc"
	"github.com/xjustloveux/jgo.web/middleware/validatorc"
	"github.com/xjustloveux/jgo.web/model/b/b102mod"
	"github.com/xjustloveux/jgo.web/service/b/b102srv"
	"github.com/xjustloveux/jgo/jlog"
	"github.com/xjustloveux/jgo/jruntime"
)

func QueryContent(ctx *gin.Context) {

	jlog.Info(fmt.Sprint(jruntime.GetFuncName(), "-", ctx.ClientIP()))
	if data, err := b102srv.QueryContent(); err != nil {

		jlog.Error(err)
		ginc.Fail(ctx, enum.QueryFail)
		return
	} else {

		ginc.OkWithData(ctx, data)
	}
}

func QueryLog(ctx *gin.Context) {

	jlog.Info(fmt.Sprint(jruntime.GetFuncName(), "-", ctx.ClientIP()))
	var req b102mod.ReqQueryLog
	if err := ctx.ShouldBindQuery(&req); err != nil {

		jlog.Error(err)
		ginc.Fail(ctx, enum.InvalidRequest)
		return
	}
	if err := validatorc.Verify(req); err != nil {

		jlog.Error(err)
		ginc.Fail(ctx, enum.InvalidRequest)
		return
	}
	if data, err := b102srv.QueryLog(req); err != nil {

		jlog.Error(err)
		ginc.Fail(ctx, enum.QueryFail)
		return
	} else {

		ginc.OkWithData(ctx, data)
	}
}

func Trigger(ctx *gin.Context) {

	jlog.Info(fmt.Sprint(jruntime.GetFuncName(), "-", ctx.ClientIP()))
	var req b102mod.ReqTrigger
	if err := ctx.ShouldBindJSON(&req); err != nil {

		jlog.Error(err)
		ginc.Fail(ctx, enum.InvalidRequest)
		return
	}
	if err := validatorc.Verify(req); err != nil {

		jlog.Error(err)
		ginc.Fail(ctx, enum.InvalidRequest)
		return
	}
	if err := b102srv.Trigger(req); err != nil {

		jlog.Error(err)
		ginc.Fail(ctx, enum.TriggerFail)
	} else {

		ginc.Ok(ctx)
	}
}
