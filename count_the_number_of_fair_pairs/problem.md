[Count the Number of Fair Pairs](https://leetcode.com/problems/count-the-number-of-fair-pairs)
===
Given a **0-indexed** integer array `nums` of size `n` and two integers `lower` and `upper`, return *the number of fair
pairs*.

A pair `(i, j)` is **fair** if:

- `0 <= i < j < n`, and
- `lower <= nums[i] + nums[j] <= upper`

### Example 1:

```text
Input: nums = [0,1,7,4,4,5], lower = 3, upper = 6
Output: 6
Explanation: There are 6 fair pairs: (0,3), (0,4), (0,5), (1,3), (1,4), and (1,5).
```

### Example 2:

```text
Input: nums = [1,7,9,2,5], lower = 11, upper = 11
Output: 1
Explanation: There is a single fair pair: (2,3).
```

### Constraints:

- `1 <= nums.length <= 1e5`
- `nums.length == n`
- `-1e9 <= nums[i] <= 1e9`
- `-1e9 <= lower <= upper <= 1e9`