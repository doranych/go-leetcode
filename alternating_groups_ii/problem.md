[Alternating Groups II](https://leetcode.com/problems/alternating-groups-ii)
===
There is a circle of red and blue tiles. You are given an array of integers `colors` and an integer `k`. The color of
tile `i` is represented by `colors[i]`:

* `colors[i] == 0` means that tile `i` is **red**.
* `colors[i] == 1` means that tile `i` is **blue**.

An **alternating** group is every `k` contiguous tiles in the circle with **alternating** colors (each tile in the group
except the first and last one has a different color from its **left** and **right** tiles).

Return _the number of **alternating** groups_.

**Note** that since `colors` represents a **circle**, the **first** and the **last** tiles are considered to be next to
each other.

**Example 1:**

```text
Input: colors = [0,1,0,1,0], k = 3
Output: 3
```

![img.png](../bin/alternating_groups_ii/img.png)

![img_1.png](../bin/alternating_groups_ii/img_1.png)

![img_2.png](../bin/alternating_groups_ii/img_2.png)

![img_3.png](../bin/alternating_groups_ii/img_3.png)

**Example 2:**

```text
Input: colors = [0,1,0,0,1,0,1], k = 6
Output: 2
```

**Example 3:**

```text
Input: colors = [1,1,0,1], k = 4
Output: 0
```

**Constraints:**

* `3 <= colors.length <= 1e5`
* `0 <= colors[i] <= 1`
* `3 <= k <= colors.length`

