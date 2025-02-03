package main

import "fmt"

func main() {
	i := 1
	// most basic with single condition
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}


	// A classic initial/condition/after for loops
	for j:= 0; j < 3; j++ {
		fmt.Println("j",j)
	}

	// N times iteration is range over an integer
	for i := range 3 {
		fmt.Println("range", i)
	}

	// for loop without condition and break the loop using break
	for {
		fmt.Println("loop")
		break
	}

	// for loop with continue for the next iteration
	for n:= range 6 {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}
}