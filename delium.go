package delium

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

type D_hash struct {
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
//   - A pointer to a D_hash struct containing:
//   - Byte_slice: A byte slice of the final SHA-256 hash after applying the deletion process the specified
//     number of times.
//   - String: A hexadecimal string representation of the final SHA-256 hash.
//
// The function works as follows:
//  1. Computes the SHA-256 hash of the input strData string.
//  2. For a specified number of repetitions, deletes every `deleteStep`-th character from the hash string
//     generated in the previous step.
//  3. Computes the SHA-256 hash of the modified string after each repetition.
//  4. Returns a D_hash struct with the final SHA-256 hash in both byte slice and hexadecimal string formats.
//
// Example:
//
//	result := delium.D256("example strData", 3, 5)
//	fmt.Println(result.Byte_slice)  // prints the byte slice of the final hash
//	fmt.Println(result.String)      // prints the hexadecimal string of the final hash
func D256(strData string, deleteStep int, repeat int) *D_hash {

	dataHash := sha256.Sum256([]byte(strData))
	var strDataHash string = hex.EncodeToString(dataHash[:])

	for i := 0; i < repeat; i++ {

		var result string = ""
		for r := 0; r < len(strDataHash); r++ {
			if (r+1)%deleteStep != 0 {
				result += string(strDataHash[r])
			}
		}

		hashByte32 := sha256.Sum256([]byte(result))
		strDataHash = hex.EncodeToString(hashByte32[:])
	}

	return &D_hash{
		Byte_slice: []byte(strDataHash),
		String:     strDataHash,
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
//   - A pointer to a D_hash struct containing:
//   - Byte_slice: A byte slice of the final SHA-512 hash after applying the deletion process the specified
//     number of times.
//   - String: A hexadecimal string representation of the final SHA-512 hash.
//
// The function works as follows:
//  1. Computes the SHA-512 hash of the input strData string.
//  2. For a specified number of repetitions, deletes every `deleteStep`-th character from the hash string
//     generated in the previous step.
//  3. Computes the SHA-512 hash of the modified string after each repetition.
//  4. Returns a D_hash struct with the final SHA-512 hash in both byte slice and hexadecimal string formats.
//
// Example:
//
//	result := delium.D512("example strData", 3, 5)
//	fmt.Println(result.Byte_slice)  // prints the byte slice of the final hash
//	fmt.Println(result.String)      // prints the hexadecimal string of the final hash
func D512(strData string, deleteStep int, repeat int) *D_hash {

	dataHash := sha512.Sum512([]byte(strData))
	var strDataHash string = hex.EncodeToString(dataHash[:])

	for i := 0; i < repeat; i++ {

		var result string = ""
		for r := 0; r < len(strDataHash); r++ {
			if (r+1)%deleteStep != 0 {
				result += string(strDataHash[r])
			}
		}

		hashByte32 := sha512.Sum512([]byte(result))
		strDataHash = hex.EncodeToString(hashByte32[:])
	}

	return &D_hash{
		Byte_slice: []byte(strDataHash),
		String:     strDataHash,
	}
}
