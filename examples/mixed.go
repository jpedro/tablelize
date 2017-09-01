package main

import (
  "github.com/jpedro/tablelize"
)

func main() {
  var data [][]string

  data = append(data, []string{"KEY", "ALL NUMBERS", "NOT ALL NUMBERS"})
  data = append(data, []string{"all numeric", "123", "123"})
  data = append(data, []string{"not numeric", "456", "456a"})

  tablelize.Rows(data)
}
