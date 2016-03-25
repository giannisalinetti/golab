package main

import "fmt"

func main() {
    
    // Make creates an empty slice of lenght 12. Capacity, if omitted is equal to lenght
    s01 := make([]int, 12)

    // The slice is filled with ints
    for i := 0; i <12; i++ {
        s01[i] = i*3
    }

    // The slice is printed. After printing the last number a newline character is added
    for i, value := range s01 {
        fmt.Printf("%d ", value)
        if i == len(s01)-1 {
            fmt.Printf("\n")
       }
    }
}

