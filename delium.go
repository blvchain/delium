package delium

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"log"
	"strconv"
	"strings"
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

/*
D256C is complex delium that performs a specific method of D256:

1. Path example: "2h4usk#5/73uytg#9/#4"
  - path := "addonString1 # deleteStep1 / addonString2 # deleteStep2 / # justDeleteStep3"

2. Purpose:
  - The function takes an initial string `strData` and a `path` (which is not used in this function directly) and processes the hash of `strData` through several iterations.
  - Each iteration involves appending additional strings and recalculating the hash.

3. Procedure:
  - Compute the SHA-256 hash of the input string `strData` to generate an initial hash.
  - Convert this hash from a byte slice to a hexadecimal string representation.
  - Split the hexadecimal hash string using "/" as the delimiter to handle different segments.
  - For each segment:
  - Split the segment by "#" to extract an additional string and a delete step value.
  - Convert the delete step value from a string to an integer.
  - Create a new string by appending the additional string to the current hash string.
  - Recompute the hash of this new string and update the current hash string.
  - Return a `D_hash` struct containing the final hash as both a byte slice and a string.

4. Detailed Steps:
  - Compute the SHA-256 hash using `sha256.Sum256([]byte(strData))`, which provides a 64-byte hash.
  - Convert this hash to a hexadecimal string with `hex.EncodeToString(dataHash[:])`.
  - Split the string into parts using `strings.Split(strDataHash, "/")` to separate different segments for further processing.
  - Iterate over these segments. For each segment:
  - Split by "#" using `strings.Split(part, "#")` to get the `addonString` and the `deleteStep`.
  - Convert `deleteStep` from a string to an integer using `strconv.Atoi(d[1])`. If conversion fails, log an error.
  - Create a new string `newString` by appending `addonString` to the current `strDataHash`.
  - Call `D256(newString, deleteStep, 1).String` to recompute the hash and update `strDataHash`.
  - Finally, create and return a `D_hash` struct containing the final `strDataHash` as a byte slice and string.

The `D_hash` struct is used to encapsulate the result, making it easier to manage and use the final hash value.

The `path` parameter is included for potential future use or compatibility with other parts of the codebase but is not utilized in the current implementation of the function.
*/
func D256C(strData string, path string) *D_hash {

	dataHash := sha256.Sum256([]byte(strData))
	var strDataHash string = hex.EncodeToString(dataHash[:])
	parts := strings.Split(path, "/")

	for _, part := range parts {
		d := strings.Split(part, "#")

		addonString := d[0]
		newString := strDataHash + addonString

		deleteStep, strconvErr := strconv.Atoi(d[1])
		if strconvErr != nil {
			log.Fatal("Hashing path is incorrect!")
		}

		strDataHash = D256(newString, deleteStep, 1).String
	}

	return &D_hash{
		Byte_slice: []byte(strDataHash),
		String:     strDataHash,
	}
}

/*
D512C is complex delium that performs a specific method of D512:

1. Path example: "2h4usk#5/73uytg#9/#4"
  - path := "addonString1 # deleteStep1 / addonString2 # deleteStep2 / # justDeleteStep3"

2. Purpose:
  - The function takes an initial string `strData` and a `path` (which is not used in this function directly) and processes the hash of `strData` through several iterations.
  - Each iteration involves appending additional strings and recalculating the hash.

3. Procedure:
  - Compute the SHA-512 hash of the input string `strData` to generate an initial hash.
  - Convert this hash from a byte slice to a hexadecimal string representation.
  - Split the hexadecimal hash string using "/" as the delimiter to handle different segments.
  - For each segment:
  - Split the segment by "#" to extract an additional string and a delete step value.
  - Convert the delete step value from a string to an integer.
  - Create a new string by appending the additional string to the current hash string.
  - Recompute the hash of this new string and update the current hash string.
  - Return a `D_hash` struct containing the final hash as both a byte slice and a string.

4. Detailed Steps:
  - Compute the SHA-512 hash using `sha512.Sum512([]byte(strData))`, which provides a 64-byte hash.
  - Convert this hash to a hexadecimal string with `hex.EncodeToString(dataHash[:])`.
  - Split the string into parts using `strings.Split(strDataHash, "/")` to separate different segments for further processing.
  - Iterate over these segments. For each segment:
  - Split by "#" using `strings.Split(part, "#")` to get the `addonString` and the `deleteStep`.
  - Convert `deleteStep` from a string to an integer using `strconv.Atoi(d[1])`. If conversion fails, log an error.
  - Create a new string `newString` by appending `addonString` to the current `strDataHash`.
  - Call `D512(newString, deleteStep, 1).String` to recompute the hash and update `strDataHash`.
  - Finally, create and return a `D_hash` struct containing the final `strDataHash` as a byte slice and string.

The `D_hash` struct is used to encapsulate the result, making it easier to manage and use the final hash value.

The `path` parameter is included for potential future use or compatibility with other parts of the codebase but is not utilized in the current implementation of the function.
*/
func D512C(strData string, path string) *D_hash {

	dataHash := sha512.Sum512([]byte(strData))
	var strDataHash string = hex.EncodeToString(dataHash[:])
	parts := strings.Split(path, "/")

	for _, part := range parts {
		d := strings.Split(part, "#")

		addonString := d[0]
		newString := strDataHash + addonString

		deleteStep, strconvErr := strconv.Atoi(d[1])
		if strconvErr != nil {
			log.Fatal("Hashing path is incorrect!")
		}

		strDataHash = D512(newString, deleteStep, 1).String
	}

	return &D_hash{
		Byte_slice: []byte(strDataHash),
		String:     strDataHash,
	}
}
