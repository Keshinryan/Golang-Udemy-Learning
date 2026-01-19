package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Now:", now.Local())

	var utc time.Time = time.Date(2024, 6, 1, 10, 0, 0, 0, time.UTC)
	fmt.Println("UTC Time:", utc)
	fmt.Println("UTC to Local:", utc.Local())

	formatter := "2006-01-02 15:04:05"

	value:="2024-06-01 15:30:00"

	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Value Time:", valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())

	var duration1 time.Duration=time.Second*100
	var duration2 time.Duration=time.Millisecond*10
	var duration3 time.Duration=duration1-duration2
	fmt.Println(duration3)
	fmt.Printf("duration : %d\n",duration3)
}