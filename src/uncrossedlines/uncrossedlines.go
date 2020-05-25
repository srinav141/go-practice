/*

We write the integers of A and B (in the order they are given) on two separate horizontal lines.

Now, we may draw connecting lines: a straight line connecting two numbers A[i] and B[j] such that:

A[i] == B[j];
The line we draw does not intersect any other connecting (non-horizontal) line.
Note that a connecting lines cannot intersect even at the endpoints: each number can only belong to one connecting line.

Return the maximum number of connecting lines we can draw in this way.



Example 1:
Input: A = [1,4,2], B = [1,2,4]
Output: 2
Explanation: We can draw 2 uncrossed lines as in the diagram.
We cannot draw 3 uncrossed lines, because the line from A[1]=4 to B[2]=4 will intersect the line from A[2]=2 to B[1]=2.
Example 2:

Input: A = [2,5,1,2,5], B = [10,5,2,1,5,2]
Output: 3
Example 3:

Input: A = [1,3,7,1,7,5], B = [1,9,2,5,1]
Output: 2


Note:

1 <= A.length <= 500
1 <= B.length <= 500
1 <= A[i], B[i] <= 2000

HINT: longest sub-sequence

*/

package main

import (
	"fmt"
)

func maxUncrossedLines(A []int, B []int) int {
	la := len(A)
	lb := len(B)


m := make([][] int ,la+1)
for i:=range m{
	m[i] = make([]int,lb+1)
}



for i:=1;i<=la;i++{
	for j:=1;j<=lb;j++{
		if A[i-1] == B[j-1]{
			m[i][j] = m[i-1][j-1]+1
		}else{
			m[i][j] = getMax(m[i][j-1],m[i-1][j])
		}
	}
}

	return m[la][lb]
}

func getMax(a,b int) int{
	if a >b{
		return a
	}
	return b
}

func main(){

	A:= []int{1,1,2,1,2}
	B:= []int{1,3,2,3,1}

	res := maxUncrossedLines(A,B)
	fmt.Println(res)
}