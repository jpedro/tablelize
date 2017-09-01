# tablelize

Formats a table (array of arrays) of strings by width.

### Usage

```go
package main

import (
  "github.com/jpedro/tablelize"
)

func main() {
  var data [][]string
  var row []string

  row = []string{"KEY", "VALUE"}
  data = append(data, row)

  row = []string{"char", "a"}
  data = append(data, row)

  row = []string{"medium", "Some text with content"}
  data = append(data, row)

  row = []string{"long", "And now for something very very loooong"}
  data = append(data, row)

  tablelize.Rows(data)
}
```

Check https://github.com/jpedro/tablelize/tree/master/examples for examples.
