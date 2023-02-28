package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"sort"

	cli "github.com/urfave/cli/v2"
)

var app = cli.App{
	Name:   "gorder",
	Action: action,
}

func action(c *cli.Context) error {
	filename := c.Args().First()
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	parsedFile, err := parser.ParseFile(
		token.NewFileSet(),
		"", bytes,
		parser.AllErrors,
	)
	if err != nil {
		return err
	}

	sort.Slice(parsedFile.Decls, func(i, j int) bool {
		return declOrder(parsedFile.Decls[i]) < declOrder(parsedFile.Decls[j])
	})

	for _, decl := range parsedFile.Decls {
		log.Printf("decl type '%T' at (%d -> %d)", decl, decl.Pos(), decl.End())
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			log.Printf("  gen decl token %v", genDecl.Tok)
		}
	}

	return nil
}

func declOrder(decl ast.Decl) int {
	if _, isFunc := decl.(*ast.FuncDecl); isFunc {
		return order[token.FUNC]
	}
	return order[decl.(*ast.GenDecl).Tok]
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var order = map[token.Token]int{
	token.IMPORT: 0,
	token.CONST:  1,
	token.VAR:    2,
	token.TYPE:   3,
	token.FUNC:   4,
}
