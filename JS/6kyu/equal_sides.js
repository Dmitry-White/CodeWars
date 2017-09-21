/**
 * Created on Wed Sep 20 00:29:44 2017
 * @author: Dmitry White
 */

 /**
  * TODO: You are going to be given an array of integers.
  * Your job is to take that array and find an index N where the sum of the integers
  * to the left of N is equal to the sum of the integers to the right of N.
  * If there is no index that would make this happen, return -1.
  */

function findEvenIndex(arr)
{
  var len = arr.length, left_sum = 0, right_sum = 0;
  var left_side = [], right_side = [];
  var sum = function(a,b){ return a+b; }

  for (var i = 0; i < len; i++) {
    left_side = arr.slice(0,i);
    right_side = arr.slice(i+1,len);
    left_sum = left_side.reduce(sum,0);
    right_sum = right_side.reduce(sum,0);
    if (left_sum === right_sum) {
      return i;
    }
  }
  return -1;
}
