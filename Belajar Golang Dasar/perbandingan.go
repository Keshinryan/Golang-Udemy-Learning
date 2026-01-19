package main


func main() {
	// operator perbandingan
	// >, <, >=, <=, ==, !=
	var a = 20
	var b = 20
	println("a =", a)
	println("b =", b)
	var hasilKebesaran = a > b
	println("a > b =", hasilKebesaran)
	var hasilKebesaranSamadengan = a >= b
	println("a >= b =", hasilKebesaranSamadengan)
	var hasilKekecilan = a < b
	println("a < b =", hasilKekecilan)
	var hasilKekecilanSamadengan = a <= b
	println("a <= b =", hasilKekecilanSamadengan)
	var hasilSamadengan = a == b
	println("a == b =", hasilSamadengan)
	var hasilTidakSamadengan = a != b
	println("a != b =", hasilTidakSamadengan)

	var name1 = "Jason"
	var name2 = "Jason"

	var result bool = name1 == name2

	println(result)

	result = hasilKebesaran && hasilKebesaranSamadengan
	println(result)

	result = hasilKebesaran || hasilKebesaranSamadengan
	println(result)

	result = !hasilKebesaranSamadengan
	println(result)
}