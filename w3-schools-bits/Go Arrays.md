Go Arrays
Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.

Declare an Array
In Go, there are two ways to declare an array:

An array in Go is a fixed-size box that holds many items of the same type.
You cannot change its size later.
Simple example:
prices := [3]int{10, 20, 30}

Use [index] to read or change items (starts from 0)
Use len(prices) to know the size
If you don’t fill all spots, empty ones become 0