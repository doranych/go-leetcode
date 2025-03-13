[Maximum Count of Positive Integer and Negative Integer](https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer)
===
Given an array `nums` sorted in **non-decreasing** order, return _the **maximum** between the number of positive
integers and the number of negative integers_.

* In other words, if the number of positive integers in `nums` is `pos` and the number of negative integers is `neg`,
  then return _the maximum of `pos` and `neg`_.

**Note** that `0` is neither positive nor negative.

**Example 1:**

```text
Input: nums = [-2,-1,-1,1,2,3]
Output: 3
Explanation: There are 3 positive integers and 3 negative integers. The maximum count among them is 3.
```

**Example 2:**

```text
Input: nums = [-3,-2,-1,0,0,1,2]
Output: 3
Explanation: There are 2 positive integers and 3 negative integers. The maximum count among them is 3.
```

**Example 3:**

```text
Input: nums = [5,20,66,1314]
Output: 4
Explanation: There are 4 positive integers and 0 negative integers. The maximum count among them is 4.
```

**Constraints:**

* `1 <= nums.length <= 2000`
* `-2000 <= nums[i] <= 2000`
* `nums` is sorted in a **non-decreasing** order.
