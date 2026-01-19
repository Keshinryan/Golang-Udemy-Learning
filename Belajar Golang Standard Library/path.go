package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	fmt.Println(path.Dir("hello/world.go"))
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Ext("hello/world.go"))
	fmt.Println(path.Join("hello", "world", "file.go"))


	fmt.Println(filepath.Dir("hello\\world.go"))
	fmt.Println(filepath.Base("hello\\world.go"))
	fmt.Println(filepath.Ext("hello\\world.go"))
	fmt.Println(filepath.IsAbs("hello\\world.go"))
	fmt.Println(filepath.IsAbs("C:\\hello\\world.go"))
	fmt.Println(filepath.IsLocal("hello\\world.go"))
	fmt.Println(filepath.Join("hello", "world", "file.go"))
}