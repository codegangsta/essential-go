// There are 2 types of programs you can create with Go. The first is a
// Command, which can be invoked via a command line. The second type of program
// you can create in Go is a Library. In this example we will focus on the former.
//
// For Go to recognize that this program is a Command, we define our code in
// the 'main' package.
package main

// Go is also a very modular programming language. In fact there is very little
// you can do without importing a package to give you more functionality. We
// want to print a "Hello world" out to our terminal, so we will import the
// ubiquitous "fmt" package to help us do our printing.
import "fmt"

// Every Command type program in Go requires you to have a 'main' function
// inside the 'main' package. This will be your entrypoint into the application
// and where you begin all of your programs work.
//
// The main function takes no arguments and has nothing to return. This makes
// it really easy to get a program up and running because you won't have to
// remember all of the bare minimum boilerplate a program needs to compile
func main() {
	// Because we imported the "fmt" package, we can now call functions that are
	// exported from that package. We can do this by identifying the package name
	// and invoking the function. The important thing to know is that Packages
	// are nonintrusive in Go. For the most part, simply importing a package will
	// not mess with your program or environment. You have to take action and
	// call a function on a package for it to really provide value to your
	// program.
	//
	// We should not be able to run our program and see that it outputs "Hello world!"
	fmt.Println("Hello", "world!")
}
