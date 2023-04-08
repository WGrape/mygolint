package ast

import (
	"go/ast"
	"log"
)

// visitor 实现了 Visitor 接口，用于遍历节点并打印信息
type visitor struct{}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	// 打印节点类型和位置信息
	if node != nil {
		log.Printf("Visit: node = %v, pos = %v, end = %v\n", node, node.Pos(), node.End())
	}
	return v
}
