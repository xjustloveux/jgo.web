// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xjustloveux/jgo/jlog"
	"reflect"
	"strings"
)

type b101 struct {
	proId  string
	proUrl string
}

func (c *b101) initRouter() {
	c.proId = strings.ToUpper(reflect.TypeOf(c).Elem().Name())
	var err error
	if c.proUrl, err = getProUrl(c.proId); err != nil {
		jlog.Error(c.proId, ": ", err)
		return
	}
	router.GET(fmt.Sprint("/", c.proId, "/QueryContent"), c.queryContent)
}

func (c *b101) queryContent(ctx *gin.Context) {
	queryContent(ctx, c.proId)
}
