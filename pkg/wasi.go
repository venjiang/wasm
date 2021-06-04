package main

// This calls a JS function from Go.
func main() {
	// println("adding two numbers:", add(2, 3)) // expecting 5
	println("load tinygo wasm")
}

// This function is exported to JavaScript, so can be called using
// exports.multiply() in JavaScript.
//export multiply
func multiply(x, y int) int {
	return x * y
}
