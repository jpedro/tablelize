package main

import (
  "github.com/jpedro/tablelize"
)

func main() {
  var data [][]string

  data = append(data, []string{"KEY", "VALUE", "NUMBER", "ALMOST_NUMBER"})
  data = append(data, []string{"char", "a", "1", "1"})
  data = append(data, []string{"longer-key-name", "Some text", "-2", "2"})
  data = append(data, []string{"key", "And now for something completely different", "3", "3a"})

  tablelize.Rows(data)
}
