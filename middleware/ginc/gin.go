// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package ginc

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

func OkWithData(ctx *gin.Context, data interface{}) {

	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}

func Fail(ctx *gin.Context, code int) {

	ctx.JSON(http.StatusOK, gin.H{"success": false, "code": code})
}

func FailWithData(ctx *gin.Context, code int, data interface{}) {

	ctx.JSON(http.StatusOK, gin.H{"success": false, "code": code, "data": data})
}
