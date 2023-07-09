package main

import (
	"sort"
)

// createDefaultPerson

type User struct {
	commits map[string]struct{}
	files   map[string]struct{}
	lines   int
}

// these Persons for sorting

var groupUser = make(map[string]User)

var sortBy string

// createLastTypePerson

type UserFinal struct {
	Name    string `json:"name"`
	Lines   int    `json:"lines"`
	Commits int    `json:"commits"`
	Files   int    `json:"files"`
}

var groupFinalUser []UserFinal

// createNewUser (construct)

func createUser() User {
	commits := make(map[string]struct{})
	files := make(map[string]struct{})
	lines := 0
	thisUser := User{commits: commits, files: files, lines: lines}

	return thisUser
}

// sorting by ....

func sortByLines(i int, j int) bool {
	if groupFinalUser[i].Lines != groupFinalUser[j].Lines {
		return groupFinalUser[i].Lines > groupFinalUser[j].Lines
	}
	if groupFinalUser[i].Commits != groupFinalUser[j].Commits {
		return groupFinalUser[i].Commits > groupFinalUser[j].Commits
	}
	if groupFinalUser[i].Files != groupFinalUser[j].Files {
		return groupFinalUser[i].Files > groupFinalUser[j].Files
	}
	return groupFinalUser[i].Name < groupFinalUser[j].Name
}

// commits

func sortByCommits(i, j int) bool {
	if groupFinalUser[i].Commits != groupFinalUser[j].Commits {
		return groupFinalUser[i].Commits > groupFinalUser[j].Commits
	}
	if groupFinalUser[i].Lines != groupFinalUser[j].Lines {
		return groupFinalUser[i].Lines > groupFinalUser[j].Lines
	}
	if groupFinalUser[i].Files != groupFinalUser[j].Files {
		return groupFinalUser[i].Files > groupFinalUser[j].Files
	}
	return groupFinalUser[i].Name < groupFinalUser[j].Name
}

func sortByFiles(i, j int) bool {
	if groupFinalUser[i].Files != groupFinalUser[j].Files {
		return groupFinalUser[i].Files > groupFinalUser[j].Files
	}
	if groupFinalUser[i].Lines != groupFinalUser[j].Lines {
		return groupFinalUser[i].Lines > groupFinalUser[j].Lines
	}
	if groupFinalUser[i].Files != groupFinalUser[j].Files {
		return groupFinalUser[i].Files > groupFinalUser[j].Files
	}
	return groupFinalUser[i].Name < groupFinalUser[j].Name
}

// thisSorting

func sortByElem() {
	for name, data := range groupUser {
		groupFinalUser = append(groupFinalUser, UserFinal{
			Name:    name,
			Commits: len(data.commits),
			Files:   len(data.files),
			Lines:   data.lines})
	}
	switch sortBy {
	case "lines":
		sort.Slice(groupFinalUser, sortByLines)
	case "commits":
		sort.Slice(groupFinalUser, sortByCommits)
	case "files":
		sort.Slice(groupFinalUser, sortByFiles)
	default:
		attention()
		break
	}
}
