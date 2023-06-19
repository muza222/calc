package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var app = &cobra.Command{
	Use:   "app",
	Short: "gitfame",
	Long:  "calculate stats for git repo",
	Run:   run,
}
var flagExtensions []string

func init() {
	app.Flags().StringVar(&path, "repository", ".", "path to repo")
	app.Flags().StringVar(&format, "format", "tabular", "output type")
	app.Flags().StringVar(&revision, "revision", "HEAD", "revision")
	app.Flags().StringVar(&sortBy, "order-by", "lines", "sort")
	app.Flags().BoolVar(&committer, "use-committer", false, "attribute files")
	app.Flags().StringSliceVar(&flagExtensions, "extensions", []string{}, "only count files with these extensions")
	app.Flags().StringSliceVar(&languages, "languages", []string{}, "count of special type e.x. '.cpp, .md, /h'")
	app.Flags().StringSliceVar(&exclude, "exclude", []string{}, "exclude these patterns")
	app.Flags().StringSliceVar(&restrict, "restrict-to", []string{}, "restrict to these patterns")
}

func main() {
	err := app.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
