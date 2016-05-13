// Comparing two structs of the same type
package main

import "fmt"

func main() {

  type myStruct struct {
    Name string
    Lastname string
    Age int
  }

  user1 := myStruct{"Gianni", "Salinetti", 36}
  user2 := myStruct{"Mario", "Rossi", 44}

  // If all field of the structs are comparable, then structs are comparable
  // So this output (True)
  fmt.Println(user1.Name == user2.Name && user1.Lastname == user2.Lastname && user1.Age == user2.Age)
  // Is identical to this
  fmt.Println(user1 == user2)
}
