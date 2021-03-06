package main

import (
	"fmt"

	"github.com/ddosakura/gklang"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "sblocker",
		Short: "A CLI to embed files into project.",
		Long:  `Embed files into project (such as a Go executable project).`,
		Run: func(cmd *cobra.Command, args []string) {
			gklang.Log(gklang.LDebug, params)

			// set config (dev/pkg/comment/algorithm)
			// choose lang / load plugin
			loadPlugin()

			// test src/target/example with force
			cleanTarget()

			// generate code
			generate()
		},
	}
)

func execute() {
	rootCmd.Version = version
	rootCmd.Execute()
}

var (
	src     string
	target  string
	example string
	force   bool

	pkg     string
	comment string

	lang      string
	algorithm string
	mod       string

	params map[string]string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&src, "src", "s", "public", "The path of the source directory.")
	rootCmd.PersistentFlags().StringVarP(&target, "target", "t", "sblock", "The destination path of the generated package.")
	rootCmd.PersistentFlags().StringVarP(&example, "example", "e", "example", "The example path to use the generated code.")
	rootCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "Overwrite destination and example file if it already exists.")

	rootCmd.PersistentFlags().StringVarP(&pkg, "package", "p", "sblock", "Name of the generated package.")
	rootCmd.PersistentFlags().StringVarP(&comment, "comment", "c", "Package contains static assets.", "The package comment.")

	rootCmd.PersistentFlags().StringVarP(&lang, "lang", "l", "golang", "The language of generated code.")
	rootCmd.PersistentFlags().StringVarP(&algorithm, "algorithm", "a", "default", "The compression algorithm to shrink the files.")
	rootCmd.PersistentFlags().StringVarP(&mod, "mod", "m", "", "The plugin of generater. This will disable lang and algorithm options!")

	rootCmd.PersistentFlags().StringToStringVarP(&params, "param", "v", make(map[string]string), "The custom parameters for plugin.")

	rootCmd.AddCommand(detailCmd)
}

var (
	detailCmd = &cobra.Command{
		Use:   "detail",
		Short: "The detail for plugin.",
		Long:  `Print the help msg in the plugin.`,
		Run: func(cmd *cobra.Command, args []string) {
			name := loadPlugin()
			println("Plugin:\n\t", name)
			println("Custom Parameters:")
			for k, v := range generater.DescriptionForParams() {
				fmt.Printf("\t%16s\t%s\n", k, v)
			}
		},
	}
)
