package main

import (
  "fmt"
)

func main()  {

  if num:=10; num%2 ==0{
    fmt.Printf("Num %v is eve\n",num)
    fmt.Printf("Type is %T\n",num)
  }

  testNum:=5
  if testNum%2==0{
    fmt.Printf("Num %v is even\n",testNum)
  }else{
    fmt.Printf("Num %v is odd and value is %v\n",testNum,testNum/2)
  }

  for i:=10; i>0;i--{
    fmt.Printf("value if i is %d\n",i)
  }

  n:=5

  for i:=0;i<n;i++{
    for j:=0;j<=i;j++{
      fmt.Print("*")
    }
    fmt.Println()
  }

  no:=65
  switch  {
  case no>=90:
    fmt.Println("Grade is A")
  case no>=75 && no<90:
    fmt.Println("Grade is B")
  case no<75 && no >= 60:
    fmt.Println("Grade is C")
  case no>=45 && no<75:
    fmt.Println("Grade is D")
  default:
    fmt.Println("Fail")

  }


  switch num := number(); { //num is not a constant
    case num < 50:
        fmt.Printf("%d is lesser than 50\n", num)
        fallthrough
    case num < 100:
        fmt.Printf("%d is lesser than 100\n", num)
        fallthrough
    case num < 200:
        fmt.Printf("%d is lesser than 200\n", num)

}

colors := []string{"Red","Green","Yellow"}

for i:= range colors{
  fmt.Println(colors[i])
}

for _, v :=range colors{
  fmt.Println(v)
}

}
func number() int {
        num := 15 * 5
        return num
}
