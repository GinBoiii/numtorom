package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
  "strings"
  "errors"
)

var roman = map[int]string {
  1000: "M",
  500: "D",
  100: "C",
  50: "L",
  10: "X",
  5: "V",
  1: "I",
}

var maxi = []int{
	1000,
	500,
	100,
	50,
	10,
	5,
	1,
}


func main() {
  readi := bufio.NewReader(os.Stdin)
  fmt.Println("Enter number: ")

  line, _ := readi.ReadString('\n')
  line =strings.TrimSuffix(line, "\n")

  num, _ := strconv.Atoi(line)
  if num <= 0 || num > 3000{
    fmt.Println(errors.New("can't return roman numeral"))
  } else {
    fmt.Println(goToRome(num))
  }
}

func goToRome(num int) string {
  ret := ""

  for num > 0 {
    val := hiVal(num)
    ret += roman[val]
    num -= val
  }

  ret = strings.ReplaceAll(ret, "IIII", "IV")
  ret = strings.ReplaceAll(ret, "VIIII", "IX")
  ret = strings.ReplaceAll(ret, "XXXX", "XL")
  ret = strings.ReplaceAll(ret, "LXXXX", "XC")
  ret = strings.ReplaceAll(ret, "CCCC", "CD")
  ret = strings.ReplaceAll(ret, "DCCCC", "CM")

  return ret
}

func hiVal(num int) int {
  for _ , val := range maxi {
    if num >= val{
      return val
    }
  }
  return 0
}
