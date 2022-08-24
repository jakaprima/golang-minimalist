package main

import "fmt"

func main() {
	// ---------------------------- FUNCTION
	sliceTutorial()
	loopingTutorial()

	// ---------------------------- CLOSURE (anonmous function) & pointer
	fmt.Println("-------------------------------- CLOSURE POINTER")
	var firstFunction = func(name_list []*string) []*string {
		result := []*string{}
		for _, data := range name_list {
			fmt.Println(data)
			result = append(result, data)
		}
		return result

	}

	jaka := "jaka"
	prima := "prima"
	name_list := []*string{&jaka, &prima}
	fmt.Println(firstFunction(name_list))

	// --------------------------------- STRUCT
	fmt.Println("-------------------------------- STRUCT")
	type Manusia struct {
		nama string
		umur int
	}

	// var orang1 = Manusia{}
	// orang1.nama = "jaka"
	// orang1.umur = 10
	// var orang2 = Manusia{}
	// orang2.nama = "prima"
	// orang2.umur = 20

	var name_list_struct = []Manusia{
		{nama: "jaka", umur: 20},
		{nama: "prima", umur: 14},
	}

	fungsiDua := func(list_data []Manusia) {
		// cetak data-data dari list
		for _, data := range list_data {
			fmt.Println("Nama Manusia: ", data.nama)
			fmt.Println("Umur: ", data.umur)
		}
	}
	fungsiDua(name_list_struct)

}

func sliceTutorial() {
	fmt.Println("-------------------------------- SLICE")
	// ----------------------- SLICE
	var students = []string{"Jaka", "Toni", "Hans", "Stevanus", "Edwin Nugroho", "Kevin Hugo", "Rizki Ramadhan", "Fajar Agus Maulana", "Iqbal Hamdani", "Fahmi Tajuddin"}
	_ = students
	for i, data := range students {
		fmt.Println("orang ke: ", i+1)
		fmt.Println("nama: ", data)
	}
}

func loopingTutorial() {
	fmt.Println("-------------------------------- LOOP, CONDITION")
	// ------------------------ LOOP, Condition
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " genap")
		} else {
			fmt.Println(i, " ganjil")
		}
	}
}
