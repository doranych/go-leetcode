package solving_questions_with_brainpower

// dp is slow and memory inefficient because of call stack and heap usage
// func mostPoints(questions [][]int) int64 {
//     n := len(questions)
//     dp := make([]int, n)

//     var dfs func(int) int
//     dfs = func(i int) int {
//         if i >= n {
//             return 0
//         }
//         if dp[i] > 0 {
//             return dp[i]
//         }
//         p, s := questions[i][0], questions[i][1]
//         dp[i] = max(dfs(i+1), p + dfs(i+s+1))
//         return dp[i]
//     }
//     return int64(dfs(0))
// }

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int, n)
	dp[n-1] = questions[n-1][0]

	for i := n - 2; i >= 0; i-- {
		dp[i] = questions[i][0]
		s := questions[i][1]
		if i+s+1 < n {
			dp[i] += dp[i+s+1]
		}
		dp[i] = max(dp[i], dp[i+1])
	}
	return int64(dp[0])
}
