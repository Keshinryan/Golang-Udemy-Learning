package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	
	_ = writer.Write([]string{"name", "age", "city"})
	_ = writer.Write([]string{"eko", "30", "jakarta"})
	_ = writer.Write([]string{"budi", "25", "bandung"})
	_ = writer.Write([]string{"toni", "35", "surabaya"})

	writer.Flush()
}