package main

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"slices"
	"testing"
)

func TestGenerateDifferentSizes(t *testing.T) {
	testCases := []struct {
		name   string
		length int
	}{
		{"8 bytes", 8},
		{"16 bytes", 16},
		{"32 bytes", 32},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			rawKeyGenerated := make([]byte, testCase.length)
			if _, err := rand.Read(rawKeyGenerated); err != nil {
				t.Fatalf("couldn't generate key, got %s and err: %v", rawKeyGenerated, err)
			}
			if len(rawKeyGenerated) != testCase.length {
				t.Fatalf("length mismatch, got %s with length of %d", rawKeyGenerated, len(rawKeyGenerated))
			}
		})
	}
}

func TestGenerateUniqueness(t *testing.T) {
	testCases := []struct {
		name       string
		encodeFunc func([]byte) string
	}{
		{
			name: "Hex",
			encodeFunc: func(b []byte) string {
				return fmt.Sprintf("%x", b)
			},
		},
		{
			name: "SafeBase64",
			encodeFunc: func(b []byte) string {
				return base64.RawURLEncoding.EncodeToString(b)
			},
		},
		{
			name: "Base64",
			encodeFunc: func(b []byte) string {
				return base64.StdEncoding.EncodeToString(b)
			},
		},
		{
			name: "Base32",
			encodeFunc: func(b []byte) string {
				return base32.StdEncoding.EncodeToString(b)
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			generatedKeys := make([]string, 0, 10)
			for i := range 1000 {
				rawKeyGenerated := make([]byte, 32)
				if _, err := rand.Read(rawKeyGenerated); err != nil {
					t.Fatalf("Couldn't generate key %v", err)
				}
				if len(rawKeyGenerated) != 32 {
					t.Errorf("Length of string generated is not 32, got %s with length of: %d", rawKeyGenerated, len(rawKeyGenerated))
				}
				generatedKeys = append(generatedKeys, testCase.encodeFunc(rawKeyGenerated))
				if len(generatedKeys[i]) == 0 {
					t.Fatalf("got empty key")
				}
			}
			slices.Sort(generatedKeys)

			for i := 1; i < len(generatedKeys); i++ {
				if generatedKeys[i] == generatedKeys[i-1] {
					t.Errorf("got two identical keys in 10-length burst: %s, %s", generatedKeys[i], generatedKeys[i-1])
				}
			}
		})
	}
}
