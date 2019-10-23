package main

import (
	"go/ast"
	"go/token"
	"strings"

	"github.com/fatih/color"
)

// GetAssignStmtRHSType GetAssignStmtRhsType
func GetAssignStmtRHSType(rhs []ast.Expr) ([]DeclType, []*SelectorExpr) {
	decls := make([]DeclType, 0)
	called := make([]*SelectorExpr, 0)
	for _, rh := range rhs {
		d, _, c := ParseExpr(rh)
		decls = append(decls, d...)
		called = append(called, c...)
	}
	return decls, called
}

// ParseExpr ParseExpr
func ParseExpr(expr ast.Expr) ([]DeclType, []*BaseValue, []*SelectorExpr) {
	decls := make([]DeclType, 0)
	called := make([]*SelectorExpr, 0)
	switch expr.(type) {
	case *ast.Ident, *ast.CompositeLit, *ast.BasicLit:
		d := GetBaseDecl(expr)
		decls = append(decls, d)
	case *ast.CallExpr:
		fun := expr.(*ast.CallExpr).Fun
		called = append(called, GetCallExp(expr.(*ast.CallExpr)))
		switch fun.(type) {
		case *ast.Ident:
			i := fun.(*ast.Ident)

			switch i.Name {
			case "make":
				decls = append(decls, GetBaseDecl(expr.(*ast.CallExpr).Args[0]))
			case "len":
				decls = append(decls, &SimpleType{Type: KindDir[token.INT]})
			case "bite":
				decls = append(decls, &SimpleType{Type: "bite"})
			default:
				if i.Obj != nil {
					d := GetBaseDecl(i.Obj.Decl.(*ast.FuncDecl).Type)
					if d != nil && d.GetRetuArgs() != nil {
						for _, v := range d.GetRetuArgs() {
							decls = append(decls, v.Type)
						}
					}
				} else {
					// 系统自带函数
					color.Green("")
				}
			}
			// if i.Obj != nil {
			// 	d := GetBaseDecl(i.Obj.Decl.(*ast.FuncDecl).Type)
			// 	if d != nil && d.GetRetuArgs() != nil {
			// 		for _, v := range d.GetRetuArgs() {
			// 			decls = append(decls, v.Type)
			// 		}
			// 	}
			// } else {
			// 	// 系统自带函数

			// }
		case *ast.SelectorExpr:
			_, c := GetAssignStmtRHSType([]ast.Expr{fun.(*ast.SelectorExpr).X, fun.(*ast.SelectorExpr).Sel})
			called = append(called, c...)
		case *ast.ArrayType:
			_, c := GetAssignStmtRHSType([]ast.Expr{fun.(*ast.ArrayType).Elt})
			called = append(called, c...)
		}
		_, c := GetAssignStmtRHSType(expr.(*ast.CallExpr).Args)
		called = append(called, c...)
	case *ast.UnaryExpr:
		exprs := []ast.Expr{expr.(*ast.UnaryExpr).X}
		d, c := GetAssignStmtRHSType(exprs)
		decls = append(decls, d...)
		called = append(called, c...)
	case *ast.SelectorExpr:
		d, c := GetAssignStmtRHSType([]ast.Expr{expr.(*ast.SelectorExpr).X, expr.(*ast.SelectorExpr).Sel})
		called = append(called, c...)
		str := make([]string, 0)
		for _, v := range d {
			str = append(str, v.GetTypeName())
		}
		decls = append(decls, &SimpleType{Type: strings.Join(str, ".")})
	case *ast.IndexExpr:
		d, c := GetAssignStmtRHSType([]ast.Expr{expr.(*ast.IndexExpr).X})
		called = append(called, c...)
		decls = append(decls, d...)
	case *ast.FuncLit:
		funcLit := expr.(*ast.FuncLit)
		f := GetBaseDecl(funcLit.Type)
		b, d, s := ParseStmt(funcLit.Body)
		f.(*FuncType).Decls = d
		f.(*FuncType).Children = b
		f.(*FuncType).CalledFuncs = s
		decls = append(decls, f)
	default:
		color.Red("GetAssignStmtRHSType:出现未知assignExpr")
		d := &SimpleType{Type: "TODO"}
		decls = append(decls, d)
	}
	return decls, nil, called

}

