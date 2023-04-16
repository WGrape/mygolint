package redis_analyzer

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

// RedisAnalyzer 定义分析器
var RedisAnalyzer = &analysis.Analyzer{
	Name: "RedisAnalyzer",
	Doc:  "This is redis analyzer",
	Run:  run,
}

// run RedisAnalyzer的入口函数
func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			callExpr, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			selector, ok := callExpr.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			if selector.Sel.Name == "Close" {
				ident, ok := selector.X.(*ast.Ident)
				if !ok {
					return true
				}

				identType := pass.TypesInfo.Uses[ident].Type()
				identName := pass.TypesInfo.Uses[ident].Name()
				identString := pass.TypesInfo.Uses[ident].String()
				identPkg := pass.TypesInfo.Uses[ident].Pkg()
				fmt.Println(identType, identName, identString, identPkg)

				if identType.String() == "*github.com/go-redis/redis.Client" {
					pass.Reportf(callExpr.Pos(), "call to redis.Close")
				}
			}
			return true
		})
	}
	return nil, nil
}
