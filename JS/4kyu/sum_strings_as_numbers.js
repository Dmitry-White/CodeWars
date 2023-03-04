/**
 * Created on Sat Mar 04 13:54:12 2023
 * @author: Dmitry White
 */

/**
 * TODO: Given the string representations of two integers,
 * return the string representation of the sum of those integers.
 * A string representation of an integer will contain no characters
 * besides the ten numerals "0" to "9".
 */

const getSum = (first, second) => {
  const diff = second.length - first.length;
  let sum = "";
  let carry = 0;

  for (let i = first.length - 1; i >= 0; i--) {
    const curr =
      Number(first.charAt(i)) + Number(second.charAt(i + diff)) + carry;

    if (curr >= 10) {
      sum = (curr % 10) + sum;
      carry = Math.floor(curr / 10);
    } else {
      sum = curr + sum;
      carry = 0;
    }
  }

  if (carry) {
    sum = carry + sum;
  }

  return sum;
};

function sumStrings(a, b) {
  const result = a.length > b.length ? getSum(a, b) : getSum(b, a);
  return result.replace(/^0+/, "");
}
