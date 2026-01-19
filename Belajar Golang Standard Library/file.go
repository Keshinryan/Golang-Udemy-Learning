package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name,os.O_CREATE|os.O_WRONLY,0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile (name string)(string,error){
    file,err:=os.OpenFile(name, os.O_RDONLY,0666)
    if err!=nil{
        return "", err
    }
    defer file.Close()

    reader:=bufio.NewReader(file)
    var message string
    for{
        line,_,err:=reader.ReadLine()
        if err==io.EOF{  // Changed: should be == not !=
            break
        }
        message+=string(line)+"\n"
    }
    return message,nil
}

func addtoFile(name string, message string) error {
	file, err := os.OpenFile(name,os.O_APPEND|os.O_RDWR,0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func main() {
    // createNewFile("newfile.txt", "Hello, this is a new file created using os.OpenFile in Go!")
	
	addtoFile("newfile.txt", "This is an appended line.\n")

	result, _ := readFile("newfile.txt")  
    fmt.Println(result)

}