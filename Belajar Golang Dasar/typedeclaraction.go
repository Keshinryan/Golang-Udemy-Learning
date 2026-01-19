package main
import "fmt"

func main(){
	type NoKTP string
	var KTP NoKTP="1111111111111111"
	var contoh string="2222222222222222"

	var KTP2 NoKTP=NoKTP(contoh)
	fmt.Println("No KTP =", KTP)
	fmt.Println("No KTP2 =", KTP2)
}