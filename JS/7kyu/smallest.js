/**
 * Created on Wed Sep 20 00:41:56 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Given an array of integers your solution should find the smallest integer.
  */

class SmallestIntegerFinder {
  findSmallestInt(args) {
    return args.sort(function(a,b) { return a-b; })[0];
  }
}
