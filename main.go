package main

import (
	"github.com/SunA0/monkey/repl"
	"syscall/js"
)

func Run(this js.Value, args []js.Value) interface{} {
	message := args[0].String()

	res := repl.Common(message)
	return res
}
func main() {
	js.Global().Set("Run", js.FuncOf(Run))
	<-make(chan bool)
}
