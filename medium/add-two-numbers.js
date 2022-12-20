/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 *
 * Input: l1 = [2,4,3], l2 = [5,6,4]
 * Output: [7,0,8]
 * Explanation: 342 + 465 = 807.
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function (l1, l2) {
  let a = "";
  let l1Head = l1;
  while (l1Head && l1Head !== undefined) {
    a = String(l1Head.val) + a;
    l1Head = l1Head.next;
  }
  a = BigInt(a);
  let b = "";
  let l2Head = l2;
  while (l2Head && l2Head.val !== undefined) {
    b = String(l2Head.val) + b;
    l2Head = l2Head.next;
  }
  b = BigInt(b);
  let sum = a + b;
  let split = sum.toLocaleString().replaceAll(",", "").split("");
  let head = null;
  split.forEach((v) => {
    const node = {
      val: parseInt(v),
      next: head,
    };
    head = node;
  });
  return head;
};
