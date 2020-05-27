package main

import "fmt"

func findMaxLength(nums []int) int {
	maxLen:=0
	n := len(nums)
	m:= make(map[int]int)
	sub := make([] int,0)

	for _,v:=range  nums{
		val:=1
		if v ==0 {
			val = -1
		}
		sub = append(sub,val)
	}
sum:=0

for i:=0;i<n;i++{
	sum += sub[i]
	if sum == 0{
		maxLen = i+1
	}

	if v, ok:= m[sum+n];ok{
		if maxLen < i - v{
maxLen = i-v
		}
	}else{
		m[sum+n] = i
	}
}



		return maxLen
}

func main()  {

	n := []int {0,0,0,1,1,1,0}
	n = []int{0,1,1,1}
	res := findMaxLength(n)
	fmt.Println(res)

}
