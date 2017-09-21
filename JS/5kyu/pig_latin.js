/**
 * Created on Thu Sep 21 23:14:11 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Move the first letter of each word to the end of it,
  * then add 'ay' to the end of the word.
  */

function pigIt(str){
  var temp = 0;
  var array = str.split(" "), arr = [], result = [];
  for (var i in array) {
    arr = array[i].split("");
    temp = arr.shift();
    result.push(arr.join("") + temp + "ay");
  }
  return result.join(" ");
}
