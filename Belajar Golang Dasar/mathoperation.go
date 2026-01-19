package main
import "fmt"

func main(){
	var a =10
	var b =3
	var hasilTambah = a + b
	var hasilKurang = a - b
	var hasilKali = a * b
	var hasilBagi = a / b
	var hasilSisaBagi = a % b
	fmt.Println("Hasil Tambah =", hasilTambah)
	fmt.Println("Hasil Kurang =", hasilKurang)
	fmt.Println("Hasil Kali =", hasilKali)
	fmt.Println("Hasil Bagi =", hasilBagi)
	fmt.Println("Hasil Sisa Bagi =", hasilSisaBagi)

	//augmented operator
	a+=2
	fmt.Println("a setelah ditambah 2 =", a)

	b-=1
	fmt.Println("b setelah dikurang 1 =", b)

	a*=3
	fmt.Println("a setelah dikali 3 =", a)

	b/=2
	fmt.Println("b setelah dibagi 2 =", b)	

	a%=4
	fmt.Println("a setelah dimodulus 4 =", a)

	//unary operator

	a++
	fmt.Println("a setelah diincrement =", a)
	b--
	fmt.Println("b setelah didecrement =", b)

}