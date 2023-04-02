package ast

import (
	"fmt"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestMyAstAnalyzer(t *testing.T) {
	// 使用analysistest.TestData()会返回当前目录下的testdata全局路径(比如xxx/mygolint/ast/testdata), 作为测试数据的目录
	var testdataDir = analysistest.TestData()
	fmt.Println(testdataDir)

	// 第3个参数传递自定义的Analyzer
	analysistest.Run(t, testdataDir, MyAstAnalyzer)
}
