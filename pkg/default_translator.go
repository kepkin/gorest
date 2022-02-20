package pkg

import (
	"github.com/kepkin/gorest/internal/generator"
	"github.com/kepkin/gorest/internal/translator"
)

func MakeTranslator() (res translator.Translator) {
	res.RefResolver = generator.GetNameFromRef
	res.MakeIdentifier = generator.MakeIdentifier
	res.MakeTitledIdentifier = generator.MakeTitledIdentifier

	res.RegisterField(translator.ObjectFieldConstructor{})
	res.RegisterField(translator.BooleanFieldConstructor{})
	res.RegisterField(translator.IntegerFieldConstructor{})
	res.RegisterField(translator.StringFieldConstructor{})
	res.RegisterField(translator.ArrayFieldConstructor{})
	res.RegisterField(translator.DecimalFieldConstructor{})

	return
}
