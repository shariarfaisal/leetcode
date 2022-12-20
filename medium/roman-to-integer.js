/**
 * 13. Roman to Integer
 * @param {string} s
 * @return {number}
 */
var romanToInt = function (s) {
  var matrixToRoman = [
    [1000, "M"],
    [900, "CM"],
    [500, "D"],
    [400, "CD"],
    [100, "C"],
    [90, "XC"],
    [50, "L"],
    [40, "XL"],
    [10, "X"],
    [9, "IX"],
    [5, "V"],
    [4, "IV"],
    [1, "I"],
  ];

  function convertToRoman(s) {
    if (s === 0) {
      return 0;
    }

    let str;
    const v = matrixToRoman.find(([n, v]) => {
      if (v === s[0]) {
        str = s.substr(1);
        return true;
      } else if (v === s.substr(0, 2)) {
        str = s.substr(2);
        return true;
      } else {
        return false;
      }
    });

    if (!v) return 0;
    return Number(v[0]) + Number(convertToRoman(str));
  }

  return convertToRoman(s);
};
