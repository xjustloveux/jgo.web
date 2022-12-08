// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jgoc

import (
	"github.com/xjustloveux/jgo.web/global"
	"github.com/xjustloveux/jgo.web/job"
	"github.com/xjustloveux/jgo.web/middleware/logc"
	"github.com/xjustloveux/jgo.web/middleware/yamlc"
	"github.com/xjustloveux/jgo/jconf"
	"github.com/xjustloveux/jgo/jcron"
	"github.com/xjustloveux/jgo/jfile"
	"github.com/xjustloveux/jgo/jlog"
	"github.com/xjustloveux/jgo/jsql"
)

func Init() (func(), error) {

	jfile.RegisterCodec(jfile.Yaml.String(), yamlc.Codec{})
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

		return nil, err
	}
	if err := conf.Convert(&global.Conf); err != nil {

		return nil, err
	}
	if err := jsql.Init(); err != nil {

		return nil, err
	}
	if err := jcron.Init(); err != nil {

		return nil, err
	}
	if err := jlog.Init(); err != nil {

		return nil, err
	}
	if err := job.Init(); err != nil {

		return nil, err
	}
	jsql.SubscribeSql(logc.SqlLog)
	var fc func()
	if global.Conf.AutoStartCron {

		fc = func() {

			jcron.Wait()
			jcron.WaitTrigger()
		}
		jcron.Start()
	}
	return fc, nil
}
