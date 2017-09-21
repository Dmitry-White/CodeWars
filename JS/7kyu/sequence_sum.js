/**
 * Created on Thu Sep 21 18:35:07 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Your task is to make function, which returns the sum of a sequence of integers.
  * The sequence is defined by 3 non-negative values: begin, end, step.
  * If begin value is greater than the end, function should returns 0
  */

const sequenceSum = (begin, end, step) => {
  var sum = 0;
  for(var i = begin;i <= end;i += step) {
    sum += i;
  }
  return sum;
};
