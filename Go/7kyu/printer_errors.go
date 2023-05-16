/*
Created on Mon May 15 22:31:27 2023
@author: Dmitry White
*/
package kata

import (
	"fmt"
	"regexp"
	"strings"
)

/*
	TODO: Write a function which given a control string
	will return the error rate of the printer as a string
	representing a rational whose numerator is the number of errors and
	the denominator the length of the control string.
	Don't reduce this fraction to a simpler expression.
	Good control string contains only characters from 'a' to 'm'.
*/

var controlStringRegex = regexp.MustCompile(`[a-m]+`)

func PrinterError(s string) string {
	denominator := len(s)
	nominator := 0

	strArr := strings.Split(s, "")
	for _, char := range strArr {
		if !controlStringRegex.MatchString(char) {
			nominator++
		}
	}

	return fmt.Sprintf("%d/%d", nominator, denominator)
}
