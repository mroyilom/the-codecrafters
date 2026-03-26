## Go Output Functions
The print function is the final out put of the go program that tells it to print the program.
Go has three functions to output text:
Print()
Println()
Printf()
The Print() function prints its arguments with their default format.
var i,j string = "Hello","World"

  fmt.Print(i)
  fmt.Print(j)

HelloWorld as the result
If we want to print the arguments in new lines, we need to use \n.
fmt.Print(i, "\n")
  fmt.Print(j, "\n")
Result:
Hello
World


It is also possible to only use one Print() for printing multiple variables.
var i,j string = "Hello","World"

  fmt.Print(i, "\n",j)
If we want to add a space between string arguments, we need to use " ":
func main() {
  var i,j string = "Hello","World"

  fmt.Print(i, " ", j)

The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end:
func main() {
  var i,j string = "Hello","World"
  fmt.Println(i,j)
The Printf() function first formats its argument based on the given formatting verb and then prints them.
Here we will use two formatting verbs:
%v is used to print the value of the arguments
%T is used to print the type of the arguments
var i string = "Hello"
  var j int = 15

  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T", j, j)
i has value: Hello and type: string
j has value: 15 and type: int
