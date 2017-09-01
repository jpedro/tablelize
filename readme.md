# tablelize

Prints out the values of a table (array of arrays) of strings nicely align by
width.


### Usage

```go
package main

import (
  "github.com/jpedro/tablelize"
)

func main() {
  var data [][]string
  var row []string

  data = append(data, []string{"KEY", "VALUE"})
  data = append(data, []string{"char", "a"})
  data = append(data, []string{"medium", "Some text with content"})
  data = append(data, []string{"long", "And now for something very very loooong"})

  tablelize.Rows(data)
}
```

Check [the directory examples](https://github.com/jpedro/tablelize/tree/master/examples)
for more exammples.
