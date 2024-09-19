package controlStructure

import (
	"fmt"
	"reflect"
)

func loopExample() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// loop through map (map itu object)
	mapData := make(map[string]int)
	mapData["key1"] = 1
	mapData["key2"] = 2

	fmt.Println("typeof", reflect.TypeOf(mapData))

	for key, data := range mapData{
		string_data := fmt.Sprintf("key=%s value=%d", key, data)
		fmt.Println(string_data)
	}

	// loop array (exactly tau isinya berapa example 5)
	// var arrayData [5]int
	// arrayData[4] = 100
	// fmt.Println(arrayData)
	arrayData2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("typeof", reflect.TypeOf(arrayData2))
	for data := range arrayData2 {
		fmt.Println("data", data)
	}
	

	// loop in slice (seperti array tapi tidak tau exacly isi array akan ada berapa length)
	sliceData := []int{1,2,3,4,5}
	fmt.Println("typeof", reflect.TypeOf(sliceData))
	for data := range sliceData {
		fmt.Println("data slice", data)
	}


}
