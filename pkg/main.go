package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	done := make(chan int, 0)
	// alert
	// js.Global().Get("alert")
	// alert := js.Global().Get("alert")
	// alert.Invoke("Hello WASM 测试!")
	// hello
	fmt.Println("hello go wasm")
	js.Global().Set("hello", js.FuncOf(jsHello))
	<-done
}

func jsHello(this js.Value, args []js.Value) interface{} {
	msg := args[0].String()

	return js.ValueOf(hello(msg))
}

func hello(msg string) string {
	return "hello " + msg
}
