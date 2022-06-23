// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xjustloveux/jgo.web/model"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jlog"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type b103 struct {
	proId  string
	proUrl string
}

func (c *b103) initRouter() {
	c.proId = strings.ToUpper(reflect.TypeOf(c).Elem().Name())
	var err error
	if c.proUrl, err = getProUrl(c.proId); err != nil {
		jlog.Error(c.proId, ": ", err)
		return
	}
	router.GET(fmt.Sprint("/", c.proId, "/QueryContent"), c.queryContent)
	router.GET(fmt.Sprint("/", c.proId, "/QueryLog"), c.queryLog)
}

func (c *b103) queryContent(ctx *gin.Context) {
	queryContent(ctx, c.proId)
}

func (c *b103) queryLog(ctx *gin.Context) {
	path := jcast.String(jlog.GetParam("path"))
	if logFile, err := c.getFolderList(path); err != nil {
		jlog.Error(err)
		ctx.JSON(http.StatusOK, gin.H{"success": false})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"success": true, "list": logFile})
	}
}

func (c *b103) getFolderList(path string) (logFile model.LogFile, err error) {
	path = strings.Trim(strings.Trim(path, "\\ "), "/ ")
	logFile.Child = make([]model.LogFile, 0)
	var fileInfo os.FileInfo
	if fileInfo, err = os.Stat(path); err != nil {
		return
	}
	logFile.Dir = fileInfo.IsDir()
	if logFile.Dir {
		logFile.Name = filepath.Base(fileInfo.Name())
		var file *os.File
		if file, err = os.Open(path); err != nil {
			return
		}
		defer func() {
			if e := file.Close(); e != nil {
				err = e
			}
		}()
		var list []os.FileInfo
		if list, err = file.Readdir(0); err != nil {
			return
		} else {
			for _, f := range list {
				if f.IsDir() {
					var child model.LogFile
					if child, err = c.getFolderList(fmt.Sprint(path, "/", f.Name())); err != nil {
						return
					}
					logFile.Child = append(logFile.Child, child)
				} else {
					logFile.Child = append(logFile.Child, model.LogFile{Name: f.Name(), Child: []model.LogFile{}})
				}
			}
		}
	} else {
		logFile.Name = fileInfo.Name()
	}
	return logFile, nil
}
