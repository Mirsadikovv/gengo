package stuble

import _ "embed"

// module — entity-level templates
//go:embed module/cmd.tpl
var Cmd string

//go:embed module/dto.tpl
var Dto string

//go:embed module/model.tpl
var Model string

//go:embed module/service.tpl
var Service string

//go:embed module/handler.tpl
var Handler string

// project — project root templates
//go:embed project/main_entry.tpl
var MainEntry string

//go:embed project/dev.tpl
var Dev string

//go:embed project/prod.tpl
var Prod string

//go:embed project/src_main.tpl
var SrcMain string

//go:embed project/env.tpl
var Env string

//go:embed project/gitignore.tpl
var Gitignore string

//go:embed project/makefile.tpl
var Makefile string

// auth — auth_service bootstrap templates
//go:embed auth/auth_dto.tpl
var AuthDto string

//go:embed auth/auth_middleware.tpl
var AuthMiddleware string
