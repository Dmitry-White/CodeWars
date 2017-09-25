/**
 * Created on Mon Sep 25 22:53:22 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Create a compose function to carry out this task,
  * which will be passed two functions or lambdas. Remember
  * that the resulting composed function may be passed multiple arguments!
  */


function compose(f,g) {
  return function() {
    return f.call(this, g.apply(this, arguments));
  }
}
