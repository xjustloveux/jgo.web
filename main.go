// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
	_ "github.com/lib/pq"
	"github.com/xjustloveux/jgo.web/controller"
	"github.com/xjustloveux/jgo.web/middleware/jgoc"
	"github.com/xjustloveux/jgo/jlog"
)

func main() {

	if fc, err := jgoc.Init(); err != nil {

		fmt.Println(err)
		return
	} else {

		defer fc()
	}
	if err := controller.Init(); err != nil {

		jlog.Error(err)
	}
}
