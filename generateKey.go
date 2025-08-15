package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"os"
)

func main() {
	bytes8 := flag.Bool("8", false, "generate key 8 bytes long")
	bytes16 := flag.Bool("16", false, "generate key 16 bytes long")
	bytes32 := flag.Bool("32", false, "generate key 32 bytes long")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of KeyCrypt :)\n")
		flag.PrintDefaults() // Prints the auto-generated flag help
		fmt.Fprintln(os.Stderr, "Defaults to 32-byte-long key if no flag is provided.")
		fmt.Fprintln(os.Stderr, "If multiple flags provided, the first one takes priority")
		fmt.Fprintln(os.Stderr, `Couple with "| clip" on Windows (or "| pbcopy" on macOS, "| xclip -selection clipboard" on Linux) for adding it directly to clipboard.`)
		fmt.Fprintln(os.Stderr, `Examples:`)
		fmt.Fprintln(os.Stderr, `  cryptkey -16`)
		fmt.Fprintln(os.Stderr, `  cryptkey -16 | clip (Windows)`)
		fmt.Fprintln(os.Stderr, `  cryptkey | pbcopy (macOS)`)
	}
	flag.Parse()

	var keyLength int

	if *bytes8 {
		keyLength = 8
	} else if *bytes16 {
		keyLength = 16
	} else if *bytes32 {
		keyLength = 32
	} else {
		keyLength = 32
	}

	keyGenerated := make([]byte, keyLength)
	if _, err := rand.Read(keyGenerated); err != nil {
		fmt.Println("Coudn't generate key: ", err.Error())
		os.Exit(1)
	}
	fmt.Printf("%x", keyGenerated)
	// Run paired with | clip in windows or equivalent to own OS for getting to clipboard
}
