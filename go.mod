module github.com/kepkin/gorest

go 1.12

require (
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.4.2 // indirect
	github.com/Masterminds/sprig v2.20.0+incompatible
	github.com/gin-gonic/gin v1.4.0
	github.com/google/uuid v1.1.1 // indirect
	github.com/huandu/xstrings v1.2.0 // indirect
	github.com/iancoleman/strcase v0.0.0-20190422225806-e506e3ef7365
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/magiconair/properties v1.8.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pkg/errors v0.8.0
	github.com/shopspring/decimal v0.0.0-20191009025716-f1972eb1d1f5
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20181202132449-6a9ea43bcacd // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.3.0
	golang.org/x/tools v0.0.0-20190311212946-11955173bddd
	gopkg.in/yaml.v2 v2.2.2
)

// https://github.com/gin-gonic/gin/issues/1673
replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
