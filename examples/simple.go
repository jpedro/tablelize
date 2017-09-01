package main

import (
  "github.com/jpedro/tablelize"
)

func main() {
  var data [][]string

  data = append(data, []string{"KEY", "VALUE"})
  data = append(data, []string{"char", "a"})
  data = append(data, []string{"medium", "Some text with content"})
  data = append(data, []string{"long", "And now for something very very loooong"})

  tablelize.Rows(data)
}
