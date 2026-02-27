package main

import (
	"os"
	"gols/functions"
)

func main() {
	args := os.Args[1:]				//get command-line arguments, skip program name
	useColor := functions.IsTerminal(os.Stdout)		//check if output is a terminal, return true if print to terminal
	functions.SimpleLS(os.Stdout, args, useColor)		//calls SimpleLS function
}