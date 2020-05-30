package main

import (
	"fmt"
	"sort"
)

func canFinish(numCourses int, prerequisites [][]int) bool {

	s := make(map[int][]int)
	for i := 0; i < len(prerequisites); i++ {
		k := prerequisites[i][0]
		v := prerequisites[i][1]
		s[k] = append(s[k], v)

	}
	for i:=0 ; i<numCourses;i++{
		if isPrereq(i,s[i],s){
			return false
		}
	}
	fmt.Println(s)

	return true
}

func isPrereq(course int,prereq []int,sm map[int][]int) bool{

for _,v:= range prereq{

	if sl, ok:= sm[v];ok {
		if isContains(course,sl){
			return true
		}
		return isPrereq(course,sl,sm)
	}
}
	return false
}

func isContains(val int, sl []int)  bool {
	r := sort.SearchInts(sl,val)
	if r < len(sl) && sl[r] == val{
		return true
	}

	return false
}

func main() {
	s := [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 0}}
	//s = [][]int {{1,0},{0,2},{2,1}}
	s = [][]int {{0,1},{3,1},{1,3},{3,2}}
	r := canFinish(4, s)
	fmt.Println(r)
}
