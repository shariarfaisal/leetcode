/*
 * @lc app=leetcode id=9 lang=javascript
 *
 * [9] Palindrome Number
 */

// @lc code=start
function isPalindrome(x) {
  if (x < 0) {
    return false;
  }

  let num = x;
  let reverse = 0;

  while (num > 0) {
    reverse = reverse * 10 + (num % 10);
    num = Math.floor(num / 10);
  }

  return reverse === x;
}

// @lc code=end
