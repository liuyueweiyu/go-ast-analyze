package main

import (
	"encoding/json"
	"go/token"

	"github.com/fatih/color"
)

// KindDir KindDir
var KindDir map[token.Token]string

var isDebug = false
var isTest = false
var jsonPrintPath string
var astPrintPath string
var srcReadPath string
var root string
var begin = 0
var end = 1

func main() {
	if isDebug {
		root = ".."
	} else {
		root = "."
	}
	InitKindDir()
	if isTest {
		jsonPrintPath = root + "/temp_res/"
		srcReadPath = root + "/temp_src/"
		astPrintPath = root + "/temp_res/"
		d, _, _ := Parse("src.go", nil)
		res, _ := json.Marshal(d)
		fileWrite(string(res[:]), jsonPrintPath+"ast.json")
	} else {
		jsonPrintPath = root + "/json/"
		astPrintPath = root + "/ast/"
		srcReadPath = root + "/src/"
		filenames := GetFilesOfDir(srcReadPath)
		length := MinInt(end, len(filenames))
		for index := begin; index < length; index++ {
			var filename = filenames[index]
			color.Blue("开始解析%s...", filename)
			d, _, _ := Parse(filename, nil)
			res, _ := json.Marshal(d)
			fileWrite(string(res[:]), jsonPrintPath+GetFileName(filename)+".json")
		}
	}

}

// InitKindDir Init
func InitKindDir() {
	KindDir = make(map[token.Token]string)
	KindDir[token.INT] = "int"
	KindDir[token.FLOAT] = "float"
	KindDir[token.IMAG] = "imag"
	KindDir[token.CHAR] = "char"
	KindDir[token.STRING] = "string"
	KindDir[token.STRUCT] = "struct"
	KindDir[token.MAP] = "map"
	KindDir[token.INTERFACE] = "interface"
	KindDir[token.CHAN] = "chan"
	KindDir[token.FUNC] = "func"
}
