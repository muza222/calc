package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var languages []string

type Lang struct {
	Name      string
	Extension []string
}

var LangExtension = make(map[string][]string)

func init() {
	var l []Lang
	_ = json.Unmarshal(language, &l)
	for _, lang := range l {
		LangExtension[strings.ToLower(lang.Name)] = lang.Extension
	}
}

func GitLog(filePath string) ([]byte, error) {
	log := exec.Command("git", "log", "--format=%H***%an", revision, "--", filePath)
	log.Dir = path
	SecondOut, err := log.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(SecondOut))
		os.Exit(1)
	}
	return SecondOut, nil
}
