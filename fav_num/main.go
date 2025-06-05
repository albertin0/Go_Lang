package main

import "fmt"

func fav_num(name string, num int) {
	fmt.Printf("Hello %s, Your Favourite number is %d.\n", name, num)
}

// The fav_num function takes two parameters: a string 'name' and an integer 'num'.
// It prints a formatted string that includes the user's name and their favorite number.
// This function demonstrates how to define and use functions in Go.
// The function uses the fmt.Printf function to format the output string.

// What are the uses of %s, %d and \n in the printf function?
// In the fmt.Printf function, %s is a placeholder for a string, %d is a placeholder for an integer,
// and \n is used to insert a newline character at the end of the output.
// The %s and %d placeholders are replaced by the values of 'name' and 'num' respectively when the function is called.
// The \n ensures that the output ends with a new line, making it more readable.

// create function fav_num_2 and return a string instead of printing
func fav_num_2(name string, num int) string {
	return fmt.Sprintf("Hello %s, Your Favourite number is %d.", name, num)
}

func main() {
	var name string
	var num int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your favourite number: ")
	fmt.Scanln(&num)

	fav_num(name, num)

	// Call fav_num_2 and print the returned string
	result := fav_num_2(name, num)
	fmt.Println(result)

}

// Output:
// Enter your name: John
// Enter your favourite number: 7
// Hello John, Your Favourite number is 7.
// This program prompts the user for their name and favourite number,
// then prints a message incorporating both inputs.
// It demonstrates basic input handling and function usage in Go.
// The program uses the fmt package for formatted I/O operations.
// The fav_num function takes a name and a number as parameters
// and prints a personalized message.
// The main function handles user input and calls the fav_num function.
// The program is simple and serves as a basic introduction to Go programming.
// It showcases how to define and call functions in Go.
// The program uses the fmt package for formatted output.
// It demonstrates how to read user input using fmt.Scanln.
// The program is structured with a main function and a separate function for the favorite number.
// The program can be extended to include more features, such as error handling.
// The program is a good starting point for beginners learning Go.
// The program can be run in any Go environment.
// The program can be compiled and executed to see the output.
// The program is a simple example of user interaction in Go.
// The program can be modified to include additional functionality.
// The program is a good exercise for understanding basic Go syntax.
// The program can be used to practice input and output in Go.
// The program can be a part of a larger application that requires user input.
// The program can be used to demonstrate function calls in Go.
// The program can be used to illustrate variable declaration and usage in Go.
// The program can be used to teach basic Go programming concepts.
// The program can be used to practice Go's fmt package features.
// The program can be used to understand how to format strings in Go.
// The program can be used to learn about Go's standard library.
