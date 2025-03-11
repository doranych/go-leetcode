[Number of Substrings Containing All Three Characters](https://leetcode.com/problems/number-of-substrings-containing-all-three-characters)
===
Given a string `s` consisting only of characters `a`, `b` and `c`.

Return _the number of substrings containing **at least** one occurrence of all these characters `a`, `b` and `c`._

**Example 1:**

```text
Input: s = "abcabc"
Output: 10
Explanation: The substrings containing at least one occurrence of the characters a, b and c are "abc", "abca",
 "abcab", "abcabc", "bca", "bcab", "bcabc", "cab", "cabc" and "abc" (again).
```

**Example 2:**

```text
Input: s = "aaacb"
Output: 3
Explanation: The substrings containing at least one occurrence of the characters a, b and c are 
 "aaacb", "aacb" and "acb".
```

**Example 3:**

```text
Input: s = "abc"
Output: 1
```

**Constraints:**

* `3 <= s.length <= 5*10e4`
* `s` only consists of `a`, `b` or `c` characters.

