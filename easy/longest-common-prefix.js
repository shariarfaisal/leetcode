/**
 * 14. Longest Common Prefix
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function (strs) {
  function x(s, a, b) {
    return strs.every((str) => {
      return str.substring(a, b) === s;
    });
  }

  let res = "";
  let min = Math.min(...strs.map((i) => i.length));
  for (let i = 0; i <= min; i++) {
    const is = x(strs[0].substring(0, i), 0, i);
    console.log(strs[0].substring(0, i));
    if (!is) {
      res = strs[0].substring(0, i - 1);
      break;
    } else {
      res = strs[0].substring(0, i);
    }
  }

  return res;
};
