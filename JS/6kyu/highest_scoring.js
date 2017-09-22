/**
 * Created on Thu Sep 22 23:12:32 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Given a string of words (x), you need to find the highest scoring word.
  * Each letter of a word scores points according to it's position in the alphabet. a=1, z=26 and everything inbetween.
  * You need to return the highest scoring word as a string.
  * If two words score the same, return the word that appears earliest in the original string.
  * All letters will be lower case and all inputs will be valid.
  */

var alpha = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'];
var words = x.split(" ");
var sum, max = 0;
var final = "";
words.forEach(function(word) {
  sum = 0;
  for (var i in word){
    sum += alpha.indexOf(word[i])+1;
  }
  if (sum > max) {
    max = sum;
    final = word;
  }

});
return final;
