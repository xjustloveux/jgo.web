// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package a101ctr

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xjustloveux/jgo.web/global/enum"
	"github.com/xjustloveux/jgo.web/middleware/ginc"
	"github.com/xjustloveux/jgo.web/middleware/validatorc"
	"github.com/xjustloveux/jgo.web/model/a/a101mod"
	"github.com/xjustloveux/jgo.web/service/a/a101srv"
	"github.com/xjustloveux/jgo/jlog"
	"github.com/xjustloveux/jgo/jruntime"
)

func QueryPage(ctx *gin.Context) {

	jlog.Info(fmt.Sprint(jruntime.GetFuncName(), "-", ctx.RemoteIP()))
	var req a101mod.ReqQueryPage
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
	if data, err := a101srv.QueryPage(req); err != nil {

		jlog.Error(err)
		ginc.Fail(ctx, enum.QueryFail)
		return
	} else {

		ginc.OkWithData(ctx, data)
	}
}

func CreateMsg(ctx *gin.Context) {

	jlog.Info(fmt.Sprint(jruntime.GetFuncName(), "-", ctx.RemoteIP()))
	var req a101mod.ReqCreateMsg
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
	if err := a101srv.CreateMsg(req, ctx.RemoteIP()); err != nil {

		jlog.Error(err)
		ginc.Fail(ctx, enum.CreateFail)
	} else {

		ginc.Ok(ctx)
	}
}
