package main

import (
  "fmt"
)

func main()  {
  fmt.Println("This is varibles test main start")
  var test1 int
  fmt.Println(test1)
  test1 = 10
  fmt.Println(test1)
  var test2 string = "test2"
  fmt.Println(test2)
  test2 = "re-declare test2"
  fmt.Println(test2)
  test3 := "test3"
  fmt.Println(test3)

  var test4 = 6
  fmt.Println(test4)

  var Test6 = 9
  fmt.Println(Test6)
  test_func(Test6)
  fmt.Println(Test6)
}

func test_func(testo int )  {
  fmt.Println(testo)
  testo += 10
  fmt.Println(testo)

}
