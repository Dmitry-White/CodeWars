/**
 * Created on Sun Oct 29 21:06:26 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Write a compose function which can compose any number
  * of functions together.
  */

function compose(...item) {
    return function(items) {
        return item.reduceRight((accumulator, element) => element(accumulator), items);
    }
}
