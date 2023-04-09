package testdata

import "fmt"

var packageIntVar = 1
var packageStrVar = "hello"

// PrintPackageVar 打印包变量
func PrintPackageVar() {
	fmt.Println(packageIntVar)
	fmt.Println(packageStrVar)
}
