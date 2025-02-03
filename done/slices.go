package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("unint", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len: ", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len: ", len(s))

	//  Append into array and save into a variable
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	// Copy into new array
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy: ", c)

	//  Slice Operator with slice[low:high]
	l := s[2:5]

	fmt.Println(l)

	// slice up to (excluding) 5
	l = s[:5]
	fmt.Println("sl2", l)
	l = s[2:]
	fmt.Println("sl3", l)

	t := []string{"g", "h", "i"}

	t2 := []string{"g", "h", "i"}

	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
} // end of main
