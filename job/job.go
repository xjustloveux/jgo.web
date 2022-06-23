// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package job

import (
	"github.com/xjustloveux/jgo/jcron"
	"github.com/xjustloveux/jgo/jlog"
)

type job interface {
	Name() string
	Init()
}

var (
	jobList []jcron.Job
)

func init() {
	jobList = []jcron.Job{new(job001), new(job002)}
}

func CreateJob() {
	for _, j := range jobList {
		tj := j.(job)
		tj.Init()
		if err := jcron.AddJob(tj.Name(), j); err != nil {
			jlog.Error(err)
		}
	}
}
