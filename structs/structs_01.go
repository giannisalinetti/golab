package main

import "fmt"

func main() {
  type myStruct struct {
    name string
    lastname string
    age int
  }

  var me myStruct
  me.name = "Gianni"
  me.lastname = "Salinetti"
  me.age = 36

  fmt.Println(me.name)
  fmt.Println(me.lastname)
  fmt.Println(me.age)
}
