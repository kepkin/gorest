package barber

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"go/parser"
	"go/scanner"
	"go/token"
	"io"
	"io/ioutil"
	"strings"
)

func PrettifySource(src io.Reader, dst io.Writer) error {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, parser.ParseComments|parser.AllErrors)
	if err != nil {
		f, _ := ioutil.ReadAll(src)

		var res strings.Builder
		for i, codeLine := range strings.Split(string(f), "\n") {
			res.WriteString(fmt.Sprintf("%d %s\n", i, codeLine))
		}

		var errBuf bytes.Buffer
		scanner.PrintError(&errBuf, err)

		return errors.New(errBuf.String() + "\nGENERATED FILE:\n" + res.String())
	}

	if err = format.Node(dst, fset, file); err != nil {
		return err
	}
	return nil
}
