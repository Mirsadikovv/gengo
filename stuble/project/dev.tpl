//go:build dev

package main

import (
	"{{ .ModPath }}/src"
	"github.com/Mirsadikovv/gengo/shared_service/sharedutil"
)

func init() {
	sharedutil.MustLoad(new(src.Env), ".env.dev")
}
