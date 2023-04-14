package testdata

import (
	"fmt"
	"math"
	"strings"
)

var packageVar = 1

// TestFunction1 this is test function no.1
// return A int number
func TestFunction1() int {
	return 1
}

// TestForMyAstAnalyzer the testdata main function
func TestForMyAstAnalyzer(i float64) float64 {
	var double = 2
	i = i * float64(double)

	var list = []string{
		"hello",
		"world",
	}
	fmt.Println(strings.Join(list, ","))

	fmt.Println(packageVar)

	return math.Abs(i)
}
