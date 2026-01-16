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
	version := flag.Bool("version", false, "display version number")

	bytes8 := flag.Bool("8", false, "generate key 8 bytes long")
	bytes16 := flag.Bool("16", false, "generate key 16 bytes long")

	toBase32 := flag.Bool("base32", false, "turn key generated into base32 encoding")
	toBase64 := flag.Bool("base64", false, "turn key generated into base64 encoding")
	toUrlSafeBase64 := flag.Bool("ubase64", false, "turn key generated into url-safe base64 encoding")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\t--Usage of KeyCrypt--\n")
		flag.PrintDefaults() // Prints the auto-generated flag help
		fmt.Fprintln(os.Stderr, "Defaults to 32-byte-long key if no flag is set.")
		fmt.Fprintln(os.Stderr, "Default to hex representation of the key bytes if not encoding option is set")
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

	if *version {
		fmt.Printf("\tKeycrypt :)\n\tVersion 3.0.0\n")
		os.Exit(0)
	}

	// Length options
	var keyLength int
	switch {
	case *bytes8:
		keyLength = 8
	case *bytes16:
		keyLength = 16
	default:
		keyLength = 32
	}

	// Key bytes generation
	rawKeyGenerated := make([]byte, keyLength)
	if _, err := rand.Read(rawKeyGenerated); err != nil {
		fmt.Println("Coudn't generate key: ", err.Error())
		os.Exit(1)
	}
	
	// Encoding options
	var encodedKey string
	switch {
	case *toBase32:
		encodedKey = base32.StdEncoding.EncodeToString(rawKeyGenerated)
		fmt.Printf("%s", encodedKey)
	case *toBase64:
		encodedKey = base64.StdEncoding.EncodeToString(rawKeyGenerated)
		fmt.Printf("%s", encodedKey)
	case *toUrlSafeBase64:
		encodedKey = base64.RawURLEncoding.EncodeToString(rawKeyGenerated)
		fmt.Printf("%s", encodedKey)
	default:
		fmt.Printf("%x", rawKeyGenerated)
	}
}
