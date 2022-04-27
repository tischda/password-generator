package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const PROG_NAME string = "password-generator"
const DEFAULT_LENGTH int = 20

var version string

var flag_help = flag.Bool("help", false, "displays this help message")
var flag_version = flag.Bool("version", false, "print version and exit")
var flag_length = flag.Int("length", DEFAULT_LENGTH, fmt.Sprintf("%s%d", "password length between 1 and ", MAX_LENGTH))
var flag_gui = flag.Bool("gui", false, "show GUI")

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [option]\n", os.Args[0])
		fmt.Fprint(os.Stderr, "\n OPTIONS:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	log.SetFlags(0)

	flag.Parse()
	if flag.Arg(0) == "version" || *flag_version {
		fmt.Printf("%s version %s\n", PROG_NAME, version)
		return
	}
	if *flag_help {
		flag.Usage()
	}
	if *flag_gui {
		processGUI()
	} else {
		processCLI()
	}
}

func processCLI() {
	err := checkLength(*flag_length)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	} else {
		fmt.Println(GeneratePassword(*flag_length))
	}
}
