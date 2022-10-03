# GOLANG notes

These are my own notes that I am taking as I go through the Golang Tutorial featured at https://go.dev/doc/tutorial/. 

## Modules & Packages

Golang programs are organized into modules and packages. Packages are simply collections of functions. Modules are collections of packages; they also specify the dependencies needed for your code to run.

## Creating a simple Go program

First, you start by creating a directory for your module and moving into it. Let's start with `hello`:
```bash
mkdir hello && cd hello
```

Next, you initialize your module using the `go mod init` command and passing in a module name using the `<prefix>/<descriptive-text>` structure ([See this link for more info](https://go.dev/doc/modules/managing-dependencies#naming_module)):

```bash
go mod init example.com/hello
```

Now that you have a directory to house your module and you have initialized it, you can start writing code in a `<filename>.go` file. We can create a simple "Hello, World!" example like so:
```golang
// Declare a main package to collect related functions
package main

// Import the `fmt` library to print text to the stdout
import "fmt"

// Declare a function that prints out "Hello, World!"
func main() {
    fmt.Println("Hello, World!")
}
```

Finally, after writing the program, we can run it using the following command (in the `hello` directory that we created):
```bash
go run .
```

The result should look like:
```bash
> go run .
Hello, World!
```

Note: In order to execute a Go program, a module must have a `main` package
and function that executes the code. In this example, our `main` package and function is in the `hello.go` file in our `hello` directory.

## Adding Dependencies & Dependency Tracking with `go.mod` file

Go uses the `go.mod` file inside of your module to keep track of package dependencies. Whenever you add new dependencies to a package or file, run the following command inside of your module directory to add that dependency to your `go.mod` file:
```bash
go mod tidy
```

This will add the dependencies to your `go.mod` file and allow you to run your code.

## Finding Available Go Packages

Golang packages are available [here](https://pkg.go.dev/search?q=quote).

## Using code from local packages

For production, it is common practice to publish your modules to a repository where Go Tools can go find and download them. However, for local development, we need to make a slight modification to our `go.mod` file to import local modules into other modules. 

The following command can be used to make a local module available for use in another local module on the same system:
```bash
go mod edit -replace <module-name>=/path/to/module/<module-name> # relative paths are OK here, too
# Make sure the file is updated
go mod tidy
```

## Error Handling

It is common pattern to return an error value from a function to another calling function. The function that is called may look like so:

```golang
// The code for the function that is being called
import (
    "fmt"
    "errors"
    "log"
)

func my_func(name string) (string, error){
    if name == "" {
        return "", errors.New("missing name")
    }

    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    // We return our message and nil in place of our error
    // to indicate that the function ran successfully
    return message, nil
}

// The code for the calling function
func main() {
    // Set up log
    log.SetPrefix("my_func: ")
    log.SetFlags(0)
    // We return our message and nil in place of our error
    // to indicate that the function ran successfully
    message, error := my_func("Martin")
    // Check for errors and stop execution & log the error 
    // if there is one
    if error != nil {
        log.Fatal(error)
    }
    // There are no errors, so we can print the message
    fmt.Println(message)
}
```

## Go Slices

Slices are a data structure similar to that of a list in Python. It is an array-like structure that can dynamically expand or contract in size. The syntax for creating one is below:

```golang
// A sample Slice
myslice := [] string {
    "I am the first entry",
    "I am the second entry",
}

// In a more general sense...
myslice := [] <datatype> {
    entries...
}
```

Accessing an index of a slice uses common syntax as well:
```golang
fmt.Println(myslice[0])
// prints "I am the first entry"
```

## Go Maps

Maps in Go are similar to other map structures in other languages. The syntax for creating a map is like so:

```golang
mymap := make(map[string]string)

// General structure
mymap := make(map[<key-data-type>]<value-data-type>)
```

Accessing values in a map is done like this:
```golang
// Adding values to the map
mymap["key"] = "first value"
mymap["key2"] = "second value"

// Checking for values in a map using a key.
// We can use the `missing` variable to handle 
// cases where a key is not in the map or is 0 or ""
missing, val := mymap["key2"]

// Creating a new map
newmap := map[string]int{
    "foo":1,
    "bar":2
}
```

More examples and a further explanation can be found [here](https://gobyexample.com/maps).

## Go Range

`range` is used like a `for` loop to iterate over elements in data structures. It returns two variables for each loop - the index number and the value at the given index. 

Here are some syntax examples over different data structures