[Count Days Without Meetings](https://leetcode.com/problems/count-days-without-meetings)
===
You are given a positive integer `days` representing the total number of days an employee is available for work (
starting from day 1). You are also given a 2D array `meetings` of size `n` where, `meetings[i] = [start_i, end_i]`
represents the starting and ending days of meeting `i` (inclusive).

Return _the count of days when the employee is available for work but no meetings are scheduled_.

**Note**: The meetings may overlap.

**Example 1:**

```text
Input: days = 10, meetings = [[5,7],[1,3],[9,10]]
Output: 2
Explanation:
There is no meeting scheduled on the 4th and 8th days.
```

**Example 2:**

```text
Input: days = 5, meetings = [[2,4],[1,3]]
Output: 1
Explanation:
There is no meeting scheduled on the 5th day.
```

**Example 3:**

```text
Input: days = 6, meetings = [[1,6]]
Output: 0
Explanation:
Meetings are scheduled for all working days.
```

**Constraints:**

* `1 <= days <= 1e9`
* `1 <= meetings.length <= 1e5`
* `meetings[i].length == 2`
* `1 <= meetings[i][0] <= meetings[i][1] <= days`

