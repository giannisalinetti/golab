package main

import "fmt"

func printSlice(s []string) {
    fmt.Printf("[ ")
    for i, m := range s {
        if i != len(s) -1  {
            fmt.Printf("'%s', ", m)
        } else {
            fmt.Printf("'%s' ", m)
        }
    }
    fmt.Printf("]\n")
}

func reverseSlice(s []string) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {  // The condition i < j is necessary to not reverse back the list.
        s[i], s[j] = s[j], s[i]
    }
}


func main() {

    // Full year array. Index 0 is empty
    months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}

    // Slices of quarters
    Q1 := months[1:4]
    Q2 := months[4:7]
    Q3 := months[7:10]
    Q4 := months[10:13]

    fmt.Println("This is the full months array. Note the first empty string.")
    for _, m := range months {  
        fmt.Printf("%s\n", m)  // Index 0 will be printed as an empty line
    }
   
    fmt.Println("Now let's print the slices.") 
    printSlice(Q1)
    printSlice(Q2)
    printSlice(Q3)
    printSlice(Q4)

    // Reverse the first quarter
    fmt.Println("Now let's reverse the first quarter slice")
    reverseSlice(Q1)  // Passing Q1 to the function. Q1 is a pointer reference to the array.
    printSlice(Q1)  // Q1 has been reversed
}

    
