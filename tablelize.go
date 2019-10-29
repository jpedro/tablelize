package tablelize

import(
  "fmt"
  "reflect"
  "strings"
  "strconv"
)

const (
  ALIGN_NUMBER = 0
  ALIGN_STRING = 1
)

func Rows(data [][]interface{}) {
  var widths []int
  var aligns []int

  started := false

  for i := range(data) {
    row := data[i]

    if started == false {
      // val := reflect.ValueOf(row)
      // fmt.Println(val.Kind())
      len := len(row)
      // fmt.Println(len)
      widths = make([]int, len)
      aligns = make([]int, len)
      started = true
    }

    for j := range(row) {
      // val := row[j]
      val := reflect.ValueOf(row[j])
      len := lenValue(val)
      if widths[j] < len {
        widths[j] = len
      }

      // 1. Skip the header
      // 2. Assume that they all are numeric, because you want to:
      // 3. Stop as soon as you find one non-numeric value and force align them
      //    all to the left
      if i > 0 && aligns[j] == ALIGN_NUMBER {
        slak := reflect.TypeOf(val).String()
        switch slak {
        case "int":
        case "float23":
        case "float64":
        case "string":
          if isNumeric(val) == false {
            aligns[j] = ALIGN_STRING
            break
          }

          // Keep it as ALIGN_NUMBER
        default:
          aligns[j] = ALIGN_STRING
        }

      }
    }
  }

  format := ""
  align := "-"
  for i := range(widths) {
    align = "-"
    if aligns[i] == ALIGN_NUMBER {
      align = ""
      format = fmt.Sprintf("%s%%%s%ds   ", format, align, widths[i])
    } else {
      format = fmt.Sprintf("%s%%%s%ds   ", format, align, widths[i])
    }
    fmt.Println(format)
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

func isNumeric(val interface{}) bool {
  switch val.(type) {
  case int:
  case float32:
  case float64:
    return true
  case string:
    str, _ := val.(string)
    _, errFloat := strconv.ParseFloat(str, 64)
    _, errInt := strconv.ParseInt(str, 10, 64)

    if errFloat == nil || errInt == nil {
      return true
    }
  }
  return false
}

func lenValue(val interface{}) int {
  slak := reflect.TypeOf(val).Elem()
  fmt.Println(slak)
  switch slak {
  case string:
    return len(val.(string))
  case int:
    i := val.(int)
    return len(strconv.Itoa(i))
  // case "float64":
  //   f := val.(float64)
  //   s := fmt.Sprintf("%f", f)
  //   return len(s)
  // case "float32":
  //   f := val.(float32)
  //   s := fmt.Sprintf("%f", f)
  //   return len(s)
  }
  return 6
}
