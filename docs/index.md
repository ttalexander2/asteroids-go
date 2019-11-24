# Go (aka Golang)
![](https://lh6.googleusercontent.com/jZkASu5zptzg0vuciQzNFuliEeIQwTLrnGX3qW6p08pNbERWZnIBvNB6gX0hPO28o2H31AQILr1lufBQ2ww_2a4u9DxMfxmIsc9G9uPYmYmjaHMtPvS4ryXvK3uZEBkLpY6eXjcl)
## History and Current Status
Development on Go began in 2007 at Google. It was publically announced in 2009 and version 1.0 was released in 2012. The primary drive behind creating Go was to provide a language with which high performance code (in terms of runtime, code creation, and code maintenance efficiency) that could easily be written for multiprocessor, networked machines running large code bases while addressing many criticisms of currently popular object oriented languages.

The current version of Go is 1.13, which was released in September 2019. Currently, Go is ranked in the top 20 most popular programming languages on the TIOBE Index. A translator for the language can be obtained from [https://golang.org/](https://golang.org/).
## Paradigm
Go is a multi-paradigm language. It is both imperative and object oriented. Go satisfies the characteristics of an object-oriented language in a unique way that attempts to solve many criticisms made of more traditional object-oriented languages.
### Imperative:

Go has branching and repeating logic as well as statements so it is an imperative language.

### Object Oriented:

##### *Objects and Classes*:

Go has the capability of creating objects and classes but does not assume that everything the programmer creates is an object. This prevents “objects” that contain only static methods and attributes, such as the Math class in Java. Since all of this class’ methods and attributes are static there is unnecessary code complexity in needing to declare every method and attribute as static so that it can be used without instantiating a Math class. Go removes this extraneous code by allowing C style imperative code to exist outside of a class declaration.

  

##### *Abstraction:*

Go contains data and control abstractions as well as class and object abstractions.

  

##### *Encapsulation:*

Go allows for encapsulation of methods, attributes, and inner classes by considering all declarations of such that start with a lowercase letter to be private. Private methods, attributes, and classes are not accessible outside of that file.

  

##### *Inheritance:*

Go departs from the traditional object-oriented programming language’s view of inheritance by only supporting composition and interfaces.This is beneficial because it avoids “Inheritance Hell” often seen in the project planning phase of other OOP, such as Java, while still providing a way to describe a relationship between two different pieces of code.

  

##### *Polymorphism:*

The implementation of interfaces in Go support a “runtime polymorphism”. Interfaces do not need to be declared in the code that implements them. Instead the code is statically checked by the Go compiler, if the code has implemented the features specified by the interface then it is considered to be valid for that interface.

## Typing System

Go is a strongly typed language where type declarations are required. New types can be made by the programmer using the type keyword and first class functions are supported.

## Control Structures

The control structures of Go are syntactically very similar to C or Java. 
#### Conditionals
Control of conditionals are indicated with `if`,  `else if`, and `else`. They can be used like so:
```Go
if x == 0 {
  fmt.Printf("%d\n", x)
} else if x == 1 {
  fmd.Printf("%d\n", x}
else {
  fmd.Printf("%d\n\n", x)
}
```
#### Switch
Additionally, Go also features `switch` control statements. A `switch` is used with multiple `case`s and a `default`. Unlike C, a break is not required at the end of each case.
```Go
switch x {
  case 0: fmt.Printf("%d\n", x)
  case 1: fmt.Printf("%d\n", x)
  case 2: fmt.Printf("%d\n", x)
  case 3: fmt.Printf("%d\n", x)
  default:  fmt.Printf("%d\n", x)
}
```
In terms of repetition, loops can be used in a similar fashion to C. Unlike C, the for loop is used for all repetition. The for loop can be used in the 3-component fashion: for i:=0; i<5; i++ However unlike C, Go does not feature a while loop. Instead, the for keyword is used with a boolean condition to achieve the same functionality: for (i < 5). Additionally, the for keyword can be used with no arguments to create an infinite loop. Like C, the continue statement can be used to immediately being the next iteration of the loop, and the break statement can be used to immediately exit the loop. In order to iterate over elements in data structures like arrays or maps, a for-each range loop can be used:

for i, s := range array. This allows the programmer to access both the index(i) and the element(s) in the collection. Unlike C, do-while loops are not features of Go. An infinite loop and a conditional break statement would need to be used to replicate the do-while behavior. Go also features goto and labels, however using these to control

## Semantics

Go is statically scoped. There is automated garbage collection. Storage is stack-dynamic and heap-dynamic. Each goroutine thread (including the main thread) created by the programmer has its own stack. This encourages threads to communicate through synchronized channels instead of using shared memory which reduces race conditions. Go supports constants of type char, string, boolean, and numeric values.

## Desirable Language Characteristics

  

## Support for Data Abstractions

## Syntax

