/**
 * Created on Sun Oct 29 21:30:42 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Write a nouveau function that replicates all
  * the behavior of the new operator.
  */

function nouveau (Constructor,...args) {
    return Reflect.construct(Constructor, args)
}
