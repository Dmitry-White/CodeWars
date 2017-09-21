/**
 * Created on Thu Sep 21 20:31:54 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Build Tower by the following given argument:
  * number of floors (integer and always greater than 0).
  * Tower block is represented as *
  */

function towerBuilder(nFloors) {
  var tower = [];
  var side = "";
  var blocks = 1;
  var floor_size = 2 * nFloors - 1;
  for( var i = 0; i < nFloors; i++) {
    side = " ".repeat((floor_size - blocks)/2);
    tower.push( side + ("*".repeat(blocks)) + side);
    blocks += 2;
  }
  return tower;
}
