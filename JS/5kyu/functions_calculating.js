/**
 * Created on Wed Sep 27 23:44:14 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Write calculations using functions and get the results.
  * Let's have a look at some examples:
  *     seven(times(five())); // must return 35
  *     four(plus(nine())); // must return 13
  *     eight(minus(three())); // must return 5
  *     six(dividedBy(two())); // must return 3
  */

function operation(num, exp)
{
  return exp === undefined ? num : exp(num);
}

function zero(exp) {
  return operation(0, exp);
}

function one(exp) {
  return operation(1, exp);
}

function two(exp) {
  return operation(2, exp);
}

function three(exp) {
  return operation(3, exp);
}

function four(exp) {
  return operation(4, exp);
}

function five(exp) {
  return operation(5, exp);
}

function six(exp) {
  return operation(6, exp);
}

function seven(exp) {
  return operation(7, exp);
}

function eight(exp) {
  return operation(8, exp);
}

function nine(exp) {
  return operation(9, exp);
}

function plus(num) {
  return function(arg) {
    return arg + num;
  };
}
function minus(num) {
  return function(arg) {
    return arg - num;
  };
}
function times(num) {
  return function(arg) {
    return arg * num;
  };}
function dividedBy(num) {
  return function(arg) {
    return arg / num;
  };
}
