package main

import "fmt"

func main() {

  // Struct declaration. Field are exported with capital letter
  type myStruct struct {
    Name string
    Lastname string
    Job string
    Age int
  }

  // Struct literal type 2: only desided fields are declared using key/value
  // Undeclared field are left to their corresponding null value
  me := myStruct{Name: "Gianni"}

  // Other fields can be updated after initialization
  me.Job = "Instructor"
  // This will print only Gianni, null string for me.Lastname, 0 for
  // me.Age and the content of me.Job
  fmt.Println(me.Name, me.Lastname, me.Age, me.Job)

  me.Lastname = "Salinetti"
  me.Age = 36

  // Print updated fields
  fmt.Println(me.Name, me.Lastname, me.Age, me.Job)
}
