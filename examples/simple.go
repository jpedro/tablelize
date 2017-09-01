package main

import (
  "github.com/jpedro/tablelize"
)

func main() {
  var data [][]string

  data = append(data, []string{"KEY", "VALUE"})
  data = append(data, []string{"A char", "a"})
  data = append(data, []string{"Some key", "With a loooooooooooonger value"})
  data = append(data, []string{"A long key with a small value", "Hi!"})

  tablelize.Rows(data)
}
