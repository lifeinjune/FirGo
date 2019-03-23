//First go program helloworld
//On command line 'go run "filename"' to run the go program
/* 'go bulid' will give the executable file which me only compile
the source code*/
// 'go fmt' is used to format the source code
// 'go install' compiles and install the package
// 'go get' download source code of someone else package
// 'go test' used to run the test file
package main

/* what is 'package'
-package is like project
-package can contain many files
-go files must declare which package it belongs to which this case
'package main' this file belong to main package

which call package main?
-there are two types of packages in Go
-One is executable and another one is reusable
-executable package is the one that will run and produce executable file
-Reusable is like libraries that some reusable codes

How to figure out which package type is
-name of the package will determine type of the package
-package name is 'main' is executable otherwise reusable
-reusable file will not produce executable file git build command
*/

import "fmt"

/*
-import statment use to access code inside of another package
-'fmt' is standard library used to print out many things
-think it as link other packages by using import statement
-information available on golang website for standard packages
*/

func main() { //-package main must contain 'main' function
	// func: declare statment is function
	// main: name of function
	// (): list of arguments passing
	fmt.Println("hi")
}

/*
Go file pettarn
1. package declaration
2. import packages need
3. Declare function
*/

//change the directory to GOPATH
