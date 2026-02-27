package functions

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func SimpleLS(w io.Writer, args []string, useColor bool) {

	if len(args) == 0 {
		args = []string{"."}	//if user do not input any, the args will be "."
	}

	var files []string
	var dirs []string

	for _, path := range args {			//for every path from user input args
		info, err := os.Lstat(path)		//use system function to obtain the info of this path, return file info and error boolean value
		if  err != nil {
			fmt.Fprintf(os.Stderr, "gols: cannot to access '%s' : %v\n", path , err)
			continue
		}

		if info.IsDir() {				//use library method to determine Directory or File
			dirs = append(dirs, path)		//add into correspond slice
		} else {
			files = append(files, path)
		}
	}

	sort.Strings(files)			//sort by lexicographical order
	sort.Strings(dirs)

	for _, f := range files {		//for evey element in the files slice
		info, err := os.Lstat(f)		//obtain file info and error condition
		if err != nil {
			fmt.Fprintf(os.Stderr, "gols: cannot access '%s' : %v\n", f, err)
			continue
		}
		ColorPrint(w, f, info, useColor)		//use function ColorPrint to determine whether requires color
	}

	for i, d := range dirs {		//for loop go through the directory slice
		if len(dirs) > 1 {			//if there are more than one directory
			fmt.Fprintf(w, "%s:\n", d)		//write directory name d to writer w
		}

		entries, err := os.ReadDir(d)		//obtain all files in the directory
		if err != nil {
			fmt.Fprintf(os.Stderr, "gols: cannot access '%s': %v\n", d, err)
			continue
		}

		entries = DirFilter(entries)		//filter the hidden files, remove .hidden and .git

		var names []string					//only save the file name
		for _, e := range entries {
			names = append(names, e.Name())
		}

		sort.Strings(names)

		for _, name := range names {			//output the content from drectory
			fullPath := filepath.Join(d, name)

			info, err := os.Lstat(fullPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "gols: cannot access '%s': %v\n", fullPath, err)
				continue
			}

			ColorPrint(w, name, info, useColor)		//only output the name and use info to determine color
		}

		if i < len(dirs)-1 {		//if current directoey is not the last one, print a empty line
			fmt.Fprintln(w)
		}
	}
}