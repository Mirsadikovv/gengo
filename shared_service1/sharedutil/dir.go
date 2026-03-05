package sharedutil

import (
	"path"
	"runtime"
)

func Dirname() string {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return ""
	}

	return path.Dir(filename)
}

func Filename() string {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return ""
	}

	return filename
}

func Basename() string {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return ""
	}

	return path.Base(filename)
}

func FunctionName() string {
	pc, _, _, ok := runtime.Caller(1)

	if !ok {
		return ""
	}

	fn := runtime.FuncForPC(pc)

	return fn.Name()
}
