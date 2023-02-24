/*
Created on Fri Feb 24 01:05:27 2023
@author: Dmitry White
*/
package kata

import (
	"crypto/md5"
	"fmt"
)

/*
TODO: Create the function that converts a given string into an md5 hash.
The return value should be encoded in hexadecimal.
*/
func PassHash(str string) string {
	hash := md5.Sum([]byte(str))

	return fmt.Sprintf("%x", hash)
}
