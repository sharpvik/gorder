package order

import (
	"go/ast"
	"go/token"
	"log"
)

var order = map[token.Token]int{
	token.IMPORT: 0,
	token.CONST:  1,
	token.VAR:    2,
	token.TYPE:   3,
	token.FUNC:   4,
}

// TODO: also sort by exported status and alphabetically.
func declOrder(decl ast.Decl) int {
	if _, isFunc := decl.(*ast.FuncDecl); isFunc {
		return order[token.FUNC]
	}
	return order[decl.(*ast.GenDecl).Tok]
}

// TODO: take non-doc comments into consideration.
func declPos(decl ast.Decl) (from, to token.Pos) {
	switch x := decl.(type) {
	case *ast.FuncDecl:
		if x.Doc != nil {
			from = x.Doc.Pos() - 1
		} else {
			from = x.Pos() - 1
		}
		return from, decl.End() - 1
	case *ast.GenDecl:
		if x.Doc != nil {
			from = x.Doc.Pos() - 1
		} else {
			from = x.Pos() - 1
		}
		return from, decl.End() - 1
	default:
		log.Fatal("found a bad declaration in file")
		return
	}
}
