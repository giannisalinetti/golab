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

  // The struct can be printed using different string formatting
  // {Gianni Salinetti 36}
  fmt.Printf("%v\n", me)
  // {name:Gianni lastname:Salinetti age:36}
  fmt.Printf("%+v\n", me)
  // main.myStruct{name:"Gianni", lastname:"Salinetti", age:36}
  fmt.Printf("%#v\n", me)
  // Just prints the type: main.myStruct
  fmt.Printf("%T\n", me)
}
