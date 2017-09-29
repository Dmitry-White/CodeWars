/**
 * Created on Fri Sep 29 23:38:54 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Create a function stringify which accepts an argument list/$list and
  * returns a string representation of the list.
  * The string representation of the list starts with the value of
  * the current Node, specified by its data/$data/Data property,
  * followed by a whitespace character, an arow and another
  * whitespace character (" -> "), followed by the rest of the list.
  * The end of the string representation of a list must always end
  * with null/NULL/None.
  */

function stringify(list) {
    var result = [];
    if (list === null) {
      return "null";
    }else if (list.next == null) {
      result.push(list.data);
      return result + " -> " + "null";
    }else {
      result.push(list.data);
      return result + " -> " + stringify(list.next);
    }
}
