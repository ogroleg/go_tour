# go_tour


### Hello World

Traditionally the first program you write in any programming language is called a “Hello World” program – a program that simply outputs Hello World to your terminal.

```
package main

import "fmt"

// this is a comment

func main() {
    fmt.Println("Hello World")
}
```
- The import keyword is how we include code from other packages to use with our program.
- The fmt package (shorthand for format) implements formatting for input and output. 
```
fmt.Println("Hello World")
```
- This statement is made of three components. First we access another function inside of the fmt package called Println (that's the fmt.Println piece, Println means Print Line). Then we create a new string that contains Hello World and invoke (also known as call or execute) that function with the string as the first and only argument.

### Variables:

 - Names must start with a letter and may contain letters, numbers or the _ (underscore) symbol.
 
 - A special way to represent multiple words in a variable name known as lower camel case (also know as mixed case, bumpy caps, camel back or hump back). The first letter of the first word is lowercase, the first letter of the subsequent words is uppercase and all the other letters are lowercase.
 
 #### Scope
 
 The range of places where you are allowed to use x is called the scope of the variable. 
 
 ### Constant:
 
 Constants are basically variables whose values cannot be changed later. They are created in the same way you create variables but instead of using the var keyword we use the ```const``` keyword.
 
