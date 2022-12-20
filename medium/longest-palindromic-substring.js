/**
 * Problem: 5. Longest Palindromic Substring
 Given a string s, return the longest palindromic substring in s
 Example 1:
    Input: s = "babad"
    Output: "bab"
    Explanation: "aba" is also a valid answer.


 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function (s) {
  let arr = s.split("");
  const len = arr.length;

  const compare = (l, r, t) => {
    let text = t;
    while (l > -1 && r <= len && arr[l] === arr[r]) {
      text = arr[l] + text + arr[r];
      l--;
      r++;
    }
    return text;
  };

  let text = "";
  arr.forEach((s, i) => {
    let v = compare(i - 1, i + 1, s);
    if (v.length >= text.length) {
      text = v;
    }
    if (s === arr[i + 1]) {
      let v = compare(i - 1, i + 2, s + s);
      if (v.length >= text.length) {
        text = v;
      }
    }
  });

  return text;
};
