package main

import "fmt"

func main() {
	// ---------------------------- FUNCTION
	sliceTutorial()
	loopingTutorial()

	// ---------------------------- CLOSURE (anonmous function) & pointer
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

}

func sliceTutorial() {
	// ----------------------- SLICE
	var students = []string{"Jaka", "Toni", "Hans", "Stevanus", "Edwin Nugroho", "Kevin Hugo", "Rizki Ramadhan", "Fajar Agus Maulana", "Iqbal Hamdani", "Fahmi Tajuddin"}
	_ = students
	for i, data := range students {
		fmt.Println("orang ke: ", i+1)
		fmt.Println("nama: ", data)
	}
}

func loopingTutorial() {
	// ------------------------ LOOP, Condition
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " genap")
		} else {
			fmt.Println(i, " ganjil")
		}
	}
}
