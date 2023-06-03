/**
 * Created on Sat Jun 03 20:27:36 2023
 * @author: Dmitry White
 */

/**
 * TODO: Create a function `correct`, which takes in integers M and N,
 * and a string of bits where the first M*N represent the content of the message,
 * the next M represent the parity bits for the rows,
 * and the final N represent the parity bits for the columns.
 * A single-bit error may or may not have appeared in the bit array.
 * The function should check to see if there is a single-bit error in the coded message,
 * correct it if it exists, and return the corrected string of bits.
 *
 * A system is transmitting messages in binary, however it is not a perfect transmission,
 * and sometimes errors will occur which result in a single bit flipping from 0 to 1, or from 1 to 0.
 * To resolve this, A 2-dimensional Parity Bit Code is used:
 * https://en.wikipedia.org/wiki/Multidimensional_parity-check_code
 */

const computeMatrix = (n, message) => {
  const bits1D = [...message];
  const bits2D = [];
  while (bits1D.length) bits2D.push(bits1D.splice(0, n));

  return bits2D;
};

const computeParities = (m, n, bits2D) => {
  const rowParitiesComputed = [];
  const columnParitiesComputed = [];

  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      if (!rowParitiesComputed[i]) {
        rowParitiesComputed[i] = bits2D[i][j];
      } else {
        rowParitiesComputed[i] += bits2D[i][j];
      }

      if (!columnParitiesComputed[j]) {
        columnParitiesComputed[j] = bits2D[i][j];
      } else {
        columnParitiesComputed[j] += bits2D[i][j];
      }

      if (i === m - 1) {
        columnParitiesComputed[j] = columnParitiesComputed[j] % 2 === 0 ? 0 : 1;
      }
    }
    rowParitiesComputed[i] = rowParitiesComputed[i] % 2 === 0 ? 0 : 1;
  }

  return {
    rowParitiesComputed,
    columnParitiesComputed,
  };
};

const compareParities = (original, computed) => {
  const errorIndex = original.findIndex(
    (bit, index) => bit !== computed[index]
  );
  return {
    index: errorIndex,
    originalBit: original[errorIndex],
    computedBit: computed[errorIndex],
  };
};

const fixParities = (parity, error) => {
  const fixedParity = [...parity];
  const wrongBit = fixedParity[error.index];
  fixedParity[error.index] = wrongBit === 0 ? 1 : 0;
  return fixedParity;
};

const fixMessage = (bits2D, rowError, columnError) => {
  const fixedBits2D = [...bits2D];
  const wrongBit = fixedBits2D[rowError.index][columnError.index];
  fixedBits2D[rowError.index][columnError.index] = wrongBit === 0 ? 1 : 0;
  const fixedMessage = fixedBits2D.flat(1);
  return fixedMessage;
};

function correct(m, n, bits) {
  const size = m * n;
  const message = bits
    .slice(0, size)
    .split("")
    .map((char) => +char);
  const rowParities = bits
    .slice(size, size + m)
    .split("")
    .map((char) => +char);
  const columnParities = bits
    .slice(size + m)
    .split("")
    .map((char) => +char);

  const bits2D = computeMatrix(n, message);

  const { rowParitiesComputed, columnParitiesComputed } = computeParities(
    m,
    n,
    bits2D
  );

  const rowError = compareParities(rowParities, rowParitiesComputed);
  const columnError = compareParities(columnParities, columnParitiesComputed);

  switch (true) {
    case rowError.index === -1 && columnError.index !== -1:
      const fixedColumnParities = fixParities(columnParities, columnError);
      return (
        message.join("") + rowParities.join("") + fixedColumnParities.join("")
      );

    case rowError.index !== -1 && columnError.index === -1:
      const fixedRowParities = fixParities(rowParities, rowError);
      return (
        message.join("") + fixedRowParities.join("") + columnParities.join("")
      );

    case rowError.index !== -1 && columnError.index !== -1:
      const fixedMessage = fixMessage(bits2D, rowError, columnError);
      return (
        fixedMessage.join("") + rowParities.join("") + columnParities.join("")
      );

    default:
      return bits;
  }
}
