package main

import(
  "fmt"
)

func main()  {

  b := 25
  a:= &b
  fmt.Println(a)
  fmt.Println(*a)

}
