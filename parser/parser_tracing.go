package parser

import (
	"fmt"
	"runtime"
	"strings"
)

var traceLevel int = 0

const traceIdentPlaceholder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", identLevel(), fs)
}

func incIdent() { traceLevel = traceLevel + 1 }
func decIdent() { traceLevel = traceLevel - 1 }

func trace() (untrace func()) {
	incIdent()
	pc, _, _, _ := runtime.Caller(1)
	tracePrint("BEGIN " + runtime.FuncForPC(pc).Name())

	return func() {
		pc, _, _, _ := runtime.Caller(1)
		tracePrint("END " + runtime.FuncForPC(pc).Name())
		decIdent()
	}
}
