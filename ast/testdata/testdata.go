package testdata

import "fmt"

var packageIntVar = 1
var packageStrVar = "hello"

func PrintPackageVar() {
	fmt.Println(packageIntVar)
	fmt.Println(packageStrVar)
}
