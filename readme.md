# tablelize

Prints out the values of a table (array of arrays) of strings nicely align by
width.


### Usage

Source:
```go
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
```

Output:
```
% go run examples/main.go
KEY               VALUE                                        NUMBER   ALMOST_NUMBER
char              a                                                 1   1
longer-key-name   Some text                                        -2   2
key               And now for something completely different        3   3a
```

Check [examples](https://github.com/jpedro/tablelize/tree/master/example).
