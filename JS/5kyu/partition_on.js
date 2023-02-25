/**
 * Created on Sat Feb 25 19:04:12 2017
 * @author: Dmitry White
 */

/**
 * TODO: Write a function which partitions a list of items based on a given predicate.
 * After the partition function is run, the list should be of the form [ F, F, F, T, T, T ]
 * where the Fs (resp. Ts) are items for which the predicate function returned false (resp. true).
 * Partition the `items` array so that all values for which `pred` returns true are at the end,
 * returning the index of the first true value.
 */
function partitionOn(pred, items) {
  const falsePart = [];
  const truePart = [];

  items.forEach((item) => {
    if (pred(item)) {
      truePart.push(item);
    } else {
      falsePart.push(item);
    }
  });

  items.length = 0;
  falsePart.forEach((item) => items.push(item));
  truePart.forEach((item) => items.push(item));

  return falsePart.length;
}

const items = [1, 2, 3, 4, 5, 6];
function isEven(n) {
  return n % 2 == 0;
}
const i = partitionOn(isEven, items);

console.log("i: ", i);
console.log("False Partition: ", items.slice(0, i)); // [1, 3, 5];
console.log("True Partition: ", items.slice(i)); // [2, 4, 6];
