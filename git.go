package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var extensions = make(map[string]int)
var exclude []string
var restrict []string

// checking for global search files + extra

// some blames from ONE file

func GitBlame(pathFile string) ([]byte, error) {
	blameFiles := exec.Command("git", "blame", "--porcelain", revision, "--", pathFile)
	blameFiles.Dir = path
	ans, err := blameFiles.CombinedOutput()
	res := string(ans)
	if err != nil {
		fmt.Println(err)
		fmt.Println(res)
		os.Exit(1)

	}
	return ans, nil
}

func GitLsTree(revision string) ([]byte, error) {
	lsFiles := exec.Command("git", "ls-tree", revision, "-r")
	lsFiles.Dir = path
	ans, err := lsFiles.CombinedOutput()
	res := string(ans)
	if err != nil {
		fmt.Println(err)
		fmt.Println(res)
		os.Exit(1)
	}

	return ans, nil
}

// for file match

// exclude

func CheckExclude(name string) bool {
	for _, item := range exclude {
		if ok, _ := filepath.Match(item, name); !ok {
			return false
		}
	}
	return true
}

// restrict

func CheckRestrict(name string) bool {
	for _, item := range restrict {
		if ok, _ := filepath.Match(item, name); !ok {
			return false
		}
	}
	return true
}

func attention() {
	fmt.Println("unknown type of sorting!")
	os.Exit(1)
}
