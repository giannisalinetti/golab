// Freely taken from D&K book "The Go Programming Language" with some mods
package main

import (
    "fmt"
    "encoding/json"
    "log"
)

type Movie struct {
  Title string
  Year  int `json:"released"`   // Field tag for Year
  Color bool `json:"color,omitempty"` // Here with omitempty option for empty/zero values
  Actors []string
}

var movies = []Movie{
  {Title: "Metropolis", Year: 1927, Color: false,
   Actors: []string{"Brigitte Helm", "Gustav Frohlich"}},
  {Title: "Ghostbusters", Year: 1984, Color: true,
   Actors: []string{"Bill Murray, Dan Aykroyd"}},
}

func main() {

  // Marshaling in byte slice output, not so human readable
  data, err := json.Marshal(movies)
  if err != nil {
    log.Fatalf("JSON marshaling failed: %s", err)
  }
  fmt.Printf("%s\n", data)

  // Marshaling with Indent
  data, err = json.MarshalIndent(movies, "", "    ")
  if err != nil {
    log.Fatalf("JSON marshaling failed: %s", err)
  }
  fmt.Printf("%s\n", data)

  // Unmarshaling
  var titles []struct{ Title string}
  if err := json.Unmarshal(data, &titles); err != nil {
    log.Fatalf("JSON unmarshaling failed: %s", err)
  }

  // Print the titles array
  fmt.Printf("%+v\n", titles)
}
