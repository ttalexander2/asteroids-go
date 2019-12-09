# Go (aka Golang)
Cameron Larson & Thomas Alexander

![](https://lh6.googleusercontent.com/jZkASu5zptzg0vuciQzNFuliEeIQwTLrnGX3qW6p08pNbERWZnIBvNB6gX0hPO28o2H31AQILr1lufBQ2ww_2a4u9DxMfxmIsc9G9uPYmYmjaHMtPvS4ryXvK3uZEBkLpY6eXjcl)
## History and Current Status
Development on Go began in 2007 at Google. It was publically announced in 2009 and version 1.0 was released in 2012. The primary drive behind creating Go was to provide a language with which high performance code (in terms of runtime, code creation, and code maintenance efficiency) that could easily be written for multiprocessor, networked machines running large code bases while addressing many criticisms of currently popular object oriented languages.

The current version of Go is 1.13, which was released in September 2019. Currently, Go is ranked in the top 20 most popular programming languages on the TIOBE Index. A translator for the language can be obtained from [https://golang.org/](https://golang.org/).
## Paradigm
Go is a multi-paradigm language. It is both imperative and object oriented. Go satisfies the characteristics of an object-oriented language in a unique way that attempts to solve many criticisms made of more traditional object-oriented languages.
### Imperative:

Go has branching and repeating logic as well as statements so it is an imperative language. Statements are executed in a procedural fashion, and uses variables to abstract data. Go has the ability to function in a sequential, “one instruction at a time” manner. However, through parallel processing, Go also has the ability this “Von Neumann bottleneck.”

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

Go is a strongly typed language where type declarations are required. New types can be made by the programmer using the `type` keyword and first class functions are supported.

## Control Structures

The control structures of Go are syntactically very similar to C or Java. 
#### Conditionals
Control of conditionals are indicated with `if`,  `else if`, and `else`. They can be used like so:

{% highlight Golang %}
if x == 0 {
  //Do something
} else if x == 1 {
  //Do something else
else {
  //Do a different thing
}
{% endhighlight %}

#### Switch
Additionally, Go also features `switch` control statements. A `switch` is used with multiple `case`s and a `default`. Unlike C, a break is not required at the end of each case.

{% highlight Golang %}
switch x {
  case 0: fmt.Printf("%d\n", x)
  case 1: fmt.Printf("%d\n", x)
  case 2: fmt.Printf("%d\n", x)
  case 3: fmt.Printf("%d\n", x)
  default:  fmt.Printf("%d\n", x)
}
{% endhighlight %}

#### Loops
In terms of repetition, loops can be used in a similar fashion to C. Unlike C, the for loop is used for all repetition. The for loop can be used in the 3-component fashion: 

{% highlight Golang %}
for i:=0; i<5; i++ {
  //Do stuff...
}
{% endhighlight %}

However unlike C, Go does not feature a while loop. Instead, the `for` keyword is used with a boolean condition to achieve the same functionality: 

{% highlight Golang %}
for i < 5{
  //Do stuff...
}
{% endhighlight %}

Additionally, the for keyword can be used with no arguments to create an infinite loop:

{% highlight Golang %}
for {
  //Do stuff...
}
{% endhighlight %}

Like C, the `continue` statement can be used to immediately being the next iteration of the loop, and the `break` statement can be used to immediately exit the loop. In order to iterate over elements in data structures like arrays or maps, a for-each range loop can be used:

{% highlight Golang %}
arr := []string{"i'm", "an", "array"}
for i, s := range arr{
  fmt.Printf("Index: %d, Element: %s", i, s)
}
{% endhighlight %}


This allows the programmer to access both the index `i` and the element `s` in the collection.
Unlike C, do-while loops are not features of Go. An infinite loop and a conditional break statement would need to be used to replicate the do-while behavior:

{% highlight Golang %}
for{
 // Do stuff...
  if(condition){
    break
  }
}
{% endhighlight %}
Go also features `goto` and labels, however their can make for rather messy code.

## Semantics

Go is statically scoped. There is automated garbage collection. Storage is stack-dynamic and heap-dynamic. Each `goroutine` thread (including the main thread) created by the programmer has its own stack. This encourages threads to communicate through synchronized channels instead of using shared memory which reduces race conditions. Go supports constants of type char, string, boolean, and numeric values.

## Desirable Language Characteristics

#### Efficiency:
Go is a language designed to be efficient. First off, Go is compiled language. Unlike languages such as Java, which is first compiled into byte-code and then read on a VM, Go is compiled directly from source code to a binary executable. This allows for quick and direct translation, and fast build speeds. In terms of execution, Go is quite efficient. Go was written to make concurrency simple for the programmer. Unlike programming languages that were originally developed to be single threaded, Go was developed from the ground up to take advantage of multi-core processors. Instead of using threads, Go has what it calls `goroutines`, as described [above](#semantics). This allows for a very efficient use of multi-core processing power, resulting in fast concurrent execution. Furthermore, Go features clean, easy to use syntax. This makes it very easy for a programmer to translate ideas into efficient code. Because of this clean and neat syntax, Go programs are very easy to maintain. Due to its simplistic approach to OOP, code is cleaner, allowing for easier maintenance—especially with larger applications.

#### Regularity:
One feature of Go that demonstrates regularity is the repetition control structure. Unlike C, instead of for, while, and do-while, Go solely uses `for` for all types of repetition. Another example is the variable naming conventions enforced by the compiler. The capital letter of the variable determines the visibility of the variable. If the variable starts with a capital letter, the variable will be visible to other files.

#### Security/Reliability
Go is statically typed, which allows the compiler to enforce type usage, minimizing programmer type errors. Go has error handling with the build-in error type. Functions with the potential to fail can return an `error` type in addition to its regular return value:

{% highlight Golang %}
f, err := os.Open("file.txt")
if err != nil {
    panic(err)
}
{% endhighlight %}

This error type is only returned if an error occurs, and can describe the error that occurs as a string. The programmer can use this returned error type to conditionally handle the error. This allows the programmer to defend against runtime errors. Go is also quite secure in how it handles concurrency. Go makes it easy to write clean, concurrent code, and encourages communication between goroutines via synchronized channels rather than shared memory. This helps reduce unsafe code that could result in race conditions.

#### Extensibility:
Like most OO languages, Go allows programmers to add new classes and types. However, unlike Java, Go does this with structs and interfaces. Instead of methods being included within an object, methods in Go can be defined on struct types. The methods can then be used in a similar fashion to class methods in Java. Instead of constructors, Go allows the programmer to implement a `New()` method.
  

## Support for Data Abstractions
An example of data abstractions in Go is the `interface`. Similar to java, Go allows the programmer to define methods. Interface types can invoke the specific type’s implementation. Unlike Java, these interfaces are implemented implicitly, meaning that the compiler is the only enforcement of the interface to the type.

## Syntax
One appealing syntax choice (for the most part) in Go is keeping the familiar C style syntax. 
The syntax choice I’d like to see changed is the variable declarations since they are backwards from the C style declarations. In Go the variable name in stated first and then the type is listed. Additionally, outside functions and methods, the longhand declaration style must be used which starts with the keyword `var` (ex: `var x int`). This is less efficient from a code writing perspective than having the type mark the statement as a variable declaration. To be fair, this choice did enable the parser to be implemented with single token look-ahead, however, which increases the compile speed.

## Works Cited

[https://golang.org/doc/effective_go.html](https://golang.org/doc/effective_go.html)

[https://www.golang-book.com/books/intro/5](https://www.golang-book.com/books/intro/5)

[https://blog.golang.org/gos-declaration-syntax](https://blog.golang.org/gos-declaration-syntax)

[https://golang.org/doc/](https://golang.org/doc/)

[https://medium.com/@kevalpatel2106/why-should-you-learn-go-f607681fad65](https://medium.com/@kevalpatel2106/why-should-you-learn-go-f607681fad65)

[http://golangtutorials.blogspot.com/2011/06/goroutines.html](http://golangtutorials.blogspot.com/2011/06/goroutines.html)

[https://stackimpact.com/docs/go-performance-tuning/](https://stackimpact.com/docs/go-performance-tuning/)

[https://www.golang-book.com/books/intro/10](https://www.golang-book.com/books/intro/10)

[https://www.golang-book.com/books/intro/9](https://www.golang-book.com/books/intro/9)


## Super Awesome Space Game!
Super Awesome Space Game! is a not-so-faithful recreation of the classic game Asteroids. The game allows you to fly around in a spaceship and shoot asteroids. The game is won when all the asteroids on the screen are gone.This game was written in GoLang utilizing a port of OpenGl (Open Graphics Library). OpenGl is a low level, open source graphics api common among many programming languages. Using this interface, all objects were created by defining vertices as arrays of floats and converting them into Vertex Array Objects. In this format, OpenGl then renders each object in line mode, by drawing lines between the coordinates. Everything from translation, rotation, asteroid generation, laser angle calculation, and other functions were all done in terms of these vertex array objects.

	The asteroids in Super Awesome Space Game! are somewhat randomly generated by creating points at a random interval in a circle around the origin. The distance from the origin for each of these points is also somewhat randomized, generating asteroids of various shapes and sizes. Additionally, the asteroid class also features a split function, which is called upon collision with a laser, resulting in the creation of 2 new, smaller asteroids. These asteroids are also randomly generated, but are always smaller and faster than the original asteroid. Collision with an asteroid results in a death animation, and game over. The game can be seen in action here:

![Super Awesome Space Game!](https://raw.githubusercontent.com/ttalexander2/csc372-final-project/master/docs/2019-12-09 16-00-21.gif)

#### Running the program:
The repository can be found here: https://github.com/ttalexander2/csc372-final-project

This game has 2 dependencies. OpenGl and GLFW must be installed for the program to run. The following commands can be used:
{% highlight Shell %}
go get 	"github.com/go-gl/gl/v4.1-core/gl"
go get 	"github.com/go-gl/glfw/v3.2/glfw"
{% endhighlight %}
Cd to the directory with the files, then use the following to run:
{% highlight Shell %}
go run *.go
{% endhighlight %}
