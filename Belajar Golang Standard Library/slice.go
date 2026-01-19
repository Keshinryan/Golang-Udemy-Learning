package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"jason", "budi", "toni", "joko"}
	values := []int{90, 80, 85, 70}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "joko"))
	fmt.Println(slices.Contains(values, 100))
	fmt.Println(slices.Index(names, "toni"))
	fmt.Println(slices.Index(values, 85))

}