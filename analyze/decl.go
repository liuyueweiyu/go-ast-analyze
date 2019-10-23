package main

/*

package->domain

domain->{
	args	//变量
	types	//定义的types

}

//-----------------------------//
//	嵌套问题用atrrs+child解决	//
//-----------------------------//

struct/interface->存在继承
map存在嵌套

funcObj->{
	"name"
	"type"->"func"
	"recv"->也有可能是func
	"return"
}
变量类型
array->attr[arrType]="type"qwq那万一是n维数组呢
chan->attr[chanType]=""
func->
interface->
map->qwq map也存在嵌套问题
struct->

Stmt

*/

// DeclFile DeclsFile
type DeclFile struct {
	PackageName string       `json:"packageName"` // 包的名称
	Decls       []DeclType   `json:"decls"`       // 各种声明
	Values      []*BaseValue `json:"values"`
	ImportList  []string     `json:"importList"`
}

// StmtDomain StmtDomain
// 一个函数带有一个域?
// type StmtDomain struct {
// 	Name       string        `json:"name"`
// 	GenDecls   []*BaseDecl   `json:"genDecls"` // 各种声明
// 	ValueDecls []*BaseValue  `json:"valueDecls"`
// 	Children   []*StmtDomain `json:"stmtChildren"`
// }

// BaseValue BaseValue
type BaseValue struct {
	Name string   `json:"name"`
	Type DeclType `json:"type"`
}

// BaseDecl 变量定义类型理论上只有struct/interface
// 但是不排除 type Token int
// 一个函数定义应该是func funcName(...args) ...agrs
// 是一个BaseValue name->funcNmae
// Type 是一个 BaseDecl,TypeName 为func
// func 内不再解析的value值，只将其定义value的Type记录
// so...func里面再定义类型==咋整->忽略(丢到attrs里面?)
// type BaseDecl struct {
// 	TypeName  string            `json:"typeName"`
// 	Type      string            `json:"type"`
// 	IsSimple  bool              `json:"isSimple"`
// 	TypeValue string            `json:"isValue"` //简单类型读这里
// 	Parents   []string          `json:"parents"`
// 	Child     *BaseDecl         `json:"child"`
// 	Children  []*BaseValue      `json:"chilren"`
// 	Attrs     map[string]string `json:"attrs"`
// 	RecvArgs  []*BaseValue      `json:"recvArgs"`
// 	RetuArgs  []*BaseValue      `json:"retuArgs"`
// }

// DeclType DeclType
type DeclType interface {
	GetDeclType() string
	SetTypeAndName(string, string)
	GetTypeName() string
	GetRetuArgs() []*BaseValue
}

// SimpleType SimpleType
type SimpleType struct {
	Type string `json:"type"`
}

// ComplexType ComplexType
type ComplexType struct {
	Type  string            `json:"type"`
	Child DeclType          `json:"child"`
	Attrs map[string]string `json:"attrs"`
}

// CustomizeType CustomizeType
type CustomizeType struct {
	TypeName string            `json:"typeName"`
	Type     string            `json:"type"`
	Parents  []string          `json:"parents"`
	Children []*BaseValue      `json:"chilren"`
	Attrs    map[string]string `json:"attrs"`
}

// ReferType 引用类型
type ReferType struct {
	Name    string `json:"typeName"`
	Package string `json:"package"`
}

// FuncType FuncTyoe
type FuncType struct {
	Type        string            `json:"type"`
	RecvArgs    []*BaseValue      `json:"recvArgs"`
	RetuArgs    []*BaseValue      `json:"retuArgs"`
	Children    []*BaseValue      `json:"chilren"`
	Decls       []DeclType        `json:"decls"` // 各种声明
	CalledFuncs []*SelectorExpr   `json:"calledFuncs"`
	Attrs       map[string]string `json:"attrs"`
}

const (
	// SelectorExprTypeFunc SelectorExprTypeFunc
	SelectorExprTypeFunc = "func"
	// SelectorExprTypeProp SelectorExprTypeProp
	SelectorExprTypeProp = "prop"
)

// SelectorExpr SelectorExpr 链式调用
type SelectorExpr struct {
	Name string        `json:"name"`
	Type string        `json:"type"`
	Next *SelectorExpr `json:"next"`
}
