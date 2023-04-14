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

func TestAST(t *testing.T) {
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

	fmt.Printf("\n\n打印一下pass.Files中每一个文件的AST节点\n")
	fmt.Printf("--------------------------------------------\n")
	ast.Inspect(GlobalF, func(n ast.Node) bool {

		fmt.Printf("node = %+v\n", n)

		if node, ok := n.(*ast.File); ok {
			fmt.Printf("ast.File: %v\n", node.Name)
		}

		// 标识符
		if node, ok := n.(*ast.Ident); ok {
			fmt.Printf("ast.Ident: %v\n", node.Name)
		}

		// 数组类型，包括数组长度和元素类型等
		if node, ok := n.(*ast.ArrayType); ok {
			fmt.Printf("ast.ArrayType: %v\n", node.Elt)
		}

		if node, ok := n.(*ast.AssignStmt); ok {
			fmt.Printf("ast.AssignStmt: %v\n", node.Tok)
		}

		if node, ok := n.(*ast.BranchStmt); ok {
			fmt.Printf("ast.AssignStmt: %v\n", node.Tok)
		}

		// 函数声明
		if decl, ok := n.(*ast.FuncDecl); ok {
			fmt.Printf("ast.FuncDecl: %v, %+v\n", decl.Name.Name, decl.Recv)
		}
		return true
	})

	ast.Walk(&visitor{}, GlobalF) // 遍历AST树
}
