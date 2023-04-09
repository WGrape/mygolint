package ast

import (
	"fmt"
	golib "github.com/WGrape/golib/string"
	"go/ast"
	"go/token"
	"log"
)

// visitor 实现了 Visitor 接口，用于遍历节点并打印信息
type visitor struct{}

var GlobalF *ast.File
var GlobalFset *token.FileSet

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return v
	}

	// 打印节点类型和位置信息
	log.Printf("Visit: node = %v, pos = %v, end = %v\n", node, node.Pos(), node.End())

	// 检查所有标识符
	switch ident := node.(type) {
	case *ast.Ident:
		// 检查是否符合驼峰命名法
		if !golib.IsCamelCase(ident.Name) {
			fmt.Printf("naming %s is not camel case\n", ident.Name)
		}

		// 检查是否符合包名全部小写的规范
		if isPackageName(ident, GlobalF) && !golib.IsLower(([]rune)(ident.Name)) {
			fmt.Printf("package %s is not lower case\n", ident.Name)
		}
	}
	return v
}

// 判断标识符是否是包名
func isPackageName(ident *ast.Ident, node *ast.File) bool {
	return ident.Obj != nil && ident.Obj.Kind == ast.Pkg && ident.Obj.Decl == node.Name
}
