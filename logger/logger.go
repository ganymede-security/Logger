package logger

import (
	"context"
	"path"
	"runtime"
	"strings"
)

type (
	Logger interface {
		Debug(c context.Context, format string, args ...interface{})
		Info(c context.Context, format string, args ...interface{})
		Warning(c context.Context, format string, args ...interface{})
		Error(c context.Context, format string, args ...interface{})
		Critical(c context.Context, format string, args ...interface{})
	}

	loggerFactory func() Logger
)

var (
	New loggerFactory
)

type callInfo struct {
	PackageName string
	FileName    string
	FuncName    string
	Line        int
}

func RetrievePackage() *callInfo {
	pc, file, line, _ := runtime.Caller(1)
	_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}
	return &callInfo{
		PackageName: packageName,
		FileName:    fileName,
		FuncName:    funcName,
		Line:        line,
	}
}
