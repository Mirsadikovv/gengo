//go:build dev

package main

import (
	"{{ .ModPath }}/src"
	"git.sriss.uz/shared/shared_service/sharedutil"
)

func init() {
	sharedutil.MustLoad(new(src.Env), ".env.dev")
}
