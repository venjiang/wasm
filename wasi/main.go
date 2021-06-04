package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	// "github.com/mathetake/gasm/wasi"
	// "github.com/mathetake/gasm/wasm"
	// "github.com/go-interpreter/wagon/exec"
	// "github.com/go-interpreter/wagon/wasm"
	"github.com/mathetake/gasm/wasi"
	"github.com/mathetake/gasm/wasm"
	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
	// wagon()
	// gasm()
	// loadGoWasm()
	loadRustWasm()
}

func loadGoWasm() {
	wasmBytes, _ := ioutil.ReadFile("../pkg/wasi.wasm")
	// Create an Engine
	engine := wasmer.NewEngine()
	// Create a Store
	store := wasmer.NewStore(engine)
	// Compiles the module
	module, err := wasmer.NewModule(store, wasmBytes)
	if err != nil {
		fmt.Println("Failed to compile module:", err)
	}

	// Create an empty import object.
	importObject := wasmer.NewImportObject()
	// Let's instantiate the Wasm module.
	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		panic(fmt.Sprintln("Failed to instantiate the module:", err))
	}
	// get function
	multiply, err := instance.Exports.GetFunction("multiply")
	if err != nil {
		panic(fmt.Sprintln("Failed to retrieve the `multiply` function:", err))
	}

	result, err := multiply(3, 5)
	fmt.Println(result)

}

func loadRustWasm() {
	wasmBytes, _ := ioutil.ReadFile("../pkg/triple-rust.wasm")
	// engine
	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)
	// Compiles the module
	module, err := wasmer.NewModule(store, wasmBytes)
	if err != nil {
		fmt.Println("Failed to compile module:", err)
	}
	// Instantiates the module
	importObject := wasmer.NewImportObject()
	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		panic(fmt.Sprintln("Failed to instantiate the module:", err))
	}
	// Gets the `sum` exported function from the WebAssembly instance.
	triple, _ := instance.Exports.GetFunction("triple")
	var data float64 = 5
	result, err := triple(data)
	if err != nil {
		panic(fmt.Sprintln("Failed to get the `add_one` function:", err))
	}
	fmt.Println(result)
}

// func wagon() {
// 	f, _ := os.Open("../static/main.wasm")
// 	m, _ := wasm.ReadModule(f, nil)
// 	vm, _ := exec.NewVM(m)
// 	// result = hello(100, 20)
// 	// msg := "调用wasm"
// 	result, err := vm.ExecCode(0, 1)
// 	fmt.Printf("ret: %v, err: %v", result, err)
// }

func gasm() {
	buf, _ := ioutil.ReadFile("../pkg/main.wasm")
	mod, _ := wasm.DecodeModule(bytes.NewBuffer(buf))
	vm, _ := wasm.NewVM(mod, wasi.New().Modules())
	// call
	// msg := "调用wasm方法"
	// arg := &msg
	ret, retTypes, err := vm.ExecExportedFunction("hello", 5)
	fmt.Println("ret: %v, retTypes: %v, err: %v", ret, retTypes, err)

}
