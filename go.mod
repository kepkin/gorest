module github.com/kepkin/gorest

go 1.16

require (
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.4.2 // indirect
	github.com/Masterminds/sprig v2.20.0+incompatible
	github.com/ghodss/yaml v1.0.0
	github.com/gin-gonic/gin v1.5.0
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.2.0
	github.com/huandu/xstrings v1.2.0 // indirect
	github.com/iancoleman/strcase v0.0.0-20190422225806-e506e3ef7365
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/shopspring/decimal v0.0.0-20200227202807-02e2044944cc
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/ugorji/go v1.1.7 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect; v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	golang.org/x/tools v0.0.0-20191125011157-cc15fab314e3 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.2
//gopkg.in/yaml.v2 v2.2.4
)

// https://github.com/gin-gonic/gin/issues/1673
replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
