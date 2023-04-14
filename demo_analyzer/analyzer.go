package analyzer

import (
	"fmt"
	"golang.org/x/tools/go/analysis"
)

// DemoAnalyzer 定义分析器
var DemoAnalyzer = &analysis.Analyzer{
	Name: "DemoAnalyzer",
	Doc:  "This is my analyzer for testing",
	Run:  run,
}

// PrintWhatIsPass 打印一下pass是啥玩意
// pass它包含了关于源代码分析过程中需要的信息和上下文。它被用作传递给分析器函数的参数，以便分析器能够访问到关于源代码的信息和上下文，例如文件的AST（抽象语法树）、token（词法分析器生成的标记）等等。
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
	for _, f := range pass.Files {
		// Unresolved []*Ident: Unresolved是ast.File结构体中的一个字段，它是一个包含所有未解析的标识符的切片。每个ast.Ident对象代表了源代码中的一个标识符，例如变量名、函数名等等。
		// 当 Go 语言的解析器（parser）在解析源代码时遇到一个标识符，它会尝试在当前的作用域（scope）中查找这个标识符的定义，以确定这个标识符表示了什么。如果解析器无法找到这个标识符的定义，那么它就会将这个标识符添加到 Unresolved 字段中，以表示这个标识符是未解析的。
		// 未解析的标识符可能有多种原因。例如，它可能是因为在源代码中拼写错误导致无法找到它的定义，或者它可能是在当前文件之外的某个文件中定义的，但是还没有被正确导入。在这种情况下，编译器或解析器就无法确定这个标识符所代表的含义，因此就将它标记为未解析的。

		// []Decl: 是指在源代码中定义变量、函数、类型、常量等等的语句。在Go语言中，声明可以分为四种类型：变量声明/类型声明/函数声明/常量声明

		// Scope 结构体表示一个作用域（scope），它包含了一组声明（declaration）和一个指向外部作用域（outer scope）的指针。
		// 外部作用域（outer scope）是指包含当前作用域的上一级作用域。例如，在一个函数内部定义的变量具有函数作用域，在这个函数的外部是无法访问这个变量的。但是，如果在函数内部嵌套了另一个函数，那么在内部函数中定义的变量就具有外部函数的作用域。因此，内部函数的作用域是包含在外部函数的作用域之内的。
		// 在go中每一个花括号都是一个作用域, 都会有一个连接上层作用域的outer指针

		fmt.Printf("pass.Files: Name = %v, Doc = %v, Pos = %v, Imports = %v, Unresolved = %v, Comments = %v, Objects = %v, Outer = %v, Decls = %v\n", f.Name, f.Doc, f.Pos(), f.Imports, f.Unresolved, f.Comments, f.Scope.Objects, f.Scope.Outer, f.Decls)

		// 打印注释相关的信息
		for _, commentGroup := range f.Comments {
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
	// ResultOf表示分析器依赖的其他分析器
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
	// TypesInfo：该属性是一个types.Info类型的指针，表示语法树的类型信息。在源代码分析过程中，分析器可以使用TypesInfo属性来查询关于变量、类型、方法等信息。
	// TypesSizes：该属性是一个types.Sizes类型的值，表示计算类型大小的函数。在源代码分析过程中，如果需要计算类型的大小，分析器可以使用TypesSizes属性中的函数进行计算。
	// TypeErrors：该属性是一个types.Error类型的切片，表示类型错误。只有当分析器使用Analyzer.RunDespiteErrors方法运行时，TypeErrors属性才会被填充。
	// 如果分析器使用Analyzer.RunDespiteErrors方法运行，并且在分析过程中出现了类型错误，那么这些错误将被添加到TypeErrors切片中。在分析器完成分析后，可以使用这些错误来生成错误报告
	fmt.Printf("pass.TypeErrors: %+v \n", pass.TypeErrors)
	// 在源代码分析过程中，如果需要查询语法树中的类型信息，可以使用TypesInfo属性。例如，可以使用TypesInfo.TypeOf函数来获取一个表达式的类型
	fmt.Printf("pass.TypesInfo: %+v \n", pass.TypesInfo)
	// 如果需要计算某个类型的大小，可以使用TypesSizes属性中的函数。例如，可以使用TypesSizes.Size函数来计算一个类型的大小
	fmt.Printf("pass.TypesSizes: %+v \n", pass.TypesSizes)
}

// run DemoAnalyzer的入口函数
// 运行DemoAnalyzer的时候就会执行这个函数
func run(pass *analysis.Pass) (interface{}, error) {
	// 打印一下pass是啥玩意
	fmt.Printf("\n\n打印一下pass是啥玩意\n")
	fmt.Printf("--------------------------------------------\n")
	PrintWhatIsPass(pass)

	return nil, nil
}
