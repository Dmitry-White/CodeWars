/**
 * Created on Thu Oct 26 21:30:28 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Given an input n, write a function always that
  * returns a function which returns n.
  */

function always (n) {
    return () => {
        return n;
    }
}
