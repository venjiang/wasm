package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/wasi"
	// "github.com/go-interpreter/wagon/exec"
	// "github.com/go-interpreter/wagon/wasm"
	"github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
	// loadTinyGoWasmer()
	// loadRustWasmer()
	loadTinyGoWazero()
}

func loadTinyGoWasmer() {
	wasmBytes, _ := os.ReadFile("../pkg/wasi.wasm")
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

func loadRustWasmer() {
	wasmBytes, _ := os.ReadFile("../pkg/triple-rust.wasm")
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

func loadTinyGoWazero() {
	buf, _ := os.ReadFile("../pkg/wasi.wasm")

	// Choose the context to use for function calls.
	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	runtime := wazero.NewRuntime()
	defer runtime.Close(ctx) // This closes everything this Runtime created.

	// wasi.wasm was compiled with TinyGo, which requires being instantiated as
	// a WASI command (to initialize memory).
	if _, err := wasi.InstantiateSnapshotPreview1(ctx, runtime); err != nil {
		fmt.Println("Failed to instantiate WASI:", err)
	}

	code, err := runtime.CompileModule(ctx, buf, wazero.NewCompileConfig())
	if err != nil {
		fmt.Println("Failed to compile module:", err)
	}

	module, err := runtime.InstantiateModule(ctx, code, wazero.NewModuleConfig())
	if err != nil {
		panic(fmt.Sprintln("Failed to instantiate the module:", err))
	}

	multiply := module.ExportedFunction("multiply")
	if multiply == nil {
		panic("Failed to locate the `multiply` function")
	}

	result, err := multiply.Call(nil, 3, 5)
	if err != nil {
		panic(fmt.Sprintln("Failed to invoke the `multiply` function:", err))
	}

	fmt.Println(result)
}
