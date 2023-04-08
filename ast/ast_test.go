package ast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"testing"
)

func TestMyAstAnalyzer(t *testing.T) {
	// 解析源代码文件
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "ast.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	ast.Walk(&visitor{}, f) // 遍历AST树
}
