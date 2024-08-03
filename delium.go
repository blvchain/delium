package delium

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

type Byte_hash struct {
	Byte_slice []byte
	String     string
}

// D256 processes a strData string by first hashing it with SHA-256, then repeatedly deleting
// characters from the resulting hash string at specified intervals, and finally hashing the
// modified string again with SHA-256.
//
// Parameters:
//   - strData: A string representing the strData phrase to be processed and hashed.
//   - deleteStep: An integer specifying the interval at which characters will be deleted from the hash string.
//   - repeat: An integer specifying how many times the deletion process should be applied.
//
// Returns:
//   - A pointer to a Byte_hash struct containing:
//   - Byte_slice: A byte slice of the final SHA-256 hash after applying the deletion process the specified
//     number of times.
//   - String: A hexadecimal string representation of the final SHA-256 hash.
//
// The function works as follows:
//  1. Computes the SHA-256 hash of the input strData string.
//  2. For a specified number of repetitions, deletes every `deleteStep`-th character from the hash string
//     generated in the previous step.
//  3. Computes the SHA-256 hash of the modified string after each repetition.
//  4. Returns a Byte_hash struct with the final SHA-256 hash in both byte slice and hexadecimal string formats.
//
// Example:
//
//	result := delium.D256("example strData", 3, 5)
//	fmt.Println(result.Byte_slice)  // prints the byte slice of the final hash
//	fmt.Println(result.String)      // prints the hexadecimal string of the final hash
func D256(strData string, deleteStep int, repeat int) *Byte_hash {

	strDataHash := sha256.Sum256([]byte(strData))
	var hashByte []byte = strDataHash[:]

	for i := 0; i < repeat; i++ {

		var result []rune
		for i, r := range string(hashByte) {
			if (i+1)%deleteStep != 0 {
				result = append(result, r)
			}
		}

		hashByte32 := sha256.Sum256([]byte(string(result)))
		hashByte = hashByte32[:]
	}

	return &Byte_hash{
		Byte_slice: hashByte,
		String:     hex.EncodeToString(hashByte),
	}
}

// D512 processes a strData string by first hashing it with SHA-512, then repeatedly deleting
// characters from the resulting hash string at specified intervals, and finally hashing the
// modified string again with SHA-512.
//
// Parameters:
//   - strData: A string representing the strData phrase to be processed and hashed.
//   - deleteStep: An integer specifying the interval at which characters will be deleted from the hash string.
//   - repeat: An integer specifying how many times the deletion process should be applied.
//
// Returns:
//   - A pointer to a Byte_hash struct containing:
//   - Byte_slice: A byte slice of the final SHA-512 hash after applying the deletion process the specified
//     number of times.
//   - String: A hexadecimal string representation of the final SHA-512 hash.
//
// The function works as follows:
//  1. Computes the SHA-512 hash of the input strData string.
//  2. For a specified number of repetitions, deletes every `deleteStep`-th character from the hash string
//     generated in the previous step.
//  3. Computes the SHA-512 hash of the modified string after each repetition.
//  4. Returns a Byte_hash struct with the final SHA-512 hash in both byte slice and hexadecimal string formats.
//
// Example:
//
//	result := delium.D512("example strData", 3, 5)
//	fmt.Println(result.Byte_slice)  // prints the byte slice of the final hash
//	fmt.Println(result.String)      // prints the hexadecimal string of the final hash
func D512(strData string, deleteStep int, repeat int) *Byte_hash {

	strDataHash := sha512.Sum512([]byte(strData))
	var hashByte []byte = strDataHash[:]

	for i := 0; i < repeat; i++ {

		var result []rune
		for i, r := range string(hashByte) {
			if (i+1)%deleteStep != 0 {
				result = append(result, r)
			}
		}

		hashByte32 := sha512.Sum512([]byte(string(result)))
		hashByte = hashByte32[:]
	}

	return &Byte_hash{
		Byte_slice: hashByte,
		String:     hex.EncodeToString(hashByte),
	}
}
