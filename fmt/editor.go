package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"go/ast"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		color.Red("Error: 请输入正确路径")
		src, _ := filepath.Abs("./temp_src/pkga/name.go")
		// EditorPackageName("pkga", src)
		AddImport("hhh", src)
		fmt.Println(src)
	} else {
		filename := os.Args[1]
		fmt.Println(filename)
	}
}

// EditorPackageName EditorPackageName
func EditorPackageName(name string, filepath string) {
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, filepath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	f.Name.Name = name
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		panic(err)
	}
	bites := buf.Bytes()
	fileWrite(string(bites[:]), filepath)
}

// AddImport AddImport
func AddImport(name string, filepath string) {
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, filepath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// ast.SortImports(fset, f)
	for _, d := range f.Decls {
		d, ok := d.(*ast.GenDecl)
		if !ok || d.Tok != token.IMPORT {
			// Not an import declaration, so we're done.
			// Imports are always first.
			break
		}
		if d.Specs
	}

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		panic(err)
	}
	bites := buf.Bytes()
	fileWrite(string(bites[:]), filepath)
}

func fileWrite(sourceStr string, path string) {
	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer f.Close()
	f.WriteString(sourceStr)
}
