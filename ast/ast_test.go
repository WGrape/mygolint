package ast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"testing"
)

func TestMyAstAnalyzer(t *testing.T) {
	var err error

	// 解析源代码文件
	fset := token.NewFileSet()
	if GlobalF, err = parser.ParseFile(fset, "./testdata/testdata.go", nil, parser.ParseComments); err != nil {
		log.Fatal(err)
		return
	}

	ast.Walk(&visitor{}, GlobalF) // 遍历AST树
}
