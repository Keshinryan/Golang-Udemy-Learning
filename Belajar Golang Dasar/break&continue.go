package main

func main() {
	// Break digunakan untuk menghentikan perulangan
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}	
		println("Perulangan ke", i)
	}
	println("Selesai")	

	// Continue digunakan untuk melewati perulangan tertentu
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			continue
		}
		println("Perulangan ke", i)
	}
	println("Selesai")
}