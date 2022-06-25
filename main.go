// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"github.com/xjustloveux/jgo.web/global"
	"github.com/xjustloveux/jgo.web/job"
	"github.com/xjustloveux/jgo.web/web"
	"github.com/xjustloveux/jgo/jcron"
	"github.com/xjustloveux/jgo/jlog"
	"github.com/xjustloveux/jgo/jsql"
)

func main() {
	if err := jgoInit(); err != nil {
		fmt.Println(err)
		return
	}
	jsql.SubscribeSql(sqlLog)
	job.CreateJob()
	if asc, err := global.Conf.Bool("autoStartCron"); err != nil {
		jlog.Error(err)
	} else if asc {
		defer func() {
			jcron.Wait()
			jcron.WaitTrigger()
		}()
		jcron.Start()
	}
	if err := routerInit(); err != nil {
		jlog.Error(err)
	}
}

func jgoInit() error {
	global.Conf.SetFileName("./config/config.json")
	jsql.SetFileName("./config/sql.json")
	jlog.SetFileName("./config/log.json")
	jcron.SetFileName("./config/cron.json")
	if err := global.Conf.Load(); err != nil {
		return err
	}
	if env, err := global.Conf.String("jEnv"); err != nil {
		return err
	} else {
		jsql.SetEnvVal(env)
		jlog.SetEnvVal(env)
	}
	if err := jsql.Init(); err != nil {
		return err
	}
	if err := jlog.Init(); err != nil {
		return err
	}
	if err := jcron.Init(); err != nil {
		return err
	}
	return nil
}

func routerInit() error {
	if env, err := global.Conf.String("jEnv"); err != nil {
		return err
	} else {
		if env == "prd" {
			gin.SetMode(gin.ReleaseMode)
		}
	}
	router := gin.Default()

	web.CreateRouter(router)

	if port, err := global.Conf.Int("port"); err != nil {
		return err
	} else {
		return router.Run(fmt.Sprint(":", port))
	}
}
