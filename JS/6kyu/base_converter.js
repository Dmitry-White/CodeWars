/**
 * Created on Fri Feb 24 19:40:25 2017
 * @author: Dmitry White
 */

/**
 * TODO: Implement a base converter, which converts positive integers between arbitrary bases / alphabets.
 * The function convert() should take an input (string),
 * the source alphabet (string) and the target alphabet (string).
 * You can assume that the input value always consists of characters from the source alphabet.
 * You don't need to validate it.
 * The maximum input value can always be encoded in a number without loss of precision in JS.
 * The function must work for any arbitrary alphabets, not only the pre-defined ones.
 * You don't have to consider negative numbers.
 */

const Alphabet = {
  BINARY: "01",
  OCTAL: "01234567",
  DECIMAL: "0123456789",
  HEXA_DECIMAL: "0123456789abcdef",
  ALPHA_LOWER: "abcdefghijklmnopqrstuvwxyz",
  ALPHA_UPPER: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
  ALPHA: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
  ALPHA_NUMERIC:
    "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
};

/**
 * Source To Decimal method:
 *
 * Right-to-left take 1 `Item` from `Input`.
 * Find its index in `Source` and store it with `Place`.
 * Then
 *      `Result += Item * (Radix ^ Place)`
 * through all stored `Items`.
 */
const sourceToDecimal = (input, source) => {
  const radix = source.length;
  const remainder = {};
  let quotient = input;
  let place = 0;

  while (quotient.length !== place) {
    remainder[place] = source.split("").indexOf(quotient[place]);

    place++;
  }

  let result = 0;
  for (let index = 1; index <= place; index++) {
    result += remainder[place - index] * Math.pow(radix, index - 1);
  }

  return result;
};

/**
 * Decimal To Target method:
 *
 * Get remainder of `Input` divided by `Target Radix`
 * to find a corresponding `Target Item` to store with `Place`.
 * Reduce `Input` to `Quotient` and repeat, until `Input` reaches 0.
 * Then combine all stored `Items` in reverse.
 */
const decimalToTarget = (input, target) => {
  const radix = target.length;
  const remainder = {};
  let quotient = input;
  let place = 0;

  do {
    remainder[place] = target[quotient % radix];
    quotient = Math.floor(quotient / radix);

    place++;
  } while (quotient !== 0);

  let result = "";
  for (let index = 1; index <= place; index++) {
    result += `${remainder[place - index]}`;
  }

  return result;
};

const convert = (input, source, target) => {
  const decimalInput = sourceToDecimal(input, source);
  return decimalToTarget(decimalInput, target);
};
