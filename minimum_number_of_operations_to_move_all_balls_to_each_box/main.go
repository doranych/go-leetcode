package minimum_number_of_operations_to_move_all_balls_to_each_box

// func minOperations(boxes string) []int {
//     // naive solution would be calculate amount of movement between
//     // each boxes and each non empty box
//     // that would be n squared solution in the worst case
//     // as amount of moves eq abs(i-j) for each box and ball
//     n := len(boxes)
//     answ := make([]int, n)
//     for i:=0; i<n; i++ {
//         for j:= 0; j<n; j++ {
//             if boxes[j] == '1' {
//                 answ[i] += int(math.Abs(float64(i-j)))
//             }
//         }
//     }
//     return answ
// }

// let's try to find linear solution, which looks possible
func minOperations(boxes string) []int {
	n := len(boxes)
	res := make([]int, n)
	sum := 0
	balls := 0
	for i := 0; i < n; i++ {
		res[i] = sum + balls
		sum += balls
		if boxes[i] == '1' {
			balls++
		}
	}

	sum = 0
	balls = 0
	for i := n - 1; i >= 0; i-- {
		res[i] = res[i] + sum + balls
		sum += balls
		if boxes[i] == '1' {
			balls++
		}
	}
	return res
}
