/**
 * Created on Fri Feb 24 23:04:54 2023
 * @author: Dmitry White
 */

/**
 * TODO: Create the function prefill that returns an array of n elements that all have the same value v.
 * If n is anything other than an integer or integer-formatted string that is >=0, throw a TypeError.
 */

function prefill(n, v) {
  try {
    if (typeof n === "boolean") {
      throw "";
    }
    const len = Number(n);
    const arr = new Array(len);

    return arr.fill(v);
  } catch (error) {
    throw new TypeError(`${n} is invalid`);
  }
}
