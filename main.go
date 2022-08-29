package main

import (
	"fmt"
	"os"
)

func main() {
	// argsWithProg := os.Args
	// argsWithoutProg := os.Args[1:]

	// argument1_int, err := strconv.ParseInt(os.Args[1], 10, 0)
	// fmt.Printf("testt %d", argument1_int)
	// fmt.Println(argument1_int, err, reflect.TypeOf(argument1_int))

	fmt.Println("os.Args[1]", len(os.Args))

	// panic("You can not divide a number by Zero")
	// if err != nil {
	// 	fmt.Println("First input parameter must be integer")
	// 	os.Exit(1)
	// }

	// fmt.Printf("type of a is %T\n", argument1)
	// fmt.Println("semua arg", argsWithProg)
	// fmt.Println("ambil semua arg", argsWithoutProg)
	// fmt.Println("get arg index ke 3", argument1)

	// var selected_index = argument1_int - 1
	// fmt.Println("selected index", selected_index)

	// type Biodata struct {
	// 	nama      string
	// 	alamat    string 2
	// 	pekerjaan string
	// }

	// biodata_list := []Biodata{
	// 	{
	// 		nama:      "jaka",
	// 		alamat:    "perum 1",
	// 		pekerjaan: "developer",
	// 	},
	// 	{
	// 		nama:      "prima",
	// 		alamat:    "perum 2",
	// 		pekerjaan: "business man",
	// 	},
	// }

	// // fungsiTiga := func(list_data []*Manusia) {
	// // 	// cetak data-data dari list
	// // 	for _, data := range list_data {
	// // 		fmt.Println("Nama Manusia: ", data.nama)
	// // 		fmt.Println("Umur: ", data.umur)
	// // 	}
	// // }
	// // fungsiTiga(biodata_list)
	// fmt.Println("Selected", biodata_list[selected_index])

	// // ---------------------------- FUNCTION
	// // sliceTutorial()
	// // loopingTutorial()

	// // // ---------------------------- CLOSURE (anonmous function) & pointer
	// // fmt.Println("-------------------------------- CLOSURE POINTER")
	// // var firstFunction = func(name_list []*string) []*string {
	// // 	result := []*string{}
	// // 	for _, data := range name_list {
	// // 		fmt.Println(data)
	// // 		result = append(result, data)
	// // 	}
	// // 	return result

	// // }

	// // jaka := "jaka"
	// // prima := "prima"
	// // name_list := []*string{&jaka, &prima}
	// // fmt.Println(firstFunction(name_list))

	// // // --------------------------------- STRUCT
	// // fmt.Println("-------------------------------- STRUCT")
	// // type Manusia struct {
	// // 	nama string
	// // 	umur int
	// // }

	// // // var orang1 = Manusia{}
	// // // orang1.nama = "jaka"
	// // // orang1.umur = 10
	// // // var orang2 = Manusia{}
	// // // orang2.nama = "prima"
	// // // orang2.umur = 20

	// // var name_list_struct = []*Manusia{
	// // 	{nama: "jaka", umur: 20},
	// // 	{nama: "prima", umur: 14},
	// // }

	// // fungsiDua := func(list_data []*Manusia) {
	// // 	// cetak data-data dari list
	// // 	for _, data := range list_data {
	// // 		fmt.Println("Nama Manusia: ", data.nama)
	// // 		fmt.Println("Umur: ", data.umur)
	// // 	}
	// // }
	// // fungsiDua(name_list_struct)

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
