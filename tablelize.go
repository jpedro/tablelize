package tablelize

import(
  "fmt"
  "strings"
)

func Rows(data [][]string) {
  widths := make([]int, 6)
  for i := range(data) {
    row := data[i]
    for j := range(row) {
      val := row[j]
      len := len(val)
      if widths[j] < len {
        widths[j] = len
      }
    }
  }

  format := ""
  for i := range(widths) {
    format = fmt.Sprintf("%s%%-%ds  ", format, widths[i])
  }
  format = strings.Trim(format, " ")

  for i := range(data) {
    row := data[i]
    args := make([]interface{}, len(row), len(row))
    for i := range row {
        args[i] = row[i]
    }
    fmt.Printf(format + "\n", args...)
  }
}
