// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package c101ctr

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xjustloveux/jgo.web/global/enum"
	"github.com/xjustloveux/jgo.web/middleware/ginc"
	"github.com/xjustloveux/jgo.web/middleware/validatorc"
	"github.com/xjustloveux/jgo.web/model/c/c101mod"
	"github.com/xjustloveux/jgo.web/service/c/c101srv"
	"github.com/xjustloveux/jgo/jlog"
	"github.com/xjustloveux/jgo/jruntime"
)

func QueryVer(ctx *gin.Context) {

	jlog.Info(fmt.Sprint(jruntime.GetFuncName(), "-", ctx.ClientIP()))
	var req c101mod.ReqQueryVer
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
	if data, err := c101srv.QueryVer(req); err != nil {

		jlog.Error(err)
		ginc.Fail(ctx, enum.QueryFail)
		return
	} else {

		ginc.OkWithData(ctx, data)
	}
}

func QueryVerUpdate(ctx *gin.Context) {

	jlog.Info(fmt.Sprint(jruntime.GetFuncName(), "-", ctx.ClientIP()))
	var req c101mod.ReqQueryVerUpdate
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
	if data, err := c101srv.QueryVerUpdate(req); err != nil {

		jlog.Error(err)
		ginc.Fail(ctx, enum.QueryFail)
		return
	} else {

		ginc.OkWithData(ctx, data)
	}
}
