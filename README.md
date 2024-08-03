# Delium

🔑 Super strong encryption method based on deleting.

This method, inspired by the film Oppenheimer, is derived from the concept of an atomic explosion. It works by visualizing a string, encoded using the SHA-256 or SHA-512 method, as a specific number of atoms placed together. In an atomic explosion, almost all atoms are destroyed, but our goal is to mathematically make the reverse engineering probability almost zero. Therefore, we remove several specific character sequences from the strings and then re-encode the resulting string using the SHA-256 or SHA-512 method.

**This method is very simple but provides a high level of security in terms of safety.**

## Installation

```bash
go get github.com/yahya-aghdam/Delium
```

## Import

```go
import "github.com/yahya-aghdam/Delium"
```

## Functions

This lib has two functions Delium256 and Delium512, The input and output is same but method of SHA encryption.

### `Delium256`

Processes a string data by first hashing it with SHA-256, then repeatedly deleting characters from the resulting hash string at specified intervals, and finally hashing the modified string again with SHA-256.

**Parameters:**

- `strData`: A string representing the data to be processed and hashed.
- `deleteStep`: An integer specifying the interval at which characters will be deleted from the hash string.
- `repeat`: An integer specifying how many times the deletion process should be applied.

**Returns:**

- A pointer to a `Byte_hash` struct containing:
  - `Byte_slice`: A byte slice of the final SHA-256 hash after applying the deletion process the specified number of times.
  - `String`: A hexadecimal string representation of the final SHA-256 hash.

### `Delium512`

Processes a string data by first hashing it with SHA-512, then repeatedly deleting characters from the resulting hash string at specified intervals, and finally hashing the modified string again with SHA-512.

**Parameters:**

- `strData`: A string representing the data to be processed and hashed.
- `deleteStep`: An integer specifying the interval at which characters will be deleted from the hash string.
- `repeat`: An integer specifying how many times the deletion process should be applied.

**Returns:**

- A pointer to a `Byte_hash` struct containing:
  - `Byte_slice`: A byte slice of the final SHA-512 hash after applying the deletion process the specified number of times.
  - `String`: A hexadecimal string representation of the final SHA-512 hash.

## Usage example
```go
package main

import (
    "fmt"
    "github.com/yahya-aghdam/Delium"
)

func main() {

    result := Delium256("example mnemonic", 3, 5)

    fmt.Println("Byte Slice:", result.Byte_slice)  // prints the byte slice of the final hash

    fmt.Println("Hex String:", result.String)      // prints the hexadecimal string of the final hash
    
}
```