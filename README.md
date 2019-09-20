# Welcome to this Repository!
This is a small project of [Jeremytjuh](https://github.com/Jeremytjuh "Jeremytjuh's profile") and [Tristangoossens](https://github.com/tristangoossens "Tristangoossens profile")
In this project we wanted to automatically generate the barebone test for a package

## Why?
We thought it would be an interesting and perhaps useful project to make!

## How to use our project??

To start off, you will need a package. This package can contain any functions which are exporteable(function names should start with a capital letter).

### ***Examples***

**Correct example**

```go
package example

import (
    "fmt"
)

func Test() {
    fmt.Println(1+1)
}
```

**Incorrect example**

```go
package main // Package cannot be main

import (
    "fmt"
)

func test() { // This function is not exportable
    fmt.Println(1+1)
}
```

***

Then make sure you have our generator.go file located in the [generator directory](https://github.com/tristangoossens/testfilegenerator/tree/master/generator) inside of your package directory.

After importing this project into your files open generator.go, scroll down into func main and enter your package file into the parentheses.

![alt text](https://github.com/tristangoossens/testfilegenerator/blob/master/readme-images/enterfilename.png "Enter file name")

Then once you are ready, run the program!

```bash
go run Generator.go
```
