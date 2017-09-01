package main

import (
  "github.com/jpedro/tablelize"
)

func main() {
  var data [][]string
  var row []string

  row = []string{"KEY", "VALUE"}
  data = append(data, row)

  row = []string{"long value", "And now for something very very loooong"}
  data = append(data, row)

  row = []string{"And now for a very very loooong key", "ok"}
  data = append(data, row)

  tablelize.Rows(data)
}
