package analyzer

import (
	"fmt"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestDemoAnalyzer(t *testing.T) {
	// 构建测试对象
	// 使用analysistest.TestData()会返回当前目录下的testdata全局路径(比如xxx/mygolint/ast/testdata), 作为测试数据的目录
	var testdataDir = analysistest.TestData()
	fmt.Println(testdataDir)

	// analysistest.Run函数是用于运行分析器测试的帮助函数，它会自动构建测试用例并运行分析器，并检查分析器输出是否符合预期
	// 在编写分析器测试时，我们可以使用 analysistest.Run 函数来简化测试用例的构建和执行过程
	// 运行分析器, 第3个参数传递自定义的Analyzer
	analysistest.Run(t, testdataDir, DemoAnalyzer)
}
