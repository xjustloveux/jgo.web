// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xjustloveux/jgo.web/model"
	"github.com/xjustloveux/jgo/jcron"
	"github.com/xjustloveux/jgo/jlog"
	"github.com/xjustloveux/jgo/jsql"
	"net/http"
	"reflect"
	"strings"
)

type b102 struct {
	proId  string
	proUrl string
}

func (c *b102) initRouter() {
	c.proId = strings.ToUpper(reflect.TypeOf(c).Elem().Name())
	var err error
	if c.proUrl, err = getProUrl(c.proId); err != nil {
		jlog.Error(c.proId, ": ", err)
		return
	}
	router.GET(fmt.Sprint("/", c.proId, "/QueryContent"), c.queryContent)
	router.POST(fmt.Sprint("/", c.proId, "/QueryLog"), c.queryLog)
	router.POST(fmt.Sprint("/", c.proId, "/Trigger"), c.trigger)
}

func (c *b102) queryContent(ctx *gin.Context) {
	queryContent(ctx, c.proId)
}

func (c *b102) queryLog(ctx *gin.Context) {
	data := make(map[string]interface{})
	if err := ctx.BindJSON(&data); err != nil {
		jlog.Error(err)
		ctx.JSON(http.StatusOK, gin.H{"success": false})
		return
	}
	type list struct {
		Rows []model.JobLog
	}
	if agent, err := jsql.GetAgent(); err != nil {
		jlog.Error(err)
		ctx.JSON(http.StatusOK, gin.H{"success": false})
		return
	} else {
		var v list
		if _, err = agent.Query("B102.QueryLog", data, &v); err != nil {
			jlog.Error(err)
			ctx.JSON(http.StatusOK, gin.H{"success": false})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "list": v.Rows})
	}
}

func (c *b102) trigger(ctx *gin.Context) {
	data := make(map[string]interface{})
	if err := ctx.BindJSON(&data); err != nil {
		jlog.Error(err)
		ctx.JSON(http.StatusOK, gin.H{"success": false})
		return
	}
	if sch, err := jcron.GetSchedule("Sch002"); err != nil {
		jlog.Error(err)
		ctx.JSON(http.StatusOK, gin.H{"success": false})
		return
	} else {
		if err = sch.Trigger(data); err != nil {
			jlog.Error(err)
			ctx.JSON(http.StatusOK, gin.H{"success": false})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"success": true})
}
