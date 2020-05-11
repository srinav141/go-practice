/**

In a town, there are N people labelled from 1 to N.  There is a rumor that one of these people is secretly the town judge.

If the town judge exists, then:

The town judge trusts nobody.
Everybody (except for the town judge) trusts the town judge.
There is exactly one person that satisfies properties 1 and 2.
You are given trust, an array of pairs trust[i] = [a, b] representing that the person labelled a trusts the person labelled b.

If the town judge exists and can be identified, return the label of the town judge.  Otherwise, return -1.



Example 1:

Input: N = 2, trust = [[1,2]]
Output: 2
Example 2:

Input: N = 3, trust = [[1,3],[2,3]]
Output: 3
Example 3:

Input: N = 3, trust = [[1,3],[2,3],[3,1]]
Output: -1
Example 4:

Input: N = 3, trust = [[1,2],[2,3]]
Output: -1
Example 5:

Input: N = 4, trust = [[1,3],[1,4],[2,3],[2,4],[4,3]]
Output: 3


Note:

1 <= N <= 1000
trust.length <= 10000
trust[i] are all different
trust[i][0] != trust[i][1]
1 <= trust[i][0], trust[i][1] <= N

*/

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
