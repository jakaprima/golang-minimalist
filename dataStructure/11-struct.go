package dataStructure

import (
	"fmt"
)


type ClassName struct {
	attribute1 string
	attribute2 string
}

func NewClassName(attribute1Value string, attribute2Value string) (*ClassName) {
	newClassName := &ClassName{
		attribute1: attribute1Value,
		attribute2: attribute2Value,
	}
	return newClassName
}

func (n *ClassName) MethodDariClassName() string {
	return "halo "+n.attribute1
}

func RunStructExample(){
	fmt.Println("run struct example")

	type UserStruct struct {
		Name string
		Age int
		Address string
	}

	// non pointer value type
	userInstance := UserStruct{
		Name: "jaka",
		Age: 1,
		Address: "aaa",
	}

	fmt.Println(userInstance)

	// struct in pointer pointer (reference type)
	userPointerInstance := &UserStruct{
		Name: "jaka",
		Age: 1,
		Address: "aaa",
	} 
	fmt.Println("userPointerInstance", userPointerInstance)

	// Perbedaan ketika diubah
	userInstance2 := userInstance
	userInstance2.Age = 2
	fmt.Println(userInstance.Age) // Output: {jaka 1 aaa}
	fmt.Println(userInstance2.Age) // Output: {jaka 2 aaa}

	userPointerInstance2 := userPointerInstance
	(*userPointerInstance2).Age = 3 // Atau bisa ditulis userPointerInstance.Age = 3
	fmt.Println(userPointerInstance.Age)
	fmt.Println(userPointerInstance2.Age) // Output: {jaka 3 aaa}

	// struct and method struct (sama seperti class  dan method class)
	classNamePointerInstance := NewClassName("isi attribute 1", "isi attribute 2")
	fmt.Println("classNameInstance.attribute1", classNamePointerInstance.attribute1)
	fmt.Println("classNameInstance.attribute2", classNamePointerInstance.attribute2)
	fmt.Println("panggil method", classNamePointerInstance.MethodDariClassName())

	


}