[Longest Strictly Increasing or Strictly Decreasing Subarray](https://leetcode.com/problems/longest-strictly-increasing-or-strictly-decreasing-subarray)
===
You are given an array of integers `nums`. Return _the length of the **longest subarray** of `nums` which is
either **strictly increasing** or **strictly decreasing**._

An array is said to be **strictly increasing** if each element is strictly **greater** than its previous one (if exists).

An array is said to be **strictly decreasing** if each element is strictly **smaller** than its previous one (if exists).

**Example 1:**

```text
Input: nums = [1,4,3,3,2]
Output: 2
Explanation:
The strictly increasing subarrays of nums are [1], [2], [3], [3], [4], and [1,4].
The strictly decreasing subarrays of nums are [1], [2], [3], [3], [4], [3,2], and [4,3].
Hence, we return 2.
```

**Example 2:**

```text
Input: nums = [3,3,3,3]
Output: 1
Explanation:
The strictly increasing subarrays of nums are [3], [3], [3], and [3].
The strictly decreasing subarrays of nums are [3], [3], [3], and [3].
Hence, we return 1.
```

**Example 3:**

```text
Input: nums = [3,2,1]
Output: 3
Explanation:
The strictly increasing subarrays of nums are [3], [2], and [1].
The strictly decreasing subarrays of nums are [3], [2], [1], [3,2], [2,1], and [3,2,1].
Hence, we return 3.
```

**Constraints:**

* `1 <= nums.length <= 50`
* `1 <= nums[i] <= 50`
