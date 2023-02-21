/*
Created on Tue Feb 21 17:31:18 2023
@author: Dmitry White
*/
package kata

import (
	"regexp"
)

/*
	TODO: Validate if a user input string is alphanumeric.
	The string has the following conditions to be alphanumeric:
	- At least one character ("" is not valid)
	- Allowed characters are uppercase/lowercase latin letters and digits from 0 to 9
	- No whitespaces / underscore
*/

const alphanumStr = "^[a-zA-Z0-9]{1,}$"

func alphanumeric(str string) bool {
	match, err := regexp.MatchString(alphanumStr, str)
	if err != nil {
		panic(err)
	}
	return match
}

// func alphanumeric(str string) bool {
// 	alphanumReg := regexp.MustCompile(alphanumStr)
// 	return alphanumReg.MatchString(str)
// }
