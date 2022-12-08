// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package logc

import (
	"github.com/xjustloveux/jgo/jlog"
)

func SqlLog(args ...interface{}) {
	jlog.Info(args...)
}
