/**
 * Created on Sun Oct 29 21:18:52 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Create a function that will add numbers together
  * when called in succession. We also want to be able to continue
  * to add numbers to our chain. A single call should return
  * the number passed in.
  */

function add(n) {
    const chain = function(item) {
        return add(n + item);
    };
    chain.valueOf = chain.toString = function() {
        return n;
    };
    return chain;
}
