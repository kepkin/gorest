package pkg

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/scanner"
	"go/token"
	"io"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"golang.org/x/tools/go/ast/astutil"
)

func BuildTemplates() (res *template.Template, err error) {
	a := sprig.GenericFuncMap()
	res = template.New("").Funcs(a)
	res = res.Funcs(template.FuncMap{
		"MakeIdentifier":    MakeIdentifier,
		"GetNameFromRef":    GetNameFromRef,
		"ToConstructorType": ToConstructorType,
		"ConvertUrl":        ConvertUrl,
	})

	dir, err := Assets.Open("/")
	if err != nil {
		return
	}
	files, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	if len(files) == 0 {
		return res, fmt.Errorf("gorest: no template files")
	}

	for _, f := range files {
		fData, err := Assets.Open(f.Name())
		if err != nil {
			return res, err
		}
		data, err := ioutil.ReadAll(fData)
		if err != nil {
			return res, err
		}

		res, err = res.New(f.Name()).Parse(string(data))
		if err != nil {
			return res, err
		}
	}

	return
}

func GenerateFromFile(swaggerPath string, packageName string, wr io.Writer) error {
	content, err := ioutil.ReadFile(swaggerPath)
	if err != nil {
		return err
	}

	spec, err := readSpec(content)
	if err != nil {
		return err
	}

	return GenerateFromSpec(spec, packageName, wr)
}

func GenerateFromSpec(spec spec, packageName string, wr io.Writer) error {
	t, err := BuildTemplates()
	if err != nil {
		return err
	}

	specMeta, err := processSpec(&spec)
	if err != nil {
		return err
	}

	result := new(strings.Builder)
	err = t.ExecuteTemplate(result, "main.tmpl", map[string]interface{}{
		"package": packageName,
		"spec":    spec,
		"meta":    specMeta,
	})
	if err != nil {
		return err
	}

	return finalizeGoSource(result.String(), wr)
}

// finalizeGoSource removes unneeded imports from the given Go source file and
// runs go fmt on it.
func finalizeGoSource(content string, wr io.Writer) error {
	// Make sure file parses and print content if it does not.
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", content, parser.ParseComments)
	if err != nil {
		var buf bytes.Buffer
		scanner.PrintError(&buf, err)
		return fmt.Errorf("%s\n========\nContent:\n%s", buf.String(), content)
	}

	// Clean unused imports
	imps := astutil.Imports(fset, file)
	for _, group := range imps {
		for _, imp := range group {
			path := strings.Trim(imp.Path.Value, `"`)
			if !astutil.UsesImport(file, path) {
				if imp.Name != nil {
					astutil.DeleteNamedImport(fset, file, imp.Name.Name, path)
				} else {
					astutil.DeleteImport(fset, file, path)
				}
			}
		}
	}
	ast.SortImports(fset, file)
	if err := format.Node(wr, fset, file); err != nil {
		return err
	}

	return nil
}

//go:generate go run -tags=dev assets_generate.go -source="github.com/kepkin/gorest/pkg".Assets
