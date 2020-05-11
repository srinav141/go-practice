package main

import (
	"fmt"
	"reflect"
	"sort"
)

func hashIntersection(a interface{}, b interface{}) []interface{} {

	set := make([]interface{}, 0)
	hash := make(map[interface{}]bool)

	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	fmt.Println(av)
	fmt.Println(bv)

	for i := 0; i < av.Len(); i++ {
		e1 := av.Index(i).Interface()
		hash[e1] = true
	}

	for i := 0; i < bv.Len(); i++ {
		e1 := bv.Index(i).Interface()
		if _, ok := hash[e1]; ok {
			set = append(set, e1)
		}
	}

	return set
}

func findJudge(N int, trust [][]int) int {
	guest := -1
	if N == 1 {
		return 1
	}

	if len(trust) == 1 {
		return trust[0][1]
	}

	ppleMap := make(map[int][]int)

	for i := 1; i <= N; i++ {
		ppleMap[i] = make([]int, 0)
	}

	for i := 0; i < len(trust); i++ {

		s := append(ppleMap[trust[i][0]], trust[i][1])
		ppleMap[trust[i][0]] = s

	}
	fmt.Println(ppleMap)
	expectedGuest := make([]int, 0)
	for k, v := range ppleMap {
		if len(v) == 0 {
			expectedGuest = append(expectedGuest, k)
		}
	}

	if len(expectedGuest) == 1 {
		guest = expectedGuest[0]
	} else {
		return -1
	}

	for k, v := range ppleMap {
		if k != guest {
			sort.Ints(v)
			i := sort.SearchInts(v, guest)

			if i < len(v) {
				if v[i] != guest {
					return -1
				}
			} else if i == len(v) {
				return -1
			}
		}
	}

	// keys := make([]int, 0, len(ppleMap))
	// for k := range ppleMap {
	// 	keys = append(keys, k)
	// }

	// common := ppleMap[keys[0]]

	// for i := 1; i < len(keys); i++ {
	// 	if len(ppleMap[i]) > 0 {

	// 		c := hashIntersection(common, ppleMap[i])
	// 		fmt.Println(common)

	// 		if len(c) > 0 {
	// 			common = make([]int, 0)
	// 			for j := 0; j < len(c); j++ {

	// 				common = append(common, c[j].(int))
	// 			}

	// 		}

	// 	}
	// }

	// if len(common) > 0 && len(ppleMap[common[0]]) == 0 {
	// 	return common[0]
	// }

	return guest
}

func main() {
	s := [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}
	s = [][]int{{1, 2}, {2, 3}}
	res := findJudge(3, s)
	fmt.Println(res)

}
