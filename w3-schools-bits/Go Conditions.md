Go Conditions help the computer make decisions (true or false).

Use if to run code only when something is true.
Example:Goif x > 10 {
    // do this
}
Add else for when it is false.
Add else if to check another condition.
Use switch when you have many possible values (like choosing between red, blue, green).

Simple example with if:
Goage := 20
if age >= 18 {
    fmt.Println("You are an adult")
} else {
    fmt.Println("You are young")
}