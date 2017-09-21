/**
 * Created on Thu Sep 21 20:05:11 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Write a reverseWords function that accepts a string a parameter,
  * and reverses each word in the string.
  * Every space should stay, so you cannot use words from Prelude.
  */

function reverseWords(str) {
  var words = str.split(" ");
  for (var i in words) {
    words[i] = words[i].split("").reverse().join("");
  }
  return words.join(" ");
}
