package cmd

import (
	"log"

	"github.com/kepkin/gorest"
	"github.com/spf13/cobra"
)

var pkgName string
var swaggerFile string
var outFile string

var rootCmd = &cobra.Command{
	Use: "gorest",
	Run: func(cmd *cobra.Command, args []string) {
		err := gorest.Generate(swaggerFile, gorest.Options{
			PackageName: pkgName,
			TargetFile:  outFile,
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.Flags().StringVar(&pkgName, "pkg", "api", "path to swagger.yaml")
	rootCmd.Flags().StringVar(&swaggerFile, "swagger", "", "path to swagger.yaml")
	rootCmd.Flags().StringVar(&outFile, "out", "", "path to output file (default stdin)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
