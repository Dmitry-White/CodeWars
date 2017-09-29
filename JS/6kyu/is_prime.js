/**
 * Created on Fri Sep 29 20:46:12 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Define a function isPrime/is_prime() that takes one integer argument
  * and returns true/True or false/False depending on if the integer is a prime.
  */

function isPrime(num) {
  if (num <= 0) {
    return false;
  }
  for(var i = 2; i < num; i++)
    if(num % i === 0) return false;
  return num !== 1;
}
