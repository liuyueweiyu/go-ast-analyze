package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"

	"github.com/fatih/color"
)

// Parse Parse
func Parse(filename string, source interface{}) (d *DeclFile, dump string, err error) {
	d = ParsePackage(filename, source)
	return d, "", nil
}

// ParsePackage ParsePackage
func ParsePackage(filename string, source interface{}) *DeclFile {
	d := &DeclFile{}
	fset := token.NewFileSet()
	src, _ := filepath.Abs(srcReadPath + filename)

	f, err := parser.ParseFile(fset, src, source, 0)
	if err != nil {
		color.Red("%s解析失败，错误原因:\n%v", filename, err)
		return nil
	}
	// 输出ast
	fileWriteAST(astPrintPath+GetFileName(filename)+".txt", fset, f)
	d.PackageName = f.Name.Name
	d.ImportList = ParseImports(f.Imports)
	d.Decls, d.Values = ParseDeclList(f.Decls)
	return d
}

// ParseImports 解析imports
func ParseImports(imports []*ast.ImportSpec) []string {
	importList := make([]string, 0)
	for _, s := range imports {
		importList = append(importList, s.Path.Value)
	}
	return importList
}

// ParseDeclList ParseDecl
func ParseDeclList(astDecls []ast.Decl) ([]DeclType, []*BaseValue) {
	based := make([]DeclType, 0)
	basev := make([]*BaseValue, 0)
	// 获取全局定义函数/变量
	for _, astDecl := range astDecls {
		d, v := ParseDecl(astDecl)
		based = append(based, d...)
		basev = append(basev, v...)
	}
	return based, basev

}

// ParseDecl ParseDecl
func ParseDecl(astDecl ast.Decl) ([]DeclType, []*BaseValue) {
	based := make([]DeclType, 0)
	basev := make([]*BaseValue, 0)
	switch astDecl.(type) {
	case *ast.GenDecl:
		genDecl := astDecl.(*ast.GenDecl)
		switch genDecl.Tok {
		case token.IMPORT:
		case token.TYPE:
			for _, value := range genDecl.Specs {
				based = append(based, ParseTypeSpec(value.(*ast.TypeSpec))...)
			}
		case token.VAR, token.CONST:
			for _, value := range genDecl.Specs {
				basev = append(basev, ParseValueSpec(value.(*ast.ValueSpec))...)
			}
		default:
			color.Red("ParseDecl:出现未知type")
		}
	case *ast.FuncDecl:
		basev = append(basev, ParseFuncDecl(astDecl.(*ast.FuncDecl)))
	default:
		color.Red("ParseDecl:全局定义函数/变量类型分析错误:出现未知变量类型")
	}
	return based, basev
}

// ParseFuncDecl BuildFuncDecl
func ParseFuncDecl(astf *ast.FuncDecl) *BaseValue {
	f := &BaseValue{}
	color.Green("正在解析FuncDecl:%s", astf.Name.Name)
	if astf.Name.Name == "InitBusinessAPI" {
		color.Green("InitBusinessAPI:%s", astf.Name.Name)
	}
	f.Name = astf.Name.Name
	d := &FuncType{}
	d.Type = "func"
	recv := make([]*BaseValue, 0)
	// 入参
	if astf.Type.Params != nil {
		recv = append(recv, GetBaseDeclListFromFieldList(astf.Type.Params.List)...)
	}
	if astf.Recv != nil {
		r := GetBaseDeclListFromFieldList(astf.Recv.List)
		if len(r) != 1 {
			color.Red("解析函数时，调用对象不止一个")
		} else {
			d.Attrs = make(map[string]string, 0)
			d.Attrs["objectName"] = r[0].Name
			d.Attrs["objectType"] = r[0].Type.GetTypeName()
		}
	}
	d.RecvArgs = recv
	// 返回值
	if astf.Type.Results != nil {
		d.RetuArgs = GetBaseDeclListFromFieldList(astf.Type.Results.List)
	}
	b, v, c := ParseStmt(astf.Body)
	d.Children = b
	d.Decls = v
	d.CalledFuncs = c
	f.Type = d

	return f
}

// ParseSpec ParseSpec
func ParseSpec(n interface{}) ([]*BaseValue, []DeclType) {
	basev := make([]*BaseValue, 0)
	based := make([]DeclType, 0)
	switch n.(type) {
	case *ast.ValueSpec:
		x := ParseValueSpec(n.(*ast.ValueSpec))
		basev = append(basev, x...)
	case *ast.TypeSpec:
		based = append(based, ParseTypeSpec(n.(*ast.TypeSpec))...)
	default:
		color.Red("ParseSpec:出现为预料变量类型")
	}
	return basev, based
}

// ParseValueSpec ParseValueSpec
func ParseValueSpec(n *ast.ValueSpec) []*BaseValue {
	values := make([]*BaseValue, 0)
	switch n.Type.(type) {
	case *ast.MapType, *ast.ArrayType, *ast.FuncType, *ast.InterfaceType, *ast.Ident, *ast.ChanType:
		for _, name := range n.Names {
			color.Green("正在解析ValueSpec:%s", name.Name)
			d := GetBaseDecl(n.Type)
			v := &BaseValue{Name: name.Name, Type: d}
			values = append(values, v)
		}
	case nil:
		for index := 0; index < len(n.Names); index++ {
			color.Green("正在解析ValueSpec:%s", n.Names[index].Name)
			d := GetBaseDecl(n.Values[index])
			v := &BaseValue{Name: n.Names[index].Name, Type: d}
			values = append(values, v)
		}
	case *ast.BasicLit:
		Log("BasicLit")
	default:
		color.Red("ParseValueSpec:出现为预料变量类型")
	}
	return values
}

// ParseTypeSpec ParseTypeSpec
func ParseTypeSpec(n *ast.TypeSpec) []DeclType {
	gdls := make([]DeclType, 0)
	color.Green("正在解析TypeSpec:%s", n.Name.Name)
	switch n.Type.(type) {
	case *ast.StructType:
		g := GetBaseDecl(n.Type)
		g.SetTypeAndName("struct", n.Name.Name)
		gdls = append(gdls, g)
	case *ast.InterfaceType:
		g := GetBaseDecl(n.Type)
		g.SetTypeAndName("interface", n.Name.Name)
		gdls = append(gdls, g)
	default:
		color.Red("ParseTypeSpec:出现为预料变量类型")
	}
	return gdls
}
