package barber

import (
	"bytes"
	"errors"
	"go/format"
	"go/parser"
	"go/scanner"
	"go/token"
	"io"
)

func PrettifySource(src io.Reader, dst io.Writer) error {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, parser.ParseComments|parser.AllErrors)
	if err != nil {
		var errBuf bytes.Buffer
		scanner.PrintError(&errBuf, err)
		return errors.New(errBuf.String())
	}

	if err = format.Node(dst, fset, file); err != nil {
		return err
	}
	return nil
}
