package main

// The print() function is a built-in function in Go that doesn't require importing fmt
// It provides basic printing functionality directly to standard output
// For more formatting options, the fmt package would be needed
// print() is simpler than fmt.Print() but has fewer capabilities
// Built-in functions that don't require imports:
// print() - prints to stdout without formatting
// println() - prints to stdout with newline
// len() - returns length of strings, arrays, slices, maps, channels
// cap() - returns capacity of arrays, slices, channels
// make() - allocates and initializes slices, maps, channels
// new() - allocates memory for a type
// append() - grows slices
// copy() - copies slice elements
// delete() - removes map entries
// close() - closes channels
// complex() - constructs complex numbers
// real() - gets real part of complex number
// imag() - gets imaginary part of complex number
// panic() - stops normal execution
// recover() - handles panics
// true, false - boolean constants
// nil - zero value for pointers, interfaces, maps, slices, channels

func main() {
	// Examples:
	print("This prints a line without carriage return. ")
	println("This prints with newline.")
	print("Characters like \\n \\ and etc are supported.\n")
	// Creates a slice of integers with length 3 and initial values of 0
	// make() allocates and initializes the slice
	// The length of 3 means it can hold 3 integers initially
	println("Create a slice (array) of integers, nums := make([]int, 3):")
	nums := make([]int, 3)
	println("Length is:")
	// len() returns the number of elements in the slice
	println(len(nums))
	println("Printing out a slice with a loop:")
	// Loop through nums slice and print each element
	for i := range nums {
		println("i =", nums[i])
	}
	// Creates a new empty map with string keys and integer values
	// make() allocates and initializes the map data structure
	println("Create an empty map (dictionary, m := make(map[string]int):")
	m := make(map[string]int)
	println("Length of map m is:", len(m))
	println("Printing out m with a loop:")
	// To print map contents without fmt, we need to iterate through key-value pairs
	for k, v := range m {
		println("Key:", k, "Value:", v)
	}
	// Fill in a map with some mock data
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	println("Filled in a map with mock data. Length is:", len(m))
	println("Printing out m with a loop:")
	for k, v := range m {
		println("Key:", k, "Value:", v)
	}
	// Deletes the entry with key "key" from the map
	// If the key doesn't exist, delete() has no effect
	// delete() is a built-in function that removes map entries
	println("Invoke delete(m, \"two\")")
	delete(m, "two")
	println("Printing out m with a loop again:")
	// To print map contents without fmt, we need to iterate through key-value pairs
	for k, v := range m {
		println("Key:", k, "Value:", v)
	}
	println("Creating a complex number c:")
	// complex() is a built-in function that creates a complex number from two float values
	// The first argument (1.0) becomes the real part
	// The second argument (2.0) becomes the imaginary part
	// This creates a complex number equivalent to 1.0 + 2.0i
	c := complex(1.0, 2.0)
	println(c)
	// Print the real part
	println(real(c))
}
