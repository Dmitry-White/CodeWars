/**
 * Created on Tue May 30 22:56:14 2023
 * @author: Dmitry White
 */

/**
 * TODO: You will be passed the dice value frequencies,
 * and your task is to write the code to return a string representing a histogram,
 * so that when it is printed it has the following format:
 * Input: [7,4,10,1,0,5]
 * Output:
 *  6|##### 5
 *  5|
 *  4|# 1
 *  3|########## 10
 *  2|### 3
 *  1|####### 7
 */

function histogram(results) {
  const resultsArr = [...results];
  const histogramArr = [];

  while (resultsArr.length > 0) {
    const index = resultsArr.length;
    const curr = resultsArr.pop();
    const bin = Array(curr).fill("#").join("");

    const histogramItem =
      curr === 0 ? `${index}|` : `${index}|${bin} ${curr ?? curr}`;
    histogramArr.push(histogramItem);
  }

  return histogramArr.join("\n") + "\n";
}
