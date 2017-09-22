/**
 * Created on Fri Sep 22 23:40:09 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Given an array of integers, remove the smallest value.
  * Do not mutate the original array/list.
  * If there are multiple elements with the same value,
  * remove the one with a lower index.
  * If you get an empty array/list, return an empty array/list.
  * Don't change the order of the elements that are left.
  */

function removeSmallest(numbers) {
  var min = 0;
  if (numbers === []) { return [] }
  min = numbers.indexOf(Math.min.apply(Math, numbers));
  numbers.splice( min, 1 );
  return numbers;
}
