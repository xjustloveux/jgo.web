// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package web

import (
	"github.com/xjustloveux/jgo/jlog"
	"reflect"
	"strings"
)

type a101 struct {
	proId  string
	proUrl string
}

func (c *a101) initRouter() {
	c.proId = strings.ToUpper(reflect.TypeOf(c).Elem().Name())
	var err error
	if c.proUrl, err = getProUrl(c.proId); err != nil {
		jlog.Error(c.proId, ": ", err)
		return
	}
}
