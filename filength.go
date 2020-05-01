// Copyright (C) 2020 Toni Helminen <GPLv3>

package main
import (
  "fmt"
  "io/ioutil"
  "os"
)

func main() {
  if len(os.Args) <= 1 {
    fmt.Println("> filength myfile.txt")
    return
  }
  dat, err := ioutil.ReadFile(os.Args[1])
  if err != nil {panic(err)}
  slen := len(string(dat)) - 1
  if slen < 0 {fmt.Println("0")} else {fmt.Println(slen)}
}
