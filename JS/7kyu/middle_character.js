/**
 * Created on Wed Sep 27 22:46:15 2017
 * @author: Dmitry White
 */

 /**
  * TODO: You are going to be given a word.
  * Your job is to return the middle character of the word.
  * If the word's length is odd, return the middle character.
  * If the word's length is even, return the middle 2 characters.
  */

function getMiddle(s) {
    var s_len = s.length;
    var central = Math.floor(s_len/2);
    return s_len % 2 === 0 ? s.slice(central - 1,central + 1)  : s[central];
}
