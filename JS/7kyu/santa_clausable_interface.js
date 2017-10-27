/**
 * Created on Fri Oct 27 08:51:31 2017
 * @author: Dmitry White
 */

 /**
  * TODO: You need to implement a method which checks for "SantaClausable"
  * interface. The SantaClausable interface is implemented,
  * if all of the following methods are defined on an object:
  * sayHoHoHo()
  * distributeGifts()
  * goDownTheChimney()
  */

function isSantaClausable(obj) {
    return (((typeof(obj.sayHoHoHo) === 'function')) &&
             (typeof(obj.distributeGifts) === 'function') &&
             (typeof(obj.goDownTheChimney) === 'function'))? true : false;
}
