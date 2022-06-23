// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xjustloveux/jgo.web/model"
	"github.com/xjustloveux/jgo/jlog"
	"github.com/xjustloveux/jgo/jsql"
	"net/http"
	"reflect"
	"strings"
)

type c101 struct {
	proId  string
	proUrl string
}

func (c *c101) initRouter() {
	c.proId = strings.ToUpper(reflect.TypeOf(c).Elem().Name())
	var err error
	if c.proUrl, err = getProUrl(c.proId); err != nil {
		jlog.Error(c.proId, ": ", err)
		return
	}
	router.POST(fmt.Sprint("/", c.proId, "/QueryVerList"), c.queryVerList)
	router.POST(fmt.Sprint("/", c.proId, "/QueryVerUpdate"), c.queryVerUpdate)
}

func (c *c101) queryVerList(ctx *gin.Context) {
	data := make(map[string]interface{})
	if err := ctx.BindJSON(&data); err != nil {
		jlog.Error(err)
		ctx.JSON(http.StatusOK, gin.H{"success": false})
		return
	}
	if v := c.queryVer(data); v == nil {
		ctx.JSON(http.StatusOK, gin.H{"success": false})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"success": true, "list": v})
	}
}

func (c *c101) queryVer(ver interface{}) []model.VerHs {
	type list struct {
		Rows []model.VerHs
	}
	if agent, err := jsql.GetAgent(); err != nil {
		jlog.Error(err)
		return nil
	} else {
		var v list
		if _, err = agent.Query("C101.QueryVerList", ver, &v); err != nil {
			jlog.Error(err)
			return nil
		}
		return v.Rows
	}
}

func (c *c101) queryVerUpdate(ctx *gin.Context) {
	data := make(map[string]interface{})
	if err := ctx.BindJSON(&data); err != nil {
		jlog.Error(err)
		ctx.JSON(http.StatusOK, gin.H{"success": false})
		return
	}
	if v := c.queryUpdate(data); v == nil {
		ctx.JSON(http.StatusOK, gin.H{"success": false})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"success": true, "list": v})
	}
}

func (c *c101) queryUpdate(ver interface{}) []model.VerUpdate {
	type list struct {
		Rows []model.VerUpdate
	}
	if agent, err := jsql.GetAgent(); err != nil {
		jlog.Error(err)
		return nil
	} else {
		var v list
		if _, err = agent.Query("C101.QueryVerUpdate", ver, &v); err != nil {
			jlog.Error(err)
			return nil
		}
		return v.Rows
	}
}
