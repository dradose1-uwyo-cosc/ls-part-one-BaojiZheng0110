package functions

import(
	"fmt"
	"os"
	"io"
)

const (
	Blue = "\033[34m"		//defines constants for color
	Green = "\033[32m"
    Reset = "\033[0m"
)

func ColorPrint(w io.Writer, name string, info os.FileInfo, useColor bool) {
	if !useColor {
		fmt.Fprintln(w, name)		//if logicaly ColorPrint is false, write to writer w and change line
		return
	}

	mode := info.Mode()

	if info.IsDir() {
		fmt.Fprintln(w, Blue+name+Reset)	//check if is directory, print out with blue color
	} else if mode.IsRegular() && (mode&0111) != 0 {
		fmt.Fprintln(w, Green+name+Reset)	//if is regular file and has execute permission, print out with green color
	} else {
		fmt.Fprintln(w, name)		//print normally
	}
}
