package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("primes:", primes)

	var primes_slice []int = primes[1:4]
	fmt.Println("primes_slice:", primes_slice)
	fmt.Println("primes_slice cap:", cap(primes_slice))
	primes_slice = append(primes_slice, 17, 19, 23)
	fmt.Println("primes_slice:", primes_slice)
	fmt.Println("primes:", primes)
	fmt.Println()

	var slc2 []int = make([]int, 4)
	fmt.Println(slc2)
	fmt.Println(cap(slc2))
	for i := 0; i < 4; i++ {
		slc2[i] = 2*i + 2
	}
	fmt.Println(slc2)
	fmt.Println()

	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.	// Other slices that share the same underlying array will see those changes
	// Other slices that share the same underlying array will see those changes.
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	slc3 := []int{5, 10, 15}
	fmt.Println(slc3)
	fmt.Println()

	slc_part1 := []bool{true, false}
	slc_part2 := []bool{false, true}
	slc_totaal := append(slc_part1, slc_part2...)
	fmt.Printf("slc_totaal: %v\n\n", slc_totaal)
}
