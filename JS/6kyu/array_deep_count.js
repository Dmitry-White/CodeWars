/**
 * Created on Fri Feb 24 22:16:17 2023
 * @author: Dmitry White
 */

/**
 * TODO: You are given an array.
 * Complete the function that returns the number of ALL elements within an array,
 * including any nested arrays.
 */

const deepCount = (arr) => {
  if (arr.length === 0) {
    return 0;
  }

  const curr = arr.reduce((acc, curr) => {
    if (!Array.isArray(curr)) {
      return acc + 1;
    }
    return acc + 1 + deepCount(curr);
  }, 0);

  return curr;
};
