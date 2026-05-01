package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Start building. No excuses")

	//Variable + type
	var name string = "Megha"
	age := 24
	isDev := true
	fmt.Printf("My name is %v and my age is %v. And this is %v\n", name, age, isDev)

	//Type
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isDev)

	//Struct
	user := User{Name: "Megha", Age: 24}
	fmt.Println(user)

}
