/**
 * Created on Thu Sep 21 21:45:15 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Write a function that will return the count of distinct
  * case-insensitive alphabetic characters and numeric digits 
  * that occur more than once in the input string.
  * The input string can be assumed to contain only
  * alphabets (both uppercase and lowercase) and numeric digits.
  */

var _ = require('lodash');
function duplicateCount(text){
  text = text.toLowerCase();
  var count = 0;
  var temp = {};
  for (var i in text) {
    if ( temp.hasOwnProperty(text[i]) ) {
      temp[text[i]] += 1
    }else {
      temp[text[i]] = 1;
    }
  }
  return _.values(temp).reduce(function(a,b) { return b>1 ?  a+=1 : a+=0; },0);
}
