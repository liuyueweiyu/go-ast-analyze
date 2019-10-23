package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// GetFilesOfDir GetFilesOfDir
func GetFilesOfDir(path string) []string {
	all := make([]string, 0)
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if !file.IsDir() {
			all = append(all, file.Name())
		}
	}
	return all
}

// GetFileName GetFileName
func GetFileName(fullName string) string {
	a := strings.Split(fullName, ".")
	a = a[0 : len(a)-1]
	return strings.Join(a, ".")
}

func fileWrite(sourceStr string, name string) {
	path, _ := filepath.Abs(name)
	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer f.Close()
	f.WriteString(sourceStr)
}

func fileWriteAST(name string, fset *token.FileSet, x interface{}) {
	var buf bytes.Buffer
	buf.Reset()
	if err := ast.Fprint(&buf, fset, x, nil); err != nil {
		Log("ast输出至文本文件错误....")
	}
	path, _ := filepath.Abs(name)
	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer f.Close()
	f.WriteString(buf.String())
}

// TestInspect TestInspect
// dfs 不能确定之间的关系就很蛋疼==
func TestInspect(filename string, source interface{}) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, source, 0)
	if err != nil {
		Log(filename, "解析失败，错误原因:", err)
	} else {
		// 输出ast
		fileWriteAST("./temp_res/ast.txt", fset, f)
		ast.Inspect(f, func(n ast.Node) bool {
			var s string
			switch x := n.(type) {
			case *ast.BasicLit:
				s = x.Value
			case *ast.Ident:
				s = x.Name
			}
			if s != "" {
				fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
			}
			return true
		})
	}
}

// MinInt MinInt
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TestFunc TestFunc
func TestFunc() {
	f := func(str string) {
		Log("hhhh", str)
	}
	f("123")
}
