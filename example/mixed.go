package main

import (
  "github.com/jpedro/tablelize"
)

func main() {
  var data [][]string
  var row []string

  row = []string{"KEY", "ALL NUMBERS", "NOT ALL NUMBERS"}
  data = append(data, row)

  row = []string{"all numeric", "123", "123"}
  data = append(data, row)

  row = []string{"not numeric", "456", "456a"}
  data = append(data, row)

  tablelize.Rows(data)
}
