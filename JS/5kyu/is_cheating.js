/**
 * Created on Sun Sep 24 13:16:11 2017
 * @author: Dmitry White
 */

 /**
  * TODO: A friend of mine takes a sequence of numbers
  * from 1 to n (where n > 0).
  * Within that sequence, he chooses two numbers, a and b.
  * He says that the product of a and b should be equal to
  * the sum of all numbers in the sequence, excluding a and b.
  * Given a number n, could you tell me the numbers
  * he excluded from the sequence?
  * [(a, b), ...] or [[a, b], ...] or {{a, b}, ...} or or [{a, b}, ...]
  * It happens that there are several possible (a, b).
  * The function returns an empty array if no possible numbers
  * are found which will prove that my friend has not told the truth!
  */

function removeNb (n) {
  var result = [];
  var j = 0;
  var sum = (n * (n+1)) / 2;
  for (var i = 1; i <= n; i++) {
    if (((sum - i) % (i + 1)) === 0) {
      j = (sum - i) / (i + 1);
      if ((j < n) && (j !== i)) {
        result.push([i,j]);
      }
    }
  }
  return result;
}
