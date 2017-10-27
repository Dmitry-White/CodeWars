/**
 * Created on Fri Oct 27 09:33:12 2017
 * @author: Dmitry White
 */

 /**
  * TODO: In this kata, you must create a digital root function.
  * A digital root is the recursive sum of all the digits in a number.
  * Given n, take the sum of the digits of n. If that value has two digits,
  * continue reducing in this way until a single-digit number is produced.
  * This is only applicable to the natural numbers.
  */

function digital_root(n) {
    let temp;
    while (n > 9) {
        temp = n;
        n = 0;
        while (temp > 0) {
            n += temp % 10;
            temp = Math.floor(temp / 10);
      }
    }
    return n;
}
