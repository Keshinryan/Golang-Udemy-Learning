package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type PersonSlice []Person

func (p PersonSlice) Len() int {
	return len(p)
}

func (p PersonSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PersonSlice) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func main() {
	person := []Person{
		{"Eko", 30},
		{"Jason", 22,},
		{"Erdi", 20,},
	}
	sort.Sort(PersonSlice(person))
	fmt.Println(person)
}