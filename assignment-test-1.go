// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	// ---------------------------------------- VALIDATE

// 	if len(os.Args) < 2 {
// 		panic("must have args input from user")
// 	}

// 	// fmt.Println("reflect", reflect.TypeOf(os.Args[1]))
// 	argument1_int, err := strconv.ParseInt(os.Args[1], 10, 0)
// 	// fmt.Println("convert", argument1_int)
// 	// fmt.Println("err", err)

// 	if err != nil {
// 		panic("First input parameter must be integer")
// 	}

// 	if argument1_int < 0 {
// 		panic("Absen dimulai dari 1")
// 	}

// 	var selected_index int = int(argument1_int) - 1

// 	type Biodata struct {
// 		nama      string
// 		alamat    string
// 		pekerjaan string
// 		alasan    string
// 	}

// 	biodata_list := []Biodata{
// 		{
// 			nama:      "jaka",
// 			alamat:    "perum 1",
// 			pekerjaan: "developer",
// 			alasan:    "golang keren kata jaka",
// 		},
// 		{
// 			nama:      "toni",
// 			alamat:    "perum 2",
// 			pekerjaan: "business man",
// 			alasan:    "golang keren kata toni",
// 		},
// 		{
// 			nama:      "hans",
// 			alamat:    "perum 2",
// 			pekerjaan: "business man",
// 			alasan:    "golang keren kata hans",
// 		},
// 		{
// 			nama:      "stevanus",
// 			alamat:    "perum 2",
// 			pekerjaan: "business man",
// 			alasan:    "golang keren kata stevanus",
// 		},
// 		{
// 			nama:      "edwin nugroho",
// 			alamat:    "perum 2",
// 			pekerjaan: "business man",
// 			alasan:    "golang keren kata edwin nugroho",
// 		},
// 		{
// 			nama:      "kevin hugo",
// 			alamat:    "perum 2",
// 			pekerjaan: "business man",
// 			alasan:    "golang keren kata kevin hugo",
// 		},
// 		{
// 			nama:      "rizki ramdhan",
// 			alamat:    "perum 2",
// 			pekerjaan: "business man",
// 			alasan:    "golang keren kata rizki ramdhan",
// 		},
// 		{
// 			nama:      "fajar agus maulana",
// 			alamat:    "perum 2",
// 			pekerjaan: "business man",
// 			alasan:    "golang keren kata fajar agus maulana",
// 		},
// 		{
// 			nama:      "iqbal hamdani",
// 			alamat:    "perum 2",
// 			pekerjaan: "business man",
// 			alasan:    "golang keren kata iqbal hamdani",
// 		},
// 		{
// 			nama:      "fahmi tajuddin",
// 			alamat:    "perum 2",
// 			pekerjaan: "business man",
// 			alasan:    "golang keren kata fahmi tajuddin",
// 		},
// 	}

// 	fungsiTiga := func(list_data []Biodata) {
// 		// cetak data-data dari list
// 		for i, data := range list_data {
// 			if i == selected_index {
// 				fmt.Printf("Nama : %s\n", data.nama)
// 				fmt.Printf("Alamat : %s\n", data.alamat)
// 				fmt.Printf("Pekerjaan : %s\n", data.pekerjaan)
// 				fmt.Printf("Alasan Kursus : %s\n", data.alasan)
// 			}
// 		}
// 	}
// 	fungsiTiga(biodata_list)

// }
