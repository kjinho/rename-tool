package cmd

import "strings"

func genNewFilename(filename string, extension string) string {
	var fullNewExtension string
	var fullNewFilename string
	if extension[0] == '.' {
		fullNewExtension = extension
	} else {
		fullNewExtension = "." + extension
	}

	if len(filename) < 5 { // too short to have a three-character extension
		fullNewFilename = filename + fullNewExtension
	} else {
		curExtension := filename[len(filename)-4:]
		if curExtension[0] == '.' && strings.Count(curExtension, ".") == 1 {
			fullNewFilename = filename[0:len(filename)-4] + fullNewExtension
		} else {
			fullNewFilename = filename + fullNewExtension
		}
	}
	return (fullNewFilename)
}
