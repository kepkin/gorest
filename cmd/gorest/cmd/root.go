package cmd

import (
	"log"

	"github.com/kepkin/gorest"
	"github.com/spf13/cobra"
)

var pkgName string
var swaggerFile string

var rootCmd = &cobra.Command{
	Use: "gorest",
	Run: func(cmd *cobra.Command, args []string) {
		err := gorest.Generate(swaggerFile, gorest.Options{
			PackageName: pkgName,
			TargetFile:  args[0],
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.Flags().StringVar(&pkgName, "pkg", "", "path to swagger.yaml")
	rootCmd.Flags().StringVar(&swaggerFile, "swagger", "", "path to swagger.yaml")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
