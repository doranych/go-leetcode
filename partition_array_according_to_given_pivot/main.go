package partition_array_according_to_given_pivot

func pivotArray(nums []int, pivot int) []int {
	n := len(nums)
	res := make([]int, n)
	lI := 0
	gI := n - 1
	for i, j := 0, n-1; i < n && j >= 0; i, j = i+1, j-1 {
		if nums[i] < pivot {
			res[lI] = nums[i]
			lI++
		}
		if nums[j] > pivot {
			res[gI] = nums[j]
			gI--
		}
	}
	for lI <= gI {
		res[lI] = pivot
		lI++
	}
	return res
}

/** naive solution
func pivotArray(nums []int, pivot int) []int {
    s1 := make([]int, 0)
    s2 := make([]int, 0)
    pCount := 0
    for i := range nums {
        if nums[i] < pivot {
            s1 = append(s1, nums[i])
        } else if nums[i] > pivot {
            s2 = append(s2, nums[i])
        } else {
            pCount++
        }
    }
    res := make([]int, 0, len(nums))
    for i := range s1 {
        res = append(res, s1[i])
    }
    for _ = range pCount {
        res = append(res, pivot)
    }
    for i := range s2 {
        res = append(res, s2[i])
    }
    return res
}
*/
