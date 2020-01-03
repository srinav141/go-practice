package main

import(
  "fmt"
  "reflect"
)

func init()  {
  fmt.Println("This is type test init")
}

func init()  {
  fmt.Println("This is type test second init")
}

func main()  {
  a:="test"
  fmt.Println(reflect.TypeOf(a))
  b:=10
  fmt.Println(reflect.TypeOf(b))
  type myString string
  var customType myString ="custom"
  fmt.Println(reflect.TypeOf(customType))

}
