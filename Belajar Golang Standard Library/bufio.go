package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader("this is long string\nwith new line\n")
	reader:=bufio.NewReader(input)

	for{
		line,_,err:=reader.ReadLine()
		if err!=nil{
			break
		}
		fmt.Println(string(line))
	}

	writer:= bufio.NewWriter(os.Stdout)
	_,_ = writer.WriteString("hello world\n")
	_,_ = writer.WriteString("Selamat Belajar\n")
	writer.Flush()

}