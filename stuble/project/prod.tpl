//go:build !dev

package main

import (
	"{{ .ModPath }}/src"
	"git.sriss.uz/shared/shared_service/sharedutil"
)

func init() {
	sharedutil.Load(new(src.Env))
}
