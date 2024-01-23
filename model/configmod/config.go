// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package configmod

type Config struct {
	// Port gin listen port
	Port int
	// UrlSSR SSR網址
	UrlSSR string
	// CORS Cross-Origin Resource Sharing
	CORS []string
	// AutoStartCron auto start cron
	AutoStartCron bool
	// Assets path to assets
	Assets string
}
