package main

import (
	"flag"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/Mirsadikovv/gengo/shared_service/sharedutil"
	"github.com/Mirsadikovv/gengo/shared_service/version/parser"
)

func main() {
	var (
		dir        string
		work       bool
		newVersion string
	)

	flag.StringVar(&dir, "d", ".", "Application version")
	flag.StringVar(&newVersion, "v", "", "Set version")
	flag.BoolVar(&work, "work", false, "Application version")

	flag.Parse()

	defer func() {
		if r := recover(); r != nil {
			if work {
				println(path.Join(TrimPrefix(), "go.work:"), "not found")
			} else {
				println(path.Join(TrimPrefix(), "go.mod:"), "not found")
			}
		}
	}()

	switch {
	case work && newVersion == "":
		{
			workfile := sharedutil.MustValue(parser.WorkFile(path.Join(dir, "go.work")))

			println(path.Join(TrimPrefix(), "go.work"), "->", workfile.Go.Version)
		}

	case !work && newVersion == "":
		{
			modfile := sharedutil.MustValue(parser.ModFile(path.Join(dir, "go.mod")))

			println(path.Join(TrimPrefix(), "go.mod"), "->", modfile.Go.Version)
		}

	case work && newVersion != "":
		{
			sharedutil.Must(parser.UpdateGoVersionInWork(path.Join(dir, "go.work"), newVersion))
			println(path.Join(TrimPrefix(), "go.work update"), "->", newVersion)
		}

	case !work && newVersion != "":
		{
			sharedutil.Must(parser.UpdateGoVersionInModFile(path.Join(dir, "go.mod"), newVersion))
			println(path.Join(TrimPrefix(), "go.mod update"), "->", newVersion)
		}
	}
}

func TrimPrefix() string {
	fullpath := sharedutil.MustValue(os.Getwd())

	buildPath := strings.TrimPrefix(fullpath, filepath.Dir(fullpath))

	return strings.TrimPrefix(buildPath, "\\")
}
