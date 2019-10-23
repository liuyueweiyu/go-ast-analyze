package main

import (
	"go/ast"
	"go/token"

	"github.com/fatih/color"
)

// GetBaseDeclListFromFieldList GetBaseDeclListFromAstField
func GetBaseDeclListFromFieldList(list []*ast.Field) []*BaseValue {
	if list == nil {
		return nil
	}
	baseValues := make([]*BaseValue, 0)
	for _, filed := range list {
		if filed.Names != nil {
			for _, name := range filed.Names {
				g := GetBaseDecl(filed.Type)
				b := &BaseValue{Name: name.Name, Type: g}
				baseValues = append(baseValues, b)
			}
		} else {
			g := GetBaseDecl(filed.Type)
			b := &BaseValue{Type: g}
			baseValues = append(baseValues, b)
		}
	}
	return baseValues
}

// GetBaseDecl GetBaseDecl
func GetBaseDecl(n interface{}) DeclType {
	switch n.(type) {
	case *ast.ArrayType:
		b := &ComplexType{}
		b.Type = "array"
		b.Child = GetBaseDecl(n.(*ast.ArrayType).Elt)
		return b
	case *ast.MapType:
		m := n.(*ast.MapType)
		b := &ComplexType{}
		b.Attrs = make(map[string]string)
		b.Attrs["keyType"] = m.Key.(*ast.Ident).Name
		b.Type = KindDir[token.MAP]
		b.Child = GetBaseDecl(m.Value)
		return b
	case *ast.ChanType:
		d := &ComplexType{}
		d.Type = KindDir[token.CHAN]
		d.Child = GetBaseDecl(n.(*ast.ChanType).Value)
		return d
	case *ast.Ident:
		b := &SimpleType{}
		b.Type = n.(*ast.Ident).Name
		return b
	case *ast.FuncType:
		funcObj := n.(*ast.FuncType)
		f := &FuncType{}
		f.Type = "func"
		if funcObj.Params != nil {
			f.RecvArgs = GetBaseDeclListFromFieldList(funcObj.Params.List)
		}
		if funcObj.Results != nil {
			f.RetuArgs = GetBaseDeclListFromFieldList(funcObj.Results.List)
		}
		return f
	case *ast.Field:
		d := &SimpleType{}
		return d
	case *ast.BasicLit:
		d := &SimpleType{}
		d.Type = KindDir[n.(*ast.BasicLit).Kind]
		return d
	case *ast.CompositeLit:
		return GetBaseDecl(n.(*ast.CompositeLit).Type)
	case *ast.InterfaceType:
		g := &CustomizeType{}
		g.Type = KindDir[token.INTERFACE]
		interfaceObj := n.(*ast.InterfaceType)
		parents := make([]string, 0)
		children := make([]*BaseValue, 0)
		fields := interfaceObj.Methods.List
		for _, field := range fields {
			if len(field.Names) != 0 {
				// 不继承interface
				for _, ident := range field.Names {
					f := GetBaseDecl(field.Type)
					b := &BaseValue{Name: ident.Name, Type: f}
					children = append(children, b)
				}
			} else {
				parents = append(parents, field.Type.(*ast.Ident).Name)
			}
		}
		g.Children = children
		g.Parents = parents
		return g
	case *ast.StructType:
		g := &CustomizeType{}
		g.Type = KindDir[token.STRUCT]
		stuctObj := n.(*ast.StructType)
		children := make([]*BaseValue, 0)
		parents := make([]string, 0)
		fields := stuctObj.Fields.List
		for _, field := range fields {
			if len(field.Names) != 0 {
				// 不继承struct
				for _, ident := range field.Names {
					d := GetBaseDecl(field.Type)
					children = append(children, &BaseValue{Name: ident.Name, Type: d})
				}
			} else {
				// 继承struct
				parents = append(parents, field.Type.(*ast.Ident).Name)
			}
		}
		g.Children = children
		g.Parents = parents
		return g
	case *ast.StarExpr:
		return GetBaseDecl(n.(*ast.StarExpr).X)
	case *ast.FuncLit:
		color.Yellow("fFuncLit:出现未知类型")
		return &SimpleType{Type: "UnKnownType"}
	case *ast.SelectorExpr:
		d := &ReferType{}
		d.Name = n.(*ast.SelectorExpr).X.(*ast.Ident).Name
		d.Package = n.(*ast.SelectorExpr).Sel.Name
		return d
	case nil:
		color.Yellow("NIL")
		return &SimpleType{Type: "UnKnownType"}
	default:
		color.Red("GetBaseDecl:出现未知类型")
		return &SimpleType{Type: "UnKnownType"}
	}
}
