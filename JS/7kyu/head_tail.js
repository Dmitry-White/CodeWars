/**
 * Created on Fri Sep 29 20:11:37 2017
 * @author: Dmitry White
 */

 /**
  * TODO: GHCi, version 7.6.3:
  * λ head [1,2,3,4,5]
  * 1
  * λ tail [1,2,3,4,5]
  * [2,3,4,5]
  * λ init [1,2,3,4,5]
  * [1,2,3,4]
  * λ last [1,2,3,4,5]
  * 5
  * Implement these functions. Make sure it doesn't edit the array.
  */

function head(array) {
    return array[0];
}

function tail(array) {
    return array.slice(1,array.length);
}

function init(array) {
    return array.slice(0,array.length-1);
}

function last(array) {
    return array[array.length-1];
}
