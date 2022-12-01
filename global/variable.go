// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package global

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/xjustloveux/jgo/jconf"
	"os"
)

var (
	Dev      = os.Getenv("jEnv") == "dev"
	Conf     = jconf.New()
	Router   = getRouter()
	Validate = validator.New()
)

func getRouter() *gin.Engine {

	if !Dev {

		gin.SetMode(gin.ReleaseMode)
	}
	return gin.Default()
}
