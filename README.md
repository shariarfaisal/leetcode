# leetcode

- [Recursion](#recursion)
- [Two Pointers](#two-pointers)
- [Sliding Window](#sliding-window)
- [Linked List](#lineked-list)
- [Queue](#queue)
- [Stack](#stack)
- [Binary Search](#binary-search)
- [Trees](#trees)
- [Sorting](#sorting)
- [Divide and Conquer](#divide-and-conquer)
- [Dynamic Programming](#dynamic-programming)
- [Graph](#graph)
- [Breadth First Search (BFS)](#breadth-first-search-bfs)
- [Depth First Search (DFS)](#depth-first-search-dfs)
- [Heap](#heap)
- [Intervals](#intervals)
- [Greedy](#greedy)
- [Backtracking](#backtracking)
- [Bit Manipulation](#bit-manipulation)
- [Math & Geometry](#math--geometry)
- [Union Find](#union-find)

### Array & Hash Table

### Recursion

Recursion is a programming concept where a function calls itself repeatedly until a certain condition is met. It allows you to solve complex problems by breaking them down into smaller, more manageable subproblems. Here's an example of a recursive function in JavaScript:

```javascript
function countdown(n) {
  // Base case: stop when n reaches 0
  if (n === 0) {
    console.log("Countdown complete!");
  } else {
    console.log(n);
    countdown(n - 1); // Recursive call with a smaller value
  }
}

countdown(5);
```

In this example, the countdown function takes an integer n as input and prints the countdown from n to 1. It starts by checking if n is equal to 0, which is the base case. If n is 0, it prints "Countdown complete!" and stops the recursion. Otherwise, it prints the current value of n and makes a recursive call to countdown with n - 1 as the argument. Each recursive call reduces the value of n until it eventually reaches 0 and completes the countdown.

### Two Pointers

The "two pointers" technique is commonly used to solve problems that involve searching, manipulating, or comparing elements in an array or a linked list. It involves using two pointers that traverse the data structure in different ways to solve the problem efficiently. Let's take a look at an example that demonstrates the two pointers technique in JavaScript:

```javascript
function findPairSum(arr, target) {
  // Sort the array in ascending order
  arr.sort((a, b) => a - b);

  let left = 0; // Pointer starting from the beginning
  let right = arr.length - 1; // Pointer starting from the end

  while (left < right) {
    const sum = arr[left] + arr[right];

    if (sum === target) {
      // Found a pair with the target sum
      return [arr[left], arr[right]];
    } else if (sum < target) {
      // The sum is smaller, move the left pointer to increase the sum
      left++;
    } else {
      // The sum is larger, move the right pointer to decrease the sum
      right--;
    }
  }

  // No pair found with the target sum
  return null;
}

const arr = [2, 4, 6, 8, 10];
const target = 12;
const pair = findPairSum(arr, target);

if (pair) {
  console.log(`Pair found: ${pair[0]} and ${pair[1]}`);
} else {
  console.log("No pair found with the target sum.");
}

```

In this example, we have an array `arr` containing numbers and a target sum `target`. The `findPairSum` function uses the two pointers technique to find a pair of numbers in the array that add up to the target sum.

We start by sorting the array in ascending order using the `sort` method. Then, we initialize two pointers: `left` points to the beginning of the array, and `right` points to the end of the array.

Inside the `while` loop, we calculate the sum of the numbers at the current positions of the two pointers. If the sum is equal to the target, we've found a pair that satisfies the condition, and we return the pair.

If the sum is smaller than the target, we move the `left` pointer one step forward to increase the sum. If the sum is larger than the target, we move the `right` pointer one step backward to decrease the sum. We continue this process until the two pointers meet or the target sum is found.

Finally, if no pair is found with the target sum, we return `null`.

### Sliding Window

The "sliding window" technique is useful for solving problems that involve finding a subset or subarray of consecutive elements in an array or string that satisfies certain conditions. It involves maintaining a window of elements while "sliding" or moving the window across the data structure to find the desired subset. Let's take a look at an example that demonstrates the sliding window technique in JavaScript:

```javascript
function findMaxSum(arr, k) {
  const n = arr.length;

  // Handle edge case
  if (n < k) {
    return null;
  }

  let windowSum = 0;
  let maxSum = 0;

  // Compute the sum of the first window
  for (let i = 0; i < k; i++) {
    windowSum += arr[i];
  }

  maxSum = windowSum;

  // Slide the window and update the sums
  for (let i = k; i < n; i++) {
    windowSum = windowSum - arr[i - k] + arr[i];
    maxSum = Math.max(maxSum, windowSum);
  }

  return maxSum;
}

const arr = [2, 1, 5, 1, 3, 2];
const k = 3;
const maxSum = findMaxSum(arr, k);

console.log(`Maximum sum of a subarray of size ${k}: ${maxSum}`);
```

In this example, we have an array arr and a window size k. The findMaxSum function uses the sliding window technique to find the maximum sum of a subarray of size k in the array.

We first handle the edge case where the array length is smaller than k. In such cases, there won't be a valid subarray, so we return null.

Next, we initialize windowSum and maxSum variables to keep track of the current sum of the window and the maximum sum found so far.

We compute the sum of the first window by iterating through the array from index 0 to k-1 and adding each element to windowSum.

Then, we set maxSum to the initial windowSum value.

Next, we slide the window across the array. Starting from index k, we subtract the element at i-k from windowSum and add the element at index i. This effectively moves the window one step forward.

At each step, we update maxSum by comparing it with the current windowSum and taking the maximum value.

Finally, we return maxSum, which represents the maximum sum of a subarray of size k in the given array.

```
Maximum sum of a subarray of size 3: 9
```

In the given array [2, 1, 5, 1, 3, 2], the subarray [5, 1, 3] has the maximum sum of 9 when k is 3.

### Lineked List

```javascript
// Define the Node class
class Node {
  constructor(data) {
    this.data = data;
    this.next = null;
  }
}

// Define the LinkedList class
class LinkedList {
  constructor() {
    this.head = null;
    this.tail = null;
  }

  // Method to add a node to the end of the linked list
  append(data) {
    const newNode = new Node(data);

    if (!this.head) {
      // If the list is empty, set the new node as the head and tail
      this.head = newNode;
      this.tail = newNode;
    } else {
      // If the list is not empty, append the new node to the tail
      this.tail.next = newNode;
      this.tail = newNode;
    }
  }

  // Method to display the linked list elements
  display() {
    let current = this.head;

    while (current) {
      console.log(current.data);
      current = current.next;
    }
  }
}

// Create a new linked list instance
const linkedList = new LinkedList();

// Append nodes to the linked list
linkedList.append(10);
linkedList.append(20);
linkedList.append(30);

// Display the linked list
linkedList.display();

```

In this example, we define two classes: `Node` and `LinkedList`.

The `Node` class represents a node in the linked list. It has two properties: `data` to store the node's value, and `next` to reference the next node in the list.

The `LinkedList` class represents the linked list itself. It has two properties: `head` to reference the first node in the list, and `tail` to reference the last node in the list.

The `LinkedList` class has two methods:

- The `append` method adds a new node with the given data to the end of the linked list. It first creates a new node using the `Node` class, and then checks if the list is empty. If it is, the new node becomes both the head and tail of the list. If the list is not empty, the new node is appended after the current tail, and the tail is updated to the new node.

- The `display` method traverses the linked list starting from the head and prints the data of each node until it reaches the end of the list.

In the main program, we create a new instance of the `LinkedList` class named `linkedList`. We then use the `append` method to add nodes with values 10, 20, and 30 to the linked list. Finally, we call the `display` method to print the values of all nodes in the linked list.

### Queue

```javascript
// Define the Queue class
class Queue {
  constructor() {
    this.items = [];
  }

  // Method to add an element to the queue
  enqueue(element) {
    this.items.push(element);
  }

  // Method to remove an element from the queue
  dequeue() {
    if (this.isEmpty()) {
      return null;
    }
    return this.items.shift();
  }

  // Method to check if the queue is empty
  isEmpty() {
    return this.items.length === 0;
  }

  // Method to get the front element of the queue
  front() {
    if (this.isEmpty()) {
      return null;
    }
    return this.items[0];
  }

  // Method to get the size of the queue
  size() {
    return this.items.length;
  }

  // Method to display the queue elements
  display() {
    console.log(this.items);
  }
}

// Create a new queue instance
const queue = new Queue();

// Enqueue elements to the queue
queue.enqueue(10);
queue.enqueue(20);
queue.enqueue(30);

// Display the queue
queue.display(); // [10, 20, 30]

// Dequeue an element from the queue
const dequeuedElement = queue.dequeue();
console.log("Dequeued element:", dequeuedElement); // Dequeued element: 10

// Display the updated queue
queue.display(); // [20, 30]

// Get the front element of the queue
const frontElement = queue.front();
console.log("Front element:", frontElement); // Front element: 20

// Get the size of the queue
const queueSize = queue.size();
console.log("Queue size:", queueSize); // Queue size: 2

// Check if the queue is empty
const isEmpty = queue.isEmpty();
console.log("Is the queue empty?", isEmpty); // Is the queue empty? false
```

In this example, we define a `Queue` class that represents a queue data structure.

The `Queue` class has an `items` array as its underlying data structure.

The `Queue` class has several methods:

- The `enqueue` method adds an element to the end of the queue by using the `push` method of the items array.

- The `dequeue` method removes and returns the element at the front of the queue using the `shift` method of the `items` array. If the queue is empty, it returns `null`.

- The `isEmpty` method checks if the queue is empty by examining the length of the `items` array.

- The `front` method returns the element at the front of the queue without removing it. If the queue is empty, it returns `null`.

- The `size` method returns the number of elements in the queue by using the `length` property of the `items` array.

- The `display` method prints the elements of the queue using `console.log`.

In the main program, we create a new instance of the `Queue` class named `queue`. We then use the `enqueue` method to add elements 10, 20, and 30 to the queue. After each operation, we display the updated queue using the `display` method.

We also demonstrate other queue operations such as `dequeue`, `front`, `size`, and `isEmpty`, and `display` their respective outputs.

When you run this code, it will output:

```
[10, 20, 30]
Dequeued element: 10
[20, 30]
Front element: 20
Queue size: 2
Is the queue empty? false
```

### Stack

### Binary Search

### Trees

### Sorting

### Divide And Conquer

### Dynamic Programming

### Graph

### Breadth First Search (BFS)

### Depth First Search (DFS)

### Heap

### Intervals

### Greedy

### Backtracking

### Bit Manipulation

### Math & Geometry

### Union Find

### Intervals

### Greedy

### Backtracking

### Bit Manipulation

### Math & Geometry

### Union Find
