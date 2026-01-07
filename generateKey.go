package main

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

func main() {
	bytes8 := flag.Bool("8", false, "generate key 8 bytes long")
	bytes16 := flag.Bool("16", false, "generate key 16 bytes long")

	toBase32 := flag.Bool("base32", false, "turn key generated into base32 encoding")
	toUrlSafeBase64 := flag.Bool("ubase64", false, "turn key generated into url-safe base64 encoding")
	toBase64 := flag.Bool("base64", false, "turn key generated into base64 encoding")

	version := flag.Bool("version", false, "display version number")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\t--Usage of KeyCrypt--\n")
		flag.PrintDefaults() // Prints the auto-generated flag help
		fmt.Fprintln(os.Stderr, "Defaults to 32-byte-long key if no flag is set.")
		fmt.Fprintln(os.Stderr, "The output format is hex representation of the key bytes if not encoding option is set")
		fmt.Fprintln(os.Stderr, "By default he program does not encode key generated to base32 or base64")
		fmt.Fprintln(os.Stderr, "If multiple flags provided either byte size or encoding, the first one in each category takes precedence")
		fmt.Fprintln(os.Stderr, `Couple with "| clip" on Windows (or "| pbcopy" on macOS, "| xclip -selection clipboard" on Linux) for adding it directly to clipboard.`)
		fmt.Fprintln(os.Stderr, `Examples:`)
		fmt.Fprintln(os.Stderr, `  keycrypt -16`)
		fmt.Fprintln(os.Stderr, `  keycrypt -16 -base64`)
		fmt.Fprintln(os.Stderr, `  keycrypt -16 | clip (Windows)`)
		fmt.Fprintln(os.Stderr, `  keycrypt | pbcopy (macOS)`)
	}
	flag.Parse()

	var keyLength int

	if *bytes8 {
		keyLength = 8
	} else if *bytes16 {
		keyLength = 16
	} else {
		keyLength = 32
	}

	if *version {
		fmt.Printf("\tKeycrypt :)\n\tVersion 3.0.0\n")
		os.Exit(0)
	}

	keyGenerated := make([]byte, keyLength)
	if _, err := rand.Read(keyGenerated); err != nil {
		fmt.Println("Coudn't generate key: ", err.Error())
		os.Exit(1)
	}
	var encodedKey string
	if *toBase32 || *toBase64 {
		if *toBase32 {
			encodedKey = base32.StdEncoding.EncodeToString(keyGenerated)
		} else if *toBase64 {
			encodedKey = base64.StdEncoding.EncodeToString(keyGenerated)
		} else if *toUrlSafeBase64 {
			encodedKey = base64.URLEncoding.EncodeToString(keyGenerated)
		}
		fmt.Printf("%s", encodedKey)
	} else {
		fmt.Printf("%x", keyGenerated)
	}
}
