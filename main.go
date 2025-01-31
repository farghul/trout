package main

import (
	"fmt"
	"log"
	"os"
)

// Launch the program and execute according to the supplied flag
func main() {
	switch os.Args[1] {
	case "-h", "--help":
		help()
	case "-r", "--run":
		serialize()
		trout = compiler()
		proceed(trout)
		changedir()
		prepare()
		release = os.Args[2]
		packagist()
		journal("Branch " + branch + release + " for Production release " + release + " created.")
	case "-v", "--version":
		version()
	default:
		alert("Unknown argument(s) -")
	}
}

// Record a message to a log file
func journal(message string) {
	file, err := os.OpenFile(assets+"logs/trout.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	inspect(err)
	log.SetOutput(file)
	log.Println(message)
}

// Print a colourized error message
func alert(message string) {
	fmt.Println(red, message, halt, reset)
	os.Exit(0)
}

// Print program version number
func version() {
	fmt.Println(yellow+"Trout", green+bv, reset)
}

// Print help information for using the program
func help() {
	fmt.Println(yellow, "\nUsage:", reset)
	fmt.Println("  [program] [flag] [release name or number]")
	fmt.Println(yellow, "\nOptions:")
	fmt.Println(green, " -h, --help", reset, "		Help Information")
	fmt.Println(green, " -r, --run", reset, "	    Run Program")
	fmt.Println(green, " -v, --version", reset, "	Display Program Version")
	fmt.Println(yellow, "\nExample:", reset)
	fmt.Println(green, "   trout -r 88")
	fmt.Println(yellow, "\nHelp:", reset)
	fmt.Println("  For more information go to:")
	fmt.Println(green, "   https://github.com/farghul/trout.git")
	fmt.Println(reset)
}
