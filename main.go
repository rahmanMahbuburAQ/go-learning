package main

import "fmt"

func main() {

	//strings
	var nameOne string = "Banyan Tree"
	var nameTwo = "Coconut Tress"
	fmt.Println(nameOne, nameTwo)

	nameOne = "Hello"
	nameTwo = "HI"
	fmt.Println(nameOne, nameTwo)
	//inside of function string can be declared like this:
	nameFour := "Nakatani"

	fmt.Println(nameFour)

	//ints
	var ageOne int = 30
	var ageTwo = 40
	ageThree := 50

	fmt.Println(ageOne, ageTwo, ageThree)

	//https://go.dev/ref/spec#Numeric_types
	//bits and memory
	var ageFive int8 = 113
	var ageSix uint = 255
	fmt.Println(ageFive, ageSix)

	//float
	var scoreOne float32 = 2.73
	var scoreTwo float64 = 7878.99
	scoreThree := 787.98
	fmt.Println(scoreOne, scoreTwo, scoreThree)

	//printing and formatting strings
	fmt.Print("Hello, ")
	fmt.Print("world! \n")
	fmt.Print("New line after Hello world! \n")

	//printing and formatting strings(new line automatically)
	fmt.Println("Hello2, ")
	fmt.Println("world2!")
	fmt.Println("New line after Hello world2! ")
	fmt.Println("Total amount is", scoreOne)

	age := "49"
	name := "Fuad"

	//Printf: formatting strings:
	fmt.Printf("I am  %v and my name is %v \n", age, name)
	fmt.Printf("I am  %q and my name is %q", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("You scored %f points! \n", 228.87)
	fmt.Printf("You scored %0.2f points! \n", 228.7887)

	//Sprintf: save formatted strings:
	var str = fmt.Sprintf("My age is %v and my name is %v", age, name)
	fmt.Println("The saved string is:", str)

}
