package main

import (
	"fmt"
	"os"
)

func loglnf(fmtStr string, vals ...interface{}) {
	fmt.Printf("%s\n", fmt.Sprintf(fmtStr, vals...))
}

func logln(i interface{}) {
	loglnf("%s", i)
}

func exitf(fmtStr string, vals ...interface{}) {
	loglnf(fmtStr, vals...)
	os.Exit(1)
}

func exitln(i interface{}) {
	logln(i)
	os.Exit(1)
}
