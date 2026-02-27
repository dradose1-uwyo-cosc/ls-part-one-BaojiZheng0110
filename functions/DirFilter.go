package functions

import (
	"os"
	"strings"
)

func DirFilter(entries []os.DirEntry) []os.DirEntry {		//this function takes a slice and return a slice
	var result []os.DirEntry		//declares aslice variable which is null slice

	for _, e := range entries {
		if !strings.HasPrefix(e.Name(), ".") {		//check if name doesn't start with ".", which means ignore hidden files
			result = append(result, e)				//append element to slice
		}
	}
	return result
}