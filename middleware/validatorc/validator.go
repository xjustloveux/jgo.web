// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package validatorc

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/xjustloveux/jgo.web/global"
	"github.com/xjustloveux/jgo/jlog"
)

func Verify(data interface{}) error {

	if err := global.Validate.Struct(data); err != nil {

		errs := err.(validator.ValidationErrors)
		for _, vErr := range errs {

			var translator ut.Translator
			errStr := fmt.Sprintf("verify error：%v info(%v.%v) struct(%v) value(%v)",
				vErr.Translate(translator),
				vErr.Field(),
				vErr.Tag(),
				vErr.StructNamespace(),
				vErr.Value())
			jlog.Error(errStr)
		}
		return err
	}
	return nil
}
