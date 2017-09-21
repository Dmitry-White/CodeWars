/**
 * Created on Wed Sep 20 00:26:21 2017
 * @author: Dmitry White
 */

 /**
  * TODO: An isogram is a word that has no repeating letters, consecutive or non-consecutive.
  * Implement a function that determines whether a string that contains only letters is an isogram.
  * Assume the empty string is an isogram. Ignore letter case.
  */

function isIsogram(str){
  var chars = {}, len = str.length;
  str = str.toLowerCase();
  for (var i = 0; i < len; ++i) {
    if (!(str[i] in chars)) {
      chars[str[i]] = 1;
    } else { return false; }
  }
  return true
}
