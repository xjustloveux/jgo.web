// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package job

import (
	"github.com/xjustloveux/jgo/jcron"
)

type job interface {
	Name() string
	Run(map[string]interface{})
}

func Init() error {

	jobs := []job{
		&job001{},
		&job002{},
	}
	for _, v := range jobs {

		if err := jcron.AddJob(v.Name(), v); err != nil {

			return err
		}
	}
	return nil
}
