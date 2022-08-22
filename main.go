package main

import "fmt"

func main() {
	// ----------------------- SLICE
	var students = []string{"Jaka", "Toni", "Hans", "Stevanus", "Edwin Nugroho", "Kevin Hugo", "Rizki Ramadhan", "Fajar Agus Maulana", "Iqbal Hamdani", "Fahmi Tajuddin"}
	_ = students
	for i, data := range students {
		fmt.Println("orang ke: ", i+1)
		fmt.Println("nama: ", data)
	}

	// ------------------------ LOOP, Condition
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " genap")
		} else {
			fmt.Println(i, " ganjil")
		}
	}

}
