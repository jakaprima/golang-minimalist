package functionPackage

import (
	"fmt"
)

func function1(string) (string, error){
	// if true {
	// 	return "", fmt.Errorf("error triggered")
	// }
	return "return data string 1", nil
}

func FunctionExample(){
	executeReturn, error := function1("test params string")
	fmt.Println("execute return", executeReturn)
	fmt.Println("error", error)
}