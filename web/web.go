// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xjustloveux/jgo.web/global"
	"github.com/xjustloveux/jgo.web/model"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jlog"
	"github.com/xjustloveux/jgo/jsql"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	router         *gin.Engine
	controllerList = make([]controller, 0)
)

func init() {
	controllerList = []controller{new(shared), new(a101), new(b101), new(b102), new(b103), new(c101)}
}

func CreateRouter(r *gin.Engine) {

	router = r

	for _, c := range controllerList {
		c.initRouter()
	}
	if env, err := global.Conf.String("jEnv"); err != nil {
		jlog.Error(err)
	} else if env != "dev" {
		var dist string
		if dist, err = global.Conf.String("dist"); err != nil {
			jlog.Error(err)
		} else {
			var list []string
			if list, err = getAngularFile(dist); err != nil {
				jlog.Error(err)
			} else {
				for _, file := range list {
					router.StaticFile(file, fmt.Sprint(dist, file))
				}
			}
			router.Static("/assets", fmt.Sprint(dist, "/assets"))
		}
		router.NoRoute(func(c *gin.Context) {
			c.File(fmt.Sprint(dist, "/index.html"))
		})
	}
}

func getAngularFile(path string) (names []string, err error) {
	var file *os.File
	if file, err = os.Open(path); err != nil {
		return
	}
	defer func() {
		if e := file.Close(); e != nil {
			err = e
		}
	}()
	names = make([]string, 0)
	var list []os.FileInfo
	if list, err = file.Readdir(0); err != nil {
		return
	} else {
		for _, f := range list {
			if !f.IsDir() {
				names = append(names, fmt.Sprint("./", filepath.Base(f.Name())))
			}
		}
	}
	return names, nil
}

func getProUrl(id string) (string, error) {
	ta := jsql.TableAgent{Table: "SYS_PRO", SelStr: "PRO_URL"}
	ta.Equal("PRO_ID", id)
	if res, err := ta.QueryRow(); err != nil {
		return "", err
	} else {
		r := res.Row()
		if r == nil {
			return "", fmt.Errorf("didn't query pro url")
		}
		url := jcast.String(r["PRO_URL"])
		if url == "" {
			return "", fmt.Errorf("didn't query pro url")
		}
		return url, nil
	}
}

func queryContent(ctx *gin.Context, proId string) {
	jlog.Info(proId, ":", "query content.")
	type list struct {
		Rows []model.PkgContent
	}
	if agent, err := jsql.GetAgent(); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"success": false})
		jlog.Error(err)
		return
	} else {
		var v list
		if _, err = agent.Query("B.QueryContent", map[string]interface{}{"PRO_ID": proId}, &v); err != nil {
			ctx.JSON(http.StatusOK, gin.H{"success": false})
			jlog.Error(err)
			return
		}
		for i, _ := range v.Rows {
			v.Rows[i].Code = strings.ReplaceAll(v.Rows[i].Code, "\r\n", "<br>")
			v.Rows[i].Code = strings.ReplaceAll(v.Rows[i].Code, "\r", "<br>")
			v.Rows[i].Code = strings.ReplaceAll(v.Rows[i].Code, "\n", "<br>")
			v.Rows[i].Code = strings.ReplaceAll(v.Rows[i].Code, "\t", "&emsp;")
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "content": v.Rows})
	}
}
