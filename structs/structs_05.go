// Struct embedding demonstration. Without Anonymous Fields
package main

import "fmt"

func main() {

  // Primitive struct for a two-dimensional point
  type Point struct {
    X int
    Y int
  }

  // The Circle struct uses Point definition
  type Circle struct {
    Center Point
    Radius int
  }

  // The Wheel struct adds Spokes count to a wheel
  type Wheel struct {
    Circle Circle
    Spokes int
  }

  // When no Anonymous Fields are used struct members must be called
  // in an explicit, nested (and long) way
  var w Wheel
  w.Circle.Center.X = 8
  w.Circle.Center.Y = 5
  w.Circle.Radius = 25
  w.Spokes = 16

  // {Circle:{Center:{X:8 Y:5} Radius:25} Spokes:16}
  fmt.Printf("%+v\n", w)
}
