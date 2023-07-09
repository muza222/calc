package main

// check all must have files

// check .type file

func CheckTypeFile(name string) int {
	lastIndex := -1
	for item := len(name) - 1; item >= 0; item-- {
		if name[item] == '.' {
			lastIndex = item
			break
		}
	}

	return lastIndex
}

// all Check

func GeneralCheck(name string) bool {
	numIndex := CheckTypeFile(name)
	if numIndex == -1 && checkingExtensions > 0 {
		return false
	}
	if numIndex != -1 && checkingExtensions > 0 {
		checkThisExtension := name[numIndex:]
		if extensions[checkThisExtension]&checkingExtensions != checkingExtensions {
			return false
		}
	}
	excludeCheck := CheckExclude(name)
	if !excludeCheck {
		return false
	}
	lenRestrict := len(restrict)
	restrictCheck := CheckRestrict(name)
	if !restrictCheck {
		return false
	}
	return lenRestrict == 0

}
