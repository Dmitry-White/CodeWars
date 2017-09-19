/**
 * Created on Wed Sep 20 00:45:03 2017
 * @author: Dmitry White
 */

 /**
  * TODO: You get an array of numbers, return the sum of all of the positives ones.
  */

function positiveSum(arr) {
  return arr.reduce(function(a,b) { return (b > 0) ? a+b : a },0);
}
