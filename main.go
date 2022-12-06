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
	"github.com/xjustloveux/jgo.web/global"
	"github.com/xjustloveux/jgo.web/job"
	"github.com/xjustloveux/jgo.web/middleware/log"
	"github.com/xjustloveux/jgo.web/middleware/yaml"
	"github.com/xjustloveux/jgo/jconf"
	"github.com/xjustloveux/jgo/jcron"
	"github.com/xjustloveux/jgo/jfile"
	"github.com/xjustloveux/jgo/jlog"
	"github.com/xjustloveux/jgo/jsql"
)

func main() {

	if err := jGoInit(); err != nil {

		fmt.Println(err)
		return
	}
	if global.Conf.AutoStartCron {

		defer func() {

			jcron.Wait()
			jcron.WaitTrigger()
		}()
		jcron.Start()
	}
	if err := controller.Init(); err != nil {

		jlog.Error(err)
	}
}

func jGoInit() error {

	jfile.RegisterCodec(jfile.Yaml.String(), yaml.Codec{})
	configFile := "config.yaml"
	conf := jconf.New()
	conf.SetFormat(jfile.Yaml)
	conf.SetFileName(configFile)
	jsql.SetFormat(jfile.Yaml)
	jsql.SetFileName(configFile)
	jcron.SetFormat(jfile.Yaml)
	jcron.SetFileName(configFile)
	jlog.SetFormat(jfile.Yaml)
	jlog.SetFileName(configFile)
	if err := conf.Load(); err != nil {

		return err
	}
	if err := conf.Convert(&global.Conf); err != nil {

		return err
	}
	if err := jsql.Init(); err != nil {

		return err
	}
	if err := jcron.Init(); err != nil {

		return err
	}
	if err := jlog.Init(); err != nil {

		return err
	}
	if err := job.Init(); err != nil {

		return err
	}
	jsql.SubscribeSql(log.SqlLog)
	return nil
}
