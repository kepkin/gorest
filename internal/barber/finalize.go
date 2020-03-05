package barber

import (
	"bytes"
	"errors"
	fmt "fmt"
	"go/format"
	"go/parser"
	"go/scanner"
	"go/token"
	"io"
	"io/ioutil"
	"strings"

	//"golang.org/x/tools/go/ast/astutil"
)

func PrettifySource(src io.Reader, dst io.Writer) error {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, parser.ParseComments|parser.AllErrors)
	if err != nil {
		f, _ := ioutil.ReadAll(src)

		var res strings.Builder
		for i, codeLine := range strings.Split(string(f), "\n") { // TODO(a.telyshev): bufio.Scanner()
			res.WriteString(fmt.Sprintf("%-6d %s\n", i, codeLine))
		}

		var errBuf bytes.Buffer
		scanner.PrintError(&errBuf, err)

		return errors.New(errBuf.String() + "\nGENERATED FILE:\n" + res.String())
	}

	// TODO(a.telyshev): Save info about used imports in generator.Generator?
	//cleanUnusedImports(fset, file)

	if err = format.Node(dst, fset, file); err != nil {
		return err
	}
	return nil
}

//func cleanUnusedImports(fset *token.FileSet, file *ast.File) {
//	imps := astutil.Imports(fset, file)
//	for _, group := range imps {
//		for _, imp := range group {
//			path := strings.Trim(imp.Path.Value, `"`)
//			if !astutil.UsesImport(file, path) {
//				if imp.Name != nil {
//					astutil.DeleteNamedImport(fset, file, imp.Name.Name, path)
//				} else {
//					astutil.DeleteImport(fset, file, path)
//				}
//			}
//		}
//	}
//	ast.SortImports(fset, file)
//}
