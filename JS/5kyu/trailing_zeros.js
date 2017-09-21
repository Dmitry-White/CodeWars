/**
 * Created on Thu Sep 22 00:30:01 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Write a program that will calculate the number
  * of trailing zeros in a factorial of a given number.
  */

function zeros(n) {

  let totalZeros = 0;

  while (n > 0) {

    n = Math.floor(n / 5);
    totalZeros += n;
  }

  return totalZeros;
}
