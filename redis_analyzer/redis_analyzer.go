package redis_analyzer

import (
	"fmt"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
)

// RedisAnalyzer 定义分析器
var RedisAnalyzer = &analysis.Analyzer{
	Name: "RedisAnalyzer",
	Doc:  "This is redis analyzer",
	Run:  run,
}

// isRedisType 是否是redis类型
func isRedisType(pass *analysis.Pass, ident *ast.Ident) bool {
	if ident == nil {
		return false
	}
	return pass.TypesInfo.Uses[ident].Type().String() == "*github.com/go-redis/redis.Client"
}

// isRedisPackageType 是否是redis package类型
func isRedisPackageType(pass *analysis.Pass, ident *ast.Ident) bool {
	if ident == nil {
		return false
	}
	var pkg = pass.TypesInfo.Uses[ident].(*types.PkgName)
	return pkg.String() == "package redis (\"github.com/go-redis/redis\")"
}

// isRedisNewClient
func isRedisNewClient(pass *analysis.Pass, n ast.Node) bool {
	callExpr, ok := n.(*ast.CallExpr)
	if !ok {
		return false
	}

	var selector *ast.SelectorExpr
	if selector, ok = callExpr.Fun.(*ast.SelectorExpr); !ok {
		return false
	}
	if selector.Sel.Name != "NewClient" {
		return false
	}

	var ident *ast.Ident
	if ident, ok = selector.X.(*ast.Ident); !ok {
		return false
	}

	return isRedisPackageType(pass, ident)
}

// isRedisClose
func isRedisClose(pass *analysis.Pass, n ast.Node) bool {
	callExpr, ok := n.(*ast.CallExpr)
	if !ok {
		return false
	}

	var selector *ast.SelectorExpr
	if selector, ok = callExpr.Fun.(*ast.SelectorExpr); !ok {
		return false
	}
	if selector.Sel.Name != "Close" {
		return false
	}

	var ident *ast.Ident
	if ident, ok = selector.X.(*ast.Ident); !ok {
		return false
	}

	return isRedisType(pass, ident)
}

// getEnclosingFunction 获取所在的函数
//     var enclosingFunction = getEnclosingFunction(pass, n)
//     fmt.Printf("\nenclosingFunction.Name = %v\n\n", enclosingFunction.Name)
//     pass.Reportf(callExpr.Pos(), "call to redis.Close")
func getEnclosingFunction(pass *analysis.Pass, node ast.Node) *ast.FuncDecl {
	pos := pass.Fset.Position(node.Pos())
	for _, file := range pass.Files {
		for _, f := range file.Decls {
			if fd, ok := f.(*ast.FuncDecl); ok {
				if pass.Fset.Position(fd.Pos()).Filename == pos.Filename && fd.Pos() <= node.Pos() && node.Pos() <= fd.End() {
					return fd
				}
			}
		}
	}
	return nil
}

// testCheck 一些测试性的验证检查
func testCheck(pass *analysis.Pass) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if isRedisNewClient(pass, n) {
				fmt.Println("开启了Redis")
			}
			if isRedisClose(pass, n) {
				fmt.Println("关闭了Redis")
			}
			return true
		})
	}
}

// run RedisAnalyzer的入口函数
func run(pass *analysis.Pass) (interface{}, error) {
	testCheck(pass)

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if isRedisNewClient(pass, n) {
				fmt.Println("创建了")
			}
			if isRedisClose(pass, n) {
				fmt.Println("关闭了")
			}
			return true
		})
	}
	return nil, nil
}
