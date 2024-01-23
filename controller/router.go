// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/xjustloveux/jgo.web/controller/a"
	"github.com/xjustloveux/jgo.web/controller/b"
	"github.com/xjustloveux/jgo.web/controller/c"
	"github.com/xjustloveux/jgo.web/controller/shared"
	"github.com/xjustloveux/jgo.web/global"
	"github.com/xjustloveux/jgo/jlog"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Init() error {

	global.Router.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     global.Conf.CORS,
				AllowMethods:     []string{"POST", "GET"},
				AllowHeaders:     []string{"Content-Type", "Content-Length"},
				AllowCredentials: true,
				ExposeHeaders:    []string{"Access-Control-Allow-Headers"},
			},
		),
	)
	global.Router.TrustedPlatform = "CF-Connecting-IP"
	a.Init()
	b.Init()
	c.Init()
	shared.Init()
	if err := initAngular(); err != nil {

		return err
	}
	return global.Router.Run(fmt.Sprint(":", global.Conf.Port))
}

func initAngular() error {

	global.Router.Static("/assets", global.Conf.Assets)
	global.Router.StaticFile("./sitemap.xml", "./sitemap.xml")
	if list, err := getAngularFile(); err != nil {

		return err
	} else {

		for _, file := range list {

			global.Router.StaticFile(file, fmt.Sprint("./dist/browser/", file))
		}
	}
	global.Router.NoRoute(func(c *gin.Context) {

		if content, err := getAngularSSR(c.Request.URL.Path); err != nil {

			c.String(http.StatusInternalServerError, err.Error())
		} else {

			c.Data(http.StatusOK, "text/html; charset=utf-8", content)
		}
	})
	return nil
}

func getAngularFile() ([]string, error) {

	if file, err := os.Open("./dist/browser"); err != nil {

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

func getAngularSSR(path string) ([]byte, error) {

	if resp, err := http.Get(fmt.Sprint(global.Conf.UrlSSR, path)); err != nil {

		jlog.Error(err)
		return nil, err
	} else {

		defer resp.Body.Close()
		return io.ReadAll(resp.Body)
	}
}
