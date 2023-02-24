/*
Created on Fri Feb 24 01:05:27 2023
@author: Dmitry White
*/
package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

/*
TODO: Create the function that converts a given string into an md5 hash.
The return value should be encoded in hexadecimal.
*/
func PassHash(str string) string {
	hash := md5.New()
	io.WriteString(hash, str)

	src := hash.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	return string(dst)
}
