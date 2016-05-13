package main

import "fmt"

func main() {

  // Struct declaration. Field are not exported
  type myStruct struct {
    name string
    lastname string
    age int
  }

  // Struct literal type 1: all field must be declared inline and in the same
  // order of the struct declaration
  me := myStruct{"Gianni", "Salinetti", 36}
  fmt.Println(me.name, me.lastname, me.age)
}
