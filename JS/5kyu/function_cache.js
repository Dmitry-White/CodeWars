/**
 * Created on Sun Oct 29 21:35:36 2017
 * @author: Dmitry White
 */

 /**
  * TODO: create a function wrapper, which takes a function and
  * caches its results depending on the arguments,
  * that were applied to the function.
  */

function cache(func) {
    let result = {};
    return function() {
        const args = JSON.stringify(arguments);
        if (result.hasOwnProperty(args) === false) {
            result[args] = func.apply(null, arguments);
        };
        return result[args];
    };
}
