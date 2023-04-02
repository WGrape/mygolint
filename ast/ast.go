package ast

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

var MyAstAnalyzer = &analysis.Analyzer{
	Name: "MyAstAnalyzer",
	Doc:  "This is my analyzer for AST testing",
	Run:  run,
}

func PrintWhatIsPass(pass *analysis.Pass) {

	// 打印分析器的基本信息
	fmt.Printf("\n\n下面是打印分析器的基本信息\n")
	fmt.Printf("--------------------------------------------\n")
	fmt.Printf("pass.Analyzer: %+v \n", pass.Analyzer)
	fmt.Printf("pass.Analyzer Name: %+v \n", pass.Analyzer.Name)
	fmt.Printf("pass.Analyzer Doc: %+v \n", pass.Analyzer.Doc)
	fmt.Printf("pass.Analyzer FactTypes: %+v \n", pass.Analyzer.FactTypes)
	fmt.Printf("pass.Analyzer Flags: %+v \n", pass.Analyzer.Flags)
	fmt.Printf("pass.Analyzer Requires: %+v \n", pass.Analyzer.Requires)
	fmt.Printf("pass.Analyzer ResultType: %+v \n", pass.Analyzer.ResultType)
	if pass.Analyzer.ResultType != nil {
		fmt.Printf("pass.Analyzer ResultType.Name(): %+v \n", pass.Analyzer.ResultType.Name())
		fmt.Printf("pass.Analyzer ResultType.String(): %+v \n", pass.Analyzer.ResultType.String())
		fmt.Printf("pass.Analyzer ResultType.String(): %+v \n", pass.Analyzer.ResultType.Len())
		fmt.Printf("pass.Analyzer ResultType.String(): %+v \n", pass.Analyzer.ResultType.Size())
		fmt.Printf("pass.Analyzer ResultType.String(): %+v \n", pass.Analyzer.ResultType.Align())
		fmt.Printf("pass.Analyzer ResultType.String(): %+v \n", pass.Analyzer.ResultType.Bits())
		fmt.Printf("pass.Analyzer ResultType.String(): %+v \n", pass.Analyzer.ResultType.NumField())
		fmt.Printf("pass.Analyzer ResultType.String(): %+v \n", pass.Analyzer.ResultType.NumMethod())
		fmt.Printf("pass.Analyzer ResultType.String(): %+v \n", pass.Analyzer.ResultType.NumIn())
		fmt.Printf("pass.Analyzer ResultType.String(): %+v \n", pass.Analyzer.ResultType.NumOut())
		fmt.Printf("pass.Analyzer ResultType.String(): %+v \n", pass.Analyzer.ResultType.Key())
		fmt.Printf("pass.Analyzer ResultType.String(): %+v \n", pass.Analyzer.ResultType.PkgPath())

	}
	fmt.Printf("pass.Analyzer RunDespiteErrors: %+v \n", pass.Analyzer.RunDespiteErrors)

	// pass.Fset表示一个文件集合(token.FileSet 类型的对象), 它用于在代码分析过程中跟踪源代码中的位置信息, 如文件名、行号、列号等
	// 在Go语言的静态代码分析中，经常需要获取分析目标代码中的某些位置信息，例如某个变量的定义位置、某个函数的调用位置等等。
	// 而pass.Fset可以帮助分析器将源代码中的位置信息与程序元素（例如函数、变量等）关联起来，从而方便分析器进行进一步的分析和处理。
	fmt.Printf("pass.Fset: %+v \n", pass.Fset)
	fmt.Printf("pass.Fset Base(): %+v \n", pass.Fset.Base())
	fmt.Printf("pass.Files: %+v \n", pass.Files)
	for _, v := range pass.Files {
		// Unresolved []*Ident: Unresolved是ast.File结构体中的一个字段，它是一个包含所有未解析的标识符的切片。每个ast.Ident对象代表了源代码中的一个标识符，例如变量名、函数名等等。
		// 当 Go 语言的解析器（parser）在解析源代码时遇到一个标识符，它会尝试在当前的作用域（scope）中查找这个标识符的定义，以确定这个标识符表示了什么。如果解析器无法找到这个标识符的定义，那么它就会将这个标识符添加到 Unresolved 字段中，以表示这个标识符是未解析的。
		// 未解析的标识符可能有多种原因。例如，它可能是因为在源代码中拼写错误导致无法找到它的定义，或者它可能是在当前文件之外的某个文件中定义的，但是还没有被正确导入。在这种情况下，编译器或解析器就无法确定这个标识符所代表的含义，因此就将它标记为未解析的。

		// []Decl: 是指在源代码中定义变量、函数、类型、常量等等的语句。在Go语言中，声明可以分为四种类型：变量声明/类型声明/函数声明/常量声明

		// Scope 结构体表示一个作用域（scope），它包含了一组声明（declaration）和一个指向外部作用域（outer scope）的指针。
		// 外部作用域（outer scope）是指包含当前作用域的上一级作用域。例如，在一个函数内部定义的变量具有函数作用域，在这个函数的外部是无法访问这个变量的。但是，如果在函数内部嵌套了另一个函数，那么在内部函数中定义的变量就具有外部函数的作用域。因此，内部函数的作用域是包含在外部函数的作用域之内的。
		// 在go中每一个花括号都是一个作用域, 都会有一个连接上层作用域的outer指针

		fmt.Printf("pass.Files: Name = %v, Doc = %v, Pos = %v, Imports = %v, Unresolved = %v, Comments = %v, Objects = %v, Outer = %v, Decls = %v\n", v.Name, v.Doc, v.Pos(), v.Imports, v.Unresolved, v.Comments, v.Scope.Objects, v.Scope.Outer, v.Decls)

		// 打印注释相关的信息
		for _, commentGroup := range v.Comments {
			fmt.Printf("Pos = %v, End = %v, Text = %v", commentGroup.Pos(), commentGroup.End(), commentGroup.Text())
			for _, comment := range commentGroup.List {
				fmt.Printf("Slash = %v, Text = %v\n", comment.Slash, comment.Text)
			}
		}
	}

	// 打印被分析的包信息
	fmt.Printf("\n\n下面是打印被分析的包信息\n")
	fmt.Printf("--------------------------------------------\n")
	fmt.Printf("pass.Pkg: %+v \n", pass.Pkg)
	fmt.Printf("pass.Pkg Path(): %+v \n", pass.Pkg.Path())
	fmt.Printf("pass.Pkg Name(): %+v \n", pass.Pkg.Name())
	fmt.Printf("pass.Pkg String(): %+v \n", pass.Pkg.String())
	fmt.Printf("pass.Pkg Complete(): %+v \n", pass.Pkg.Complete())
	fmt.Printf("pass.Pkg Imports(): %+v \n", pass.Pkg.Imports())

	// 打印被分析的包对象的Scope作用域信息
	fmt.Printf("\n\n下面是打印被分析的包对象的作用域信息\n")
	fmt.Printf("--------------------------------------------\n")
	fmt.Printf("pass.Pkg Scope(): %+v \n", pass.Pkg.Scope())
	fmt.Printf("pass.Pkg Scope().String(): %+v \n", pass.Pkg.Scope().String())
	fmt.Printf("pass.Pkg Scope().Pos(): %+v \n", pass.Pkg.Scope().Pos())
	fmt.Printf("pass.Pkg Scope().End(): %+v \n", pass.Pkg.Scope().End())
	fmt.Printf("pass.Pkg Scope().Len(): %+v \n", pass.Pkg.Scope().Len())
	fmt.Printf("pass.Pkg Scope().Parent(): %+v \n", pass.Pkg.Scope().Parent())
	fmt.Printf("pass.Pkg Scope().NumChildren(): %+v \n", pass.Pkg.Scope().NumChildren())
	fmt.Printf("pass.Pkg Scope().Names(): %+v \n", pass.Pkg.Scope().Names())
	for _, name := range pass.Pkg.Scope().Names() {
		// <重要>
		// 大部分场景都是需要对作用域内的函数、变量等分析
		var obj = pass.Pkg.Scope().Lookup(name)
		fmt.Printf("pass.Pkg Scope().Lookup(): obj = %+v \n, Type = %v, Id = %v, Name = %v, String = %v, Pos = %v, Exported = %v \n", obj, obj.Type(), obj.Id(), obj.Name(), obj.String(), obj.Pos(), obj.Exported())
	}

	// 打印被分析的包对象的ResultOf信息
	fmt.Printf("\n\n打印被分析的包对象的ResultOf信息\n")
	fmt.Printf("--------------------------------------------\n")
	fmt.Printf("pass.ResultOf: %+v \n", pass.ResultOf)
	fmt.Printf("pass.TypesInfo: %+v \n", pass.TypesInfo)
	fmt.Printf("pass.TypesInfo.Types: %+v \n", pass.TypesInfo.Types)
	fmt.Printf("pass.TypesInfo.Defs: %+v \n", pass.TypesInfo.Defs)
	fmt.Printf("pass.TypesInfo.Implicits: %+v \n", pass.TypesInfo.Implicits)
	for k, v := range pass.TypesInfo.Implicits {
		fmt.Printf("pass.TypesInfo.Implicits : ast.Node = %+v, Object = %+v \n", k, v)
	}

	// 打印被分析的包对象的一些关于文件相关的信息
	fmt.Printf("\n\n打印被分析的包对象的一些关于文件相关的信息\n")
	fmt.Printf("--------------------------------------------\n")
	fmt.Printf("Files = %v, OtherFiles = %v, IgnoredFiles = %v", pass.Files, pass.OtherFiles, pass.IgnoredFiles)

	// 打印被分析的包对象的一些其他无分类的信息
	fmt.Printf("\n\n打印被分析的包对象的一些其他无分类的信息\n")
	fmt.Printf("--------------------------------------------\n")
	fmt.Printf("pass.String(): %+v \n", pass.String())
}

func run(pass *analysis.Pass) (interface{}, error) {
	PrintWhatIsPass(pass)

	for _, f := range pass.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			if decl, ok := n.(*ast.FuncDecl); ok {
				fmt.Printf("Function %q\n", decl.Name.Name)
			}
			return true
		})
	}
	return nil, nil
}
