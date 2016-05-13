// Example taken from GOPL.IO D&K

package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    seen := make(map[string]bool)        // Map creation with string keys and bool values
    input := bufio.NewScanner(os.Stdin)  // Using stdin for lines input
    for input.Scan() {                   // Scan a new line
        line := input.Text()  
        if !seen[line] {                 // If line is not already printed its initial zero value is false
            seen[line] = true
            fmt.Println(line)
        }
    }
    
    if err := input.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
        os.Exit(1)
    }
}
 
