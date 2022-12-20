/**
 * 1480. Running Sum of 1d Array
 * @param {number[]} nums
 * @return {number[]}
 */
var runningSum = function (nums) {
  let sum = 0;
  return nums.map((i) => {
    sum += i;
    return sum;
  });
};
