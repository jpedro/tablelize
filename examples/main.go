package main

import (
  "github.com/jpedro/tablelize"
)

func main() {
  var data [][]interface{}

  data = append(data, []interface{}{"KEY", "VALUE", "NUMBER", "ALMOST_NUMBER"})
  data = append(data, []interface{}{"char", "a", 1, 1})
  data = append(data, []interface{}{"longer-key-name", "Some text", -2, 2})
  data = append(data, []interface{}{"key", "And now for something completely different", 3, "3a"})

  tablelize.Rows(data)

  // ok := [][]interface{}
  // ok = append(ok, []interface{"Ok", 2})
  // fmt.Printf("Len: %d\n", len(ok[0]))
}