// ParseStmt ParsrStmt
func ParseStmt(stmt interface{}) ([]*BaseValue, []DeclType, []*SelectorExpr) {
	values := make([]*BaseValue, 0)
	decls := make([]DeclType, 0)
	called := make([]*SelectorExpr, 0)
	switch stmt.(type) {
	case *ast.BadStmt:
		Log("BadStmt")
	case *ast.BlockStmt:
		list := stmt.(*ast.BlockStmt).List
		for _, s := range list {
			b, d, c := ParseStmt(s)
			values = append(values, b...)
			decls = append(decls, d...)
			called = append(called, c...)
		}
	case *ast.BranchStmt:
		Log("BranchStmt")
	case *ast.DeclStmt:
		d, b := ParseDecl(stmt.(*ast.DeclStmt).Decl)
		values = append(values, b...)
		decls = append(decls, d...)
	case *ast.DeferStmt:
		Log("DeferStmt")
	case *ast.ExprStmt:
		b := GetCallExp(stmt.(*ast.ExprStmt).X.(*ast.CallExpr))
		called = append(called, b)
		_, c := GetAssignStmtRHSType(stmt.(*ast.ExprStmt).X.(*ast.CallExpr).Args)
		called = append(called, c...)
	case *ast.EmptyStmt:
		Log("EmptyStmt")
	case *ast.ForStmt:
		v, d, c := ParseStmt(stmt.(*ast.ForStmt).Body)
		values = append(values, v...)
		decls = append(decls, d...)
		called = append(called, c...)
	case *ast.GoStmt:
		Log("GoStmt")
	case *ast.IfStmt:
		v, d, c := ParseStmt(stmt.(*ast.IfStmt).Body)
		called = append(called, c...)
		values = append(values, v...)
		decls = append(decls, d...)
		v, d, c = ParseStmt(stmt.(*ast.IfStmt).Else)
		called = append(called, c...)
		values = append(values, v...)
		decls = append(decls, d...)
	case *ast.IncDecStmt:
		Log("IncDecStmt")
	case *ast.LabeledStmt:
		Log("LabeledStmt")
	case *ast.RangeStmt:
		b, d, c := ParseStmt(stmt.(*ast.RangeStmt).Body)
		values = append(values, b...)
		decls = append(decls, d...)
		called = append(called, c...)
	case *ast.ReturnStmt:
		Log("ReturnStmt")
	case *ast.SelectStmt:
		Log("SelectStmt")
	case *ast.SendStmt:
		Log("SendStmt")
	case *ast.SwitchStmt:
		v, d, c := ParseStmt(stmt.(*ast.SwitchStmt).Body)
		values = append(values, v...)
		decls = append(decls, d...)
		called = append(called, c...)
	case *ast.AssignStmt:
		if stmt.(*ast.AssignStmt).Tok == token.DEFINE {
			assign := stmt.(*ast.AssignStmt)
			rhDecls, cs := GetAssignStmtRHSType(assign.Rhs)
			called = append(called, cs...)
			if len(rhDecls) == len(assign.Lhs) {
				for index := 0; index < len(assign.Lhs); index++ {
					b := &BaseValue{
						Name: assign.Lhs[index].(*ast.Ident).Name,
						Type: rhDecls[index],
					}
					values = append(values, b)
				}
			} else {
				color.Red("parse:AssignStmt len(rhDecls) != len(assign.Lhs)")
				color.Red("rhs lhs: %d %d", len(rhDecls), len(assign.Lhs))
			}
		}
	case *ast.TypeSwitchStmt:
		Log("TypeSwitchStmt")
	case *ast.CaseClause:
		for _, s := range stmt.(*ast.CaseClause).Body {
			v, d, c := ParseStmt(s)
			values = append(values, v...)
			decls = append(decls, d...)
			called = append(called, c...)
		}
	default:
		Log("自闭解析stmt")
	}
	return values, decls, called
}
