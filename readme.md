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

  data = append(data, []string{"KEY", "VALUE"})
  data = append(data, []string{"A char", "a"})
  data = append(data, []string{"Some key", "With a loooooooooooonger value"})
  data = append(data, []string{"A long key with a small value", "Hi!"})

  tablelize.Rows(data)
}
```

Output:
```
% go run simple.go
KEY     VALUE
char    a
medium  Some text with content
long    And now for something very very loooong
```

Check [the directory examples](https://github.com/jpedro/tablelize/tree/master/examples)
for more exammples.
