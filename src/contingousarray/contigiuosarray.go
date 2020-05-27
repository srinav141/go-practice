package main

import "fmt"

func findMaxLength(nums []int) int {
	numLength := len(nums)
	if numLength == 0 || numLength ==1{
		return 0
	}

	maxLen:=0
	i:=0
	le :=0

	for i < numLength-1{
		le,i = getLength(nums,i)
		if le > maxLen{
			maxLen = le
		}
	}

	return maxLen
}

func getLength(nums []int,startIndex int) (le,ind int) {
	exp := nums[startIndex]
	l:=0

	for i:=startIndex;i<len(nums);i++{
		if nums[i] == exp{
			l++
		}else {
			ind = i
			break
		}
		ind = i
		if nums[i]==0 {
			exp =1
		}else{
			exp = 0
		}
	}
	if l%2 == 0{
		le = l
	}else{
		le = l-1
	}
	return
}


func main()  {

	n := []int {0,1,0}
	res := findMaxLength(n)
	fmt.Println(res)

}