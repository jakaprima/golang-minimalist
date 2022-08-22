package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "adalah Angka genap")
		} else {
			fmt.Println(i, "adalah Angka ganjil")
		}

	}
}
