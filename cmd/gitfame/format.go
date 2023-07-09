package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

type stats interface {
	write(uf UserFinal)
	flush()
}

/////////////

type tabWriter struct {
	*tabwriter.Writer
}

func (w tabWriter) write(obj UserFinal) {
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", obj.Name, obj.Lines, obj.Commits, obj.Files)
}
func (w tabWriter) flush() {
	_ = w.Flush()
}

func formatTabular() tabWriter {
	tab := new(tabwriter.Writer)
	tab.Init(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(tab, "Name\tLines\tCommits\tFiles")
	res := tabWriter{tab}
	return res
}

////////////////

type csvWriter struct {
	*csv.Writer
}

func (w csvWriter) write(uf UserFinal) {
	_ = w.Write([]string{
		uf.Name,
		strconv.Itoa(uf.Lines),
		strconv.Itoa(uf.Commits),
		strconv.Itoa(uf.Files)})
}

func (w csvWriter) flush() {
	w.Flush()
}

func formatCSV() csvWriter {
	w := csv.NewWriter(os.Stdout)
	headers := []string{"Name", "Lines", "Commits", "Files"}
	_ = w.Write(headers)

	res := csvWriter{w}
	return res
}

/////////////////////////

type jsonWriter struct {
}

func (_ jsonWriter) write(entry UserFinal) {
	s, _ := json.Marshal(entry)
	res := string(s)
	fmt.Println(res)
}
func (_ jsonWriter) flush() {}

////////////////////
