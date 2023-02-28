package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
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
	fileName := c.Args().First()
	fileContents, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	parsedFile, err := parser.ParseFile(
		token.NewFileSet(),
		"", fileContents,
		parser.AllErrors,
	)
	if err != nil {
		return err
	}

	sort.Slice(parsedFile.Decls, func(i, j int) bool {
		return declOrder(parsedFile.Decls[i]) < declOrder(parsedFile.Decls[j])
	})
	orderedFileContents := prettyPrint(parsedFile, fileContents)

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, orderedFileContents)
	return err
}

func prettyPrint(ordered *ast.File, initial []byte) *bytes.Buffer {
	w := new(bytes.Buffer)

	fmt.Fprintf(w, "package %s\n\n", ordered.Name)
	for _, decl := range ordered.Decls {
		w.Write(initial[decl.Pos()-1 : decl.End()])
		w.WriteByte('\n')
	}

	return w
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
