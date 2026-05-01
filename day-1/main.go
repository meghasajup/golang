package main

import "fmt"

type User struct {
	Name string
	Age  int
}

//Function
func add(a int, b int) int {
	return a + b
}

//Func with no return
func greet(name string) {
	fmt.Println("Hello", name)
}

//Multiple return values (VERY important in Go)
func divide(c int, d int) (int, error) {
	if d == 0 {
		return 0, fmt.Errorf("Couldn't divide by zero")
	}
	return 0, nil
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

	//If
	if age > 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	//Loop
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	//Function
	result := add(3, 3)
	fmt.Println(result)

	//Multiple return values (VERY important in Go)
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
