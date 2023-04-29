package server_analyzer

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"os"
)

// ServerAnalyzer 服务分析器
var ServerAnalyzer = &analysis.Analyzer{
	Name: "ServerAnalyzer",
	Doc:  "This is server analyzer",
	Run:  run,
}

// IsServerComponentInit 是否调用了server.Init()
func IsServerComponentInit(pass *analysis.Pass, n ast.Node) bool {
	callExpr, ok := n.(*ast.CallExpr)
	if !ok {
		return false
	}

	var selector *ast.SelectorExpr
	if selector, ok = callExpr.Fun.(*ast.SelectorExpr); !ok {
		return false
	}
	if selector.Sel.Name != "Init" {
		return false
	}

	var ident *ast.Ident
	if ident, ok = selector.X.(*ast.Ident); !ok {
		return false
	}
	if ident == nil {
		return false
	}
	if ident.Name != "server" {
		return false
	}

	return pass.TypesInfo.Uses[ident].String() == "package server (\"github.com/xgo/server/\")"
}

// run ServerAnalyzer的入口函数
func run(pass *analysis.Pass) (interface{}, error) {
	var isServerComponentInit bool
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if IsServerComponentInit(pass, n) {
				isServerComponentInit = true
				return false
			}
			return true
		})
	}

	// 开启了服务配置, 但是未初始化服务
	var err error
	var configFile = "testdata/config.toml"
	if err = isServerConfigOk(configFile); err == nil && !isServerComponentInit {
		fmt.Println("Error: 未开启服务, 但Toml文件存在服务配置")
		os.Exit(1)
	}

	// 初始化服务, 但是解析服务配置异常
	if isServerComponentInit && err != nil {
		fmt.Printf("Error: 开启了服务, 但Toml文件存在异常: %+v", err.Error())
		os.Exit(1)
	}

	fmt.Println("服务检测通过")

	return nil, nil
}
