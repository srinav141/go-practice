package main

import (
	"fmt"
	"sort"
	"strings"
)

func frequencySort(s string) string {
 m:= make(map[string]int)

 for _,v:= range s{
 	if c, ok:= m[string(v)]; ok{
		m[string(v)] = c+1

	}else{
		m[string(v)] = 1
	}
 }

 set := make(map[int]bool)
 uv:= make([]int,0)

 for _,v:=range m{
 	 if !set[v]{
 	 	set[v] = true
 	 	uv = append(uv, v)
	 }

 }


 sort.Sort(sort.Reverse(sort.IntSlice(uv)))
 res:= ""
 for i:=0;i<len(uv);i++{
 	r:= uv[i]
for k,v := range m{
	if v == r{
		res = res+strings.Repeat(k,r)
	}
}
 }


	return res
}

func main()  {
res:= frequencySort("Aabb")

fmt.Println(res)
}