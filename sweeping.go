package minesweeper

func Sweeping(no []int) int {
	var countMine int
	for i := 0; i <= len(no)-1; i++ {
		if no[i] == 1 {
			countMine += 1
			if i != 0 {
				if no[i-1] == 1 {
					countMine -= 1
				}
			}
			if i < 2 == false {
				if no[i-2] == 2 {
					countMine += 1
				}
			}

		}
	}
	return countMine
}
