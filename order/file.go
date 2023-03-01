package order

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"sort"
)

type File struct {
	pkg *ast.File
	raw []byte
}

func ReadFile(name string) (*File, error) {
	fileContents, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	parsedFile, err := parser.ParseFile(
		token.NewFileSet(),
		"", fileContents,
		parser.ParseComments|parser.AllErrors,
	)
	if err != nil {
		return nil, err
	}
	return NewFile(fileContents, parsedFile), nil
}

func NewFile(raw []byte, pkg *ast.File) *File {
	sort.Slice(pkg.Decls, func(i, j int) bool {
		return declOrder(pkg.Decls[i]) < declOrder(pkg.Decls[j])
	})
	return &File{
		pkg: pkg,
		raw: raw,
	}
}

func (file *File) Pretty() *bytes.Buffer {
	w := &bytes.Buffer{}

	if file.pkg.Doc != nil {
		for _, each := range file.pkg.Doc.List {
			w.WriteString(each.Text + "\n")
		}
	}
	fmt.Fprintf(w, "package %s\n\n", file.pkg.Name)

	for i, decl := range file.pkg.Decls {
		from, to := declPos(decl)
		w.Write(file.raw[from:to])
		if i < len(file.pkg.Decls)-1 {
			w.WriteString("\n\n")
		} else {
			w.WriteByte('\n') //* Only 1 newline after the last declaration.
		}
	}

	return w
}
