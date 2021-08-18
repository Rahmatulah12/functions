package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func main() {
	// var names = []string{"bagus", "cahyo"}
	// printMessage("Hello selamat datang,", names)

	// rand.Seed(time.Now().Unix())
	// var randomValue int
	// randomValue = randomWithRange(2, 10)
	// fmt.Println("random number:", randomValue)
	// randomValue = randomWithRange(2, 10)
	// fmt.Println("random number:", randomValue)
	// randomValue = randomWithRange(2, 10)
	// fmt.Println("random number:", randomValue)
	// divideNumber(10, 2)
	// divideNumber(4, 0)
	// divideNumber(8, -4)

	// var diameter float64 = 11
	// var area, keliling = calculate(diameter)

	// fmt.Printf("Luas lingkaran: %2.f\n", area)
	// fmt.Printf("Keliling Lingakaran %2.f\n", keliling)

	/*
		fmt.Sprintf()
		pada dasarnya sama dengan fmt.Printf() ,
		hanya saja fungsi ini tidak menampilkan nilai,
		 melainkan mengembalikan nilainya dalam bentuk string.
	*/
	// var avg = calculateAvarage(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	// var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	// fmt.Println(msg)

	/*
		Pengisian function variadic dapat menggunakan slice
		dengan cara menambahkan titik 3 setalah nama variable (number ...)
	*/
	// var numbers = []int{10, 12, 15, 5, 9, 10, 12}
	// var avg1 = calculateAvarage(numbers...)
	// msg1 := fmt.Sprintf("Rata-rata yang kedua: %2.f", avg1)
	// fmt.Println(msg1)
	var hobbies = []string{"koding", "makan", "tidur"}
	yourHobbies("Rahmatulah Sidik", hobbies...)

	/*
		Closure
		sebuah fungsi yang bisa disimpan dalam variabel
		Closure bisa disimpan dalam variabel.
	*/
	// closure yang menerima parameter slice int dan mengembalikan value 2 int
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e

			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	// %v digunakan untuk menampilkan segala jenis data. Bisa array, int, float, bool, dan lainnya
	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numbers, min, max)

	/*
		Immediately-Invoked Function Expression (IIFE)
		Closure jenis ini dieksekusi langsung pada saat deklarasinya.
		Biasa digunakan untuk membungkus proses yang hanya dilakukan sekali,
		bisa mengembalikan nilai, bisa juga tidak.
	*/
	// mengembalikan slice int
	// Ciri khas IIFE adalah adanya kurung parameter tepat setelah deklarasi closure berakhir.
	//Jika ada parameter, bisa juga dituliskan dalam kurung parameternya.
	var newNumbers = func(min int) []int {
		var r []int //slice int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println("original number :", numbers)
	fmt.Println("filtered number :", newNumbers)

	/*
		cara memanggil function sebagai parameter
	*/
	var data = []string{
		"wick",
		"jason",
		"ethan",
	}

	var dataContains0 = filter(data, func(each string) bool {
		// mencari string yang mengandung karakter o
		return strings.Contains(each, "h")
	})

	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data:", data)
	fmt.Println(dataContains0)
	fmt.Println(dataLength5)

}

func printMessage(message string, names []string) {
	var joinString = strings.Join(names, " ")
	fmt.Println(message, joinString)
}

// function with return value, with same data type
func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

func divideNumber(m, n int) {
	// return untuk menghentikan proses dalam fungsi
	if n == 0 {
		fmt.Printf("invalid divide. %d can not divide by %d.\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d \n", m, n, res)

}

/*
	Multiple Return value
	Note:
	math.Pow(), digunakan untuk memangkat nilai. math.Pow(2, 3) berarti 2 pangkat 3, hasilnya 8.
	math.Pi, adalah konstanta bawaan package math yang merepresentasikan Pi atau 22/7.
*/
func calculate(num float64) (float64, float64) {
	// hitung luas area
	var area = math.Pi * math.Pow(num/2, 2)

	// hitung keliling
	var keliling = math.Pi * num

	return area, keliling
}

/*
	Variadic Functions
	pembuatan fungsi dengan parameter tak terbatas
*/
func calculateAvarage(numbers ...int) float64 {
	var total int = 0

	for _, number := range numbers {
		total += number
	}

	// membagi nilai total dengan panjang jumlah numbers
	var avarage = float64(total) / float64(len(numbers))
	return avarage
}

/*
	Penggabungan parameter biasa/tunggal
	dengan parameter variadic
	parameter variadic harus ditulis pada akhir parameter
*/
// void
func yourHobbies(name string, hobbies ...string) {
	// penggabungan slice menjadi string
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Println("Hello My name is :", name)
	fmt.Println("My Hobby is :", hobbiesAsString)
}

/*
	Penerapan Function sebagai parameter
	func filter(data []string, callback func(string) bool) []string
*/
func filter(data []string, callback func(string) bool) []string {
	var result []string
	// looping data
	for _, each := range data {
		// memanggil callback function
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

/*
	Alias Skema Closure
	type FilterCallback func(string) bool
*/
