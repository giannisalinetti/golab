// Struct embedding demonstration. With Anonymous Fields
// Anonymous Fields let create a syntactic shortcut to dot expressions
package main

import "fmt"

func main() {

  // Primitive struct for a two-dimensional point
  type Point struct {
    X int
    Y int
  }

  // The Circle struct uses a Point Anonymous Field
  type Circle struct {
    Point
    Radius int
  }

  // The Wheel struct uses a Circle Anonymous Field
  type Wheel struct {
    Circle
    Spokes int
  }

  // No need to use the full nested dot expressions
  var w Wheel
  w.X = 8
  w.Y = 5
  w.Radius = 25
  w.Spokes = 16

  // {Circle:{Point:{X:8 Y:5} Radius:25} Spokes:16}
  fmt.Printf("%+v\n", w)

  // Anonymous Fields won't apply to Struct Literals. Extended notation needed
  r := Wheel{Circle{Point{10, 10}, 30}, 24}
  //  {Circle:{Point:{X:10 Y:10} Radius:30} Spokes:24}
  fmt.Printf("%+v\n", r)
}
