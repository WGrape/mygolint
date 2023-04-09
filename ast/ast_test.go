package ast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
	"testing"
)

func TestMyAstAnalyzer(t *testing.T) {
	var err error

	// 解析源代码文件
	GlobalFset = token.NewFileSet()
	if GlobalF, err = parser.ParseFile(GlobalFset, "./testdata/testdata.go", nil, parser.ParseComments); err != nil {
		log.Fatal(err)
		return
	}

	// 遍历文件中的所有函数声明
	for _, decl := range GlobalF.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			// 检查函数声明是否有注释
			if fn.Doc == nil || len(fn.Doc.List) == 0 {
				fmt.Printf("Warning: Function %s in %s:%d has no comment\n", fn.Name.Name, GlobalF.Name, GlobalFset.Position(fn.Pos()).Line)
				continue
			}

			var comment = fn.Doc.List[0]
			if strings.Index(comment.Text, fmt.Sprintf("// %s ", fn.Name)) != 0 {
				fmt.Printf("Warning: The comment of function %s is not standard\n", fn.Name.Name)
			}
		}
	}

	ast.Walk(&visitor{}, GlobalF) // 遍历AST树
}
