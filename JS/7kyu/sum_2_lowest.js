/**
 * Created on Wed Sep 20 00:36:18 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Create a function that returns the sum of
  * the two lowest positive numbers given an array of minimum 4 integers.
  * No floats or empty arrays will be passed.
  */

function sumTwoSmallestNumbers(numbers) {
    numbers = numbers.sort(function(a,b) { return a-b; });
    return numbers[0] + numbers[1];
}
