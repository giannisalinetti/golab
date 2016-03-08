package main

import "fmt"

// This demonstrates that an init function is always called, even with no code execued in main
func init() { fmt.Printf("Hello ") }
func init() { fmt.Printf("World\n") }  // Multiple init() functions can be called.

func main() {}

