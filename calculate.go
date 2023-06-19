package main

import (
	"strconv"
	"strings"
)

var author = "author"

func calcStats(pathFile string) {
	gb, _ := GitBlame(pathFile)
	if len(gb) == 0 {
		gl, _ := GitLog(pathFile)
		fileLog := string(gl)
		checkFiles := strings.Split(strings.Split(fileLog, "\n")[0], "***")
		checkCommit, checkName := checkFiles[0], checkFiles[1]
		var person User
		var check bool
		if person, check = groupUser[checkName]; !check {
			person = createUser()
		}
		person.files[pathFile] = struct{}{}
		person.commits[checkCommit] = struct{}{}
		groupUser[checkName] = person
	}
	blameFiles := strings.Split(string(gb), "\n")
	var thisCommit string
	var thisLines int
	bucketCommits := make(map[string]string)
	for item := 0; item < len(blameFiles); item++ {
		info := strings.Fields(blameFiles[item])
		if len(info) == 4 {
			thisCommit = info[0]
			thisLines, _ = strconv.Atoi(info[3])
			if name, check := bucketCommits[thisCommit]; check {
				personItem := groupUser[name]
				personItem.lines += thisLines
				groupUser[name] = personItem
			} else {
			parseFields:
				for {
					item++
					f := strings.Fields(blameFiles[item])
					thisCheck := f[0]
					switch thisCheck {
					case author:
						lengthSlice := len(thisCheck) + 1
						name := blameFiles[item][lengthSlice:]
						var person User
						var check bool
						if person, check = groupUser[name]; !check {
							person = createUser()
						}
						person.commits[thisCommit] = struct{}{}
						person.files[pathFile] = struct{}{}
						person.lines += thisLines
						groupUser[name] = person
						bucketCommits[thisCommit] = name
					case "filename":
						break parseFields
					}
				}
			}
		}
		item++
	}
}
