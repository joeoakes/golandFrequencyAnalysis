package main

import (
	"fmt"
	"strings"
)

// Expected letter frequencies in English
var letterFrequencies = map[rune]float64{
	'a': 8.2, 'b': 1.5, 'c': 2.8, 'd': 4.3, 'e': 12.7, 'f': 2.2, 'g': 2.0, 'h': 6.1,
	'i': 7.0, 'j': 0.2, 'k': 0.8, 'l': 4.0, 'm': 2.4, 'n': 6.7, 'o': 7.5, 'p': 1.9,
	'q': 0.1, 'r': 6.0, 's': 6.3, 't': 9.1, 'u': 2.8, 'v': 1.0, 'w': 2.4, 'x': 0.2,
	'y': 2.0, 'z': 0.1,
}

func main() {
	ciphertext := "VQREQFGT" // Replace with your ciphertext
	bestShift := findBestShift(ciphertext)
	decryptedText := caesarDecrypt(ciphertext, bestShift)
	fmt.Printf("Ciphertext: %s\n", ciphertext)
	fmt.Printf("Best Shift: %d\n", bestShift)
	fmt.Printf("Decrypted Text: %s\n", decryptedText)
}

func caesarDecrypt(ciphertext string, shift int) string {
	var decrypted strings.Builder
	for _, char := range ciphertext {
		if 'A' <= char && char <= 'Z' {
			decryptedChar := 'A' + (char-'A'-rune(shift)+26)%26
			decrypted.WriteRune(decryptedChar)
		} else if 'a' <= char && char <= 'z' {
			decryptedChar := 'a' + (char-'a'-rune(shift)+26)%26
			decrypted.WriteRune(decryptedChar)
		} else {
			// Preserve non-alphabetic characters as is
			decrypted.WriteRune(char)
		}
	}
	return decrypted.String()
}

func calculateScore(text string) float64 {
	text = strings.ToLower(text)
	score := 0.0
	for _, char := range text {
		if freq, ok := letterFrequencies[char]; ok {
			score += freq
		}
	}
	return score
}

func findBestShift(ciphertext string) int {
	bestShift := 0
	bestScore := 0.0
	for shift := 1; shift <= 25; shift++ {
		decryptedText := caesarDecrypt(ciphertext, shift)
		score := calculateScore(decryptedText)
		if score > bestScore {
			bestScore = score
			bestShift = shift
		}
	}
	return bestShift
}
