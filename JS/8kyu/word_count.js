/**
 * Created on Wed Sep 27 22:44:41 2017
 * @author: Dmitry White
 */

 /**
  * TODO: You need comie up with a function that returns word count
  * from a given string
  */

function countWords(str) {
  var words = str.match(/[a-zA-Z0-9\-'`]+/g);
  return (str == '' || !words) ? 0 : words.length;
}
