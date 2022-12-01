// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package b103srv

import (
	"fmt"
	"github.com/xjustloveux/jgo.web/model/b/b103mod"
	"github.com/xjustloveux/jgo.web/model/globalmod"
	"github.com/xjustloveux/jgo.web/service/globalsrv"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jlog"
	"os"
	"path/filepath"
	"strings"
)

func QueryContent() ([]globalmod.PkgContent, error) {

	return globalsrv.QueryContent("B103")
}

func QueryLog() (b103mod.ResLog, error) {

	path := jcast.String(jlog.GetParam("path"))
	return getFolderList(path)
}

func getFolderList(path string) (log b103mod.ResLog, err error) {

	path = strings.Trim(strings.Trim(path, "\\ "), "/ ")
	log.Child = make([]b103mod.ResLog, 0)
	var fileInfo os.FileInfo
	if fileInfo, err = os.Stat(path); err != nil {

		return
	}
	log.Dir = fileInfo.IsDir()
	if log.Dir {

		log.Name = filepath.Base(fileInfo.Name())
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

					var child b103mod.ResLog
					if child, err = getFolderList(fmt.Sprint(path, "/", f.Name())); err != nil {

						return
					}
					log.Child = append(log.Child, child)
				} else {

					log.Child = append(log.Child, b103mod.ResLog{Name: f.Name(), Child: []b103mod.ResLog{}})
				}
			}
		}
	} else {

		log.Name = fileInfo.Name()
	}
	return log, nil
}
