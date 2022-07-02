package main

import (
	"context"
	"logger/logger"
)

func main() {
	c := context.Background()
	name := "Namsdgfasdfaefdafawfawdfae"
	logg := logger.New()

	logg.Info(c, "testing %s", name)

	name = "Name"
	logg.Error(c, "testing %s", name)
	// logObj := logger.RetrievePackage()
	// fmt.Println(logObj.PackageName, logObj.FileName, logObj.FuncName, logObj.Line)
}
