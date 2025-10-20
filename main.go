package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

const DEFAULT_LENGTH int = 20
const MAX_LENGTH int = 128

// https://goreleaser.com/cookbooks/using-main.version/
var (
	name    string
	version string
	date    string
	commit  string
)

// flags
type Config struct {
	gui 	bool
	length 	int
	help    bool
	version bool
}

// initFlags initializes and parses command-line flags.
// It returns a pointer to a Config struct with the parsed values.
func initFlags() *Config {
	cfg := &Config{}
	flag.IntVar(&cfg.length, "l", DEFAULT_LENGTH, "")
	flag.IntVar(&cfg.length, "length", DEFAULT_LENGTH, fmt.Sprintf("%s%d", "password length between 1 and ", MAX_LENGTH))
	flag.BoolVar(&cfg.gui, "g", false, "")
	flag.BoolVar(&cfg.gui, "gui", false, "show GUI")
	flag.BoolVar(&cfg.help, "?", false, "")
	flag.BoolVar(&cfg.help, "help", false, "displays this help message")
	flag.BoolVar(&cfg.version, "v", false, "")
	flag.BoolVar(&cfg.version, "version", false, "print version and exit")
	return cfg
}

// main is the entry point of the application.
func main() {
	log.SetFlags(0)
	cfg := initFlags()
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: "+name+` [OPTIONS]

Generates a random password with a default length of `+fmt.Sprintf("%d", DEFAULT_LENGTH)+` characters.

OPTIONS:

  -l, --length
        password length between 1 and `+fmt.Sprintf("%d", MAX_LENGTH)+`
  -g, --gui
        show GUI
  -?, --help
        display this help message
  -v, --version
        print version and exit

EXAMPLES:`)

		fmt.Fprintln(os.Stderr, "\n  $ "+name+` --length 15
  dtB{#hP_C)C^D4a`)
	}
	flag.Parse()

	if flag.Arg(0) == "version" || cfg.version {
		fmt.Printf("%s %s, built on %s (commit: %s)\n", name, version, date, commit)
		return
	}

	if cfg.help {
		flag.Usage()
		return
	}

	if flag.NArg() > 0 {
		flag.Usage()
		os.Exit(1)
	}

	if cfg.gui {
		processGUI(cfg.length)
	} else {
		processCLI(cfg.length)
	}
}

// processCLI handles the command-line interface logic.
// It generates a password of the given length and prints it to stdout.
// The length parameter specifies the desired password length.
func processCLI(length int) {
	err := checkLength(length)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	} else {
		fmt.Println(GeneratePassword(length))
	}
}

// checkLength validates if the given password length is within the allowed range.
// The len parameter is the length to validate.
// It returns an error if the length is out of bounds, otherwise nil.
func checkLength(len int) error {
	if len < 0 || len > MAX_LENGTH {
		return fmt.Errorf("password length must be between 1 and %d", MAX_LENGTH)
	}
	return nil
}

// lengthValidator is a validator for GUI input. It checks if the provided string
// represents a valid password length.
// The s parameter is the string to validate.
// It returns an error if the string is not a valid integer or if the length is
// out of bounds, otherwise nil.
func lengthValidator(s string) error {
	len, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	return checkLength(len)
}

