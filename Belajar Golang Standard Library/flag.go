package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your database host")
	var username *string = flag.String("username", "root", "Put your database username")
	var password *string = flag.String("password", "root", "Put your database password")
	var port *int = flag.Int("port",0,"Put your database port")

	flag.Parse()

	fmt.Println("Host",*host)
	fmt.Println("Username",*username)
	fmt.Println("Password",*password)
	fmt.Println("Port",*port)
	//how to run go run flag.go -username=jason -password="rahasia banget" -host=1.2.2.3 -port=3306
}