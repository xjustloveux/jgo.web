// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xjustloveux/jgo.web/controller/a"
	"github.com/xjustloveux/jgo.web/controller/b"
	"github.com/xjustloveux/jgo.web/controller/c"
	"github.com/xjustloveux/jgo.web/controller/shared"
	"github.com/xjustloveux/jgo.web/global"
	"os"
	"path/filepath"
)

func Init() error {

	a.Init()
	b.Init()
	c.Init()
	shared.Init()
	if err := initAngular(); err != nil {

		return err
	}
	if port, err := global.Conf.String("port"); err != nil {

		return err
	} else if len(port) > 0 {

		return global.Router.Run(fmt.Sprint(":", port))
	}
	return global.Router.Run(":8080")
}

func initAngular() error {

	if assets, err := global.Conf.String("assets"); err != nil {

		return err
	} else {

		global.Router.Static("/assets", assets)
		global.Router.StaticFile("./sitemap.xml", "./sitemap.xml")
		if !global.Dev {

			var list []string
			if list, err = getAngularFile(); err != nil {

				return err
			} else {

				for _, file := range list {

					global.Router.StaticFile(file, fmt.Sprint("./dist/", file))
				}
			}
			global.Router.NoRoute(func(c *gin.Context) {

				c.File("./dist/index.html")
			})
		}
	}
	return nil
}

func getAngularFile() ([]string, error) {

	if file, err := os.Open("./dist"); err != nil {

		return nil, err
	} else {

		defer func() {

			if e := file.Close(); e != nil {
				err = e
			}
		}()
		names := make([]string, 0)
		var list []os.FileInfo
		if list, err = file.Readdir(0); err != nil {

			return nil, err
		} else {

			for _, f := range list {

				if !f.IsDir() {

					names = append(names, filepath.Base(f.Name()))
				}
			}
		}
		return names, nil
	}
}
