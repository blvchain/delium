<p align="center">
  <img src="./delium.png" alt="delium logo" width="200" height="200" style="display: block; margin: 30 auto" />
</p>

![License](https://img.shields.io/badge/License-MIT-blue)
![Version](https://img.shields.io/badge/Version-2.0.0-orange)

# Delium

**🔑 This method is very simple but provides a high level of security in terms of safety.**

This method, inspired by the film Oppenheimer, is derived from the concept of an atomic explosion. It works by visualizing a string, encoded using the SHA-256 or SHA-512 method, as a specific number of atoms placed together. In an atomic explosion, almost all atoms are destroyed, but our goal is to mathematically make the reverse engineering probability almost zero. Therefore, we remove several specific character sequences from the strings and then re-encode the resulting string using the SHA-256 or SHA-512 method.

- [Delium](#delium)
  - [Methods](#methods)
  - [Functions](#functions)
    - [Simple delium](#simple-delium)
      - [`D256`](#d256)
      - [`D512`](#d512)
    - [Complex delium](#complex-delium)
      - [`D256C`](#d256c)
      - [`D512C`](#d512c)
  - [Installation](#installation)
  - [Import](#import)
  - [Usage example](#usage-example)

## Methods

Delium has two method for encryption, 1-Simple delium 2-Complex delium.
In simple delium input string will hash with sha-256/sha-512 and in every cycle of hashing will delete specific chars from hashid string. But in complex delium, we have a `path` for adding and deleting for every cycle.

## Functions

### Simple delium

#### `D256`

Processes a string data by first hashing it with SHA-256, then repeatedly deleting characters from the resulting hash string at specified intervals, and finally hashing the modified string again with SHA-256.

**Parameters:**

- `strData`: A string representing the data to be processed and hashed.
- `deleteStep`: An integer specifying the interval at which characters will be deleted from the hash string.
- `repeat`: An integer specifying how many times the deletion process should be applied.

**Returns:**

- A pointer to a `D_hash` struct containing:
  - `Byte_slice`: A byte slice of the final SHA-256 hash after applying the deletion process the specified number of times.
  - `String`: A hexadecimal string representation of the final SHA-256 hash.

#### `D512`

Processes a string data by first hashing it with SHA-512, then repeatedly deleting characters from the resulting hash string at specified intervals, and finally hashing the modified string again with SHA-512.

**Parameters:**

- `strData`: A string representing the data to be processed and hashed.
- `deleteStep`: An integer specifying the interval at which characters will be deleted from the hash string.
- `repeat`: An integer specifying how many times the deletion process should be applied.

**Returns:**

- A pointer to a `D_hash` struct containing:
  - `Byte_slice`: A byte slice of the final SHA-256 hash after applying the deletion process the specified number of times.
  - `String`: A hexadecimal string representation of the final SHA-256 hash.

### Complex delium

Processes a string data by first hashing it with SHA-256, then add strings based on path to the end of hashed string and delete chars based on path. Path can has NOT any addon string but must has at least the deleting chars number.

Path example:
`"2h4usk#5/73uytg#9/#4"`

In this path we use delium 3 times. First we use delium with adding `2h4usk` to end of hashing and deleting chars based on `5` intervals. Then we add `73uytg` to end of new hash and delete every chars based on `9` intervals. At least we DONT add any string but we do simple delium just once with `4` delete step.
⚠️⚠️⚠️ PATH STRING NOT HAS ANY SPACE ⚠️⚠️⚠️

#### `D256C`

**Parameters:**

- `strData`: A string representing the data to be processed and hashed.
- `path`: A string representing the path of complex delium.

**Returns:**

- A pointer to a `D_hash` struct containing:
  - `Byte_slice`: A byte slice of the final SHA-256 hash after applying the complex delium.
  - `String`: A hexadecimal string representation of the final SHA-256 hash.

#### `D512C`

**Parameters:**

- `strData`: A string representing the data to be processed and hashed.
- `path`: A string representing the path of complex delium.

**Returns:**

- A pointer to a `D_hash` struct containing:
  - `Byte_slice`: A byte slice of the final SHA-512 hash after applying the complex delium.
  - `String`: A hexadecimal string representation of the final SHA-512 hash.

## Installation

```bash
go get github.com/yahya-aghdam/Delium
```

## Import

```go
import "github.com/yahya-aghdam/Delium"
```

## Usage example

```go
package main

import (
    "fmt"
    "github.com/yahya-aghdam/Delium"
)

func main() {

  // Simple delium
  simpleDelium := delium.D256("example mnemonic", 3, 5)

  fmt.Println("Byte Slice:", simpleDelium.Byte_slice)  // prints the byte slice of the simple delium hash
  fmt.Println("Hex String:", simpleDelium.String)      // prints the hexadecimal string of the simple delium hash

  // Complex delium
  path:= "2h4usk#5/73uytg#9/#4"
  complexDelium := delium.D256C("example mnemonic", path)

  fmt.Println("Byte Slice:", complexDelium.Byte_slice)  // prints the byte slice of the complex delium hash
  fmt.Println("Hex String:", complexDelium.String)      // prints the hexadecimal string of the complex delium hash

}
```
