package main

import (
	"go/ast"
)

// GetCallExp GetCallExp
//  []DeclType, []*BaseValue,
func GetCallExp(n interface{}) *SelectorExpr {
	switch n.(type) {
	case *ast.CallExpr:
		s := GetCallExp(n.(*ast.CallExpr).Fun)
		if s != nil {
			cur := s
			for cur.Next != nil {
				cur = cur.Next
			}
			cur.Type = SelectorExprTypeFunc
			return s
		}
	case *ast.SelectorExpr:
		se := n.(*ast.SelectorExpr)
		s := GetCallExp(se.X)
		cur := s
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = GetCallExp(se.Sel)
		return s
	case *ast.Ident:
		s := &SelectorExpr{}
		s.Name = n.(*ast.Ident).Name
		s.Next = nil
		s.Type = SelectorExprTypeProp
		return s
	}
	return nil
}
