package main

import (
	"fmt"
	"os"
)

// Launch the program and execute according to the supplied flag
func main() {
	logo()
	switch os.Args[1] {
	case "-h", "--help":
		help()
	case "-r", "--run":
		credits()
		serialize()
		trout = compiler()
		proceed(trout)
		changedir()
		prepare()
		release = os.Args[2]
		packagist()
		journal("Branch " + branch + release + " for Production release " + release + " created.")
	case "-v", "--version":
	default:
		alert("Unknown argument(s) -")
	}
}

// Provide and highlight an informational message
func inform(message string) {
	Yellow.Printf("%s", "** ")
	fmt.Print(message)
	Yellow.Println(" **")
}

// Print a colourized error message
func alert(message string) {
	Red.Printf("\n%s", "Error: ")
	fmt.Printf("%s", message)
	BGRed.Println(halt)
	inform("Use -h to display help information")
	os.Exit(0)
}

// Print help information for using the program
func help() {
	Yellow.Println("\nUsage:")
	fmt.Println("  [program] [flag] [release name or number]")
	Yellow.Println("\nOperational Flags:")
	Green.Printf("%s", "  -h, --help")
	fmt.Println("		Help Information")
	Green.Printf("%s", "  -r, --run")
	fmt.Println("		Run Program")
	Green.Printf("%s", "  -v, --version")
	fmt.Println("		Display Program Version")
	Yellow.Println("\nExample:")
	fmt.Println("  Adding your path to file if necessary, run:")
	Green.Printf("%s", "    trout -r 88")
	Yellow.Println("\nHelp:")
	fmt.Println("  For more information go to:")
	Green.Println("    https://github.com/farghul/trout.git")
}

// Print a colourized logo, indicating the program is running
func logo() {
	Green.Println("▗▄▄▄▖▗▄▄▖  ▗▄▖ ▗▖ ▗▖▗▄▄▄▖")
	Green.Println("  █  ▐▌ ▐▌▐▌ ▐▌▐▌ ▐▌  █  ")
	Green.Println("  █  ▐▛▀▚▖▐▌ ▐▌▐▌ ▐▌  █  ")
	Green.Println("  █  ▐▌ ▐▌▝▚▄▞▘▝▚▄▞▘  █  ")
	Green.Println(bv)
}

// Print the program mission statement and creator credit
func credits() {
	fmt.Println("\nA `Release to Production` install tool for WordPress plugins")
	fmt.Println("Created by Byron Stuike")
}
