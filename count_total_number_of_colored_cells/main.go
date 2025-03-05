package count_total_number_of_colored_cells

func coloredCells(n int) int64 {
	return int64(1 + 2*(n-1)*n)
}

/* intuitive solution
func coloredCells(n int) int64 {
    sum := 1
    for i := 2; i <= n; i++ {
        sum += (i-1)*4
    }
    return int64(sum)
}
*/
