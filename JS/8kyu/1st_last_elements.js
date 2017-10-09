/**
 * Created on Mon Oct 09 20:40:11 2017
 * @author: Dmitry White
 */

 /**
  * TODO: You are given a list of characters as a comma separated string.
  * Write a function to return a string containing all except the first
  * and last characters, separated by spaces. If the input string is empty,
  * or the removal of the first and last items would cause the string
  * to be empty, return null value.
  */

function array(arr){
  arr = arr.split(",");
  if (arr.length < 3) return null;
  arr.shift();
  arr.pop();
  return arr.join(" ");
}
