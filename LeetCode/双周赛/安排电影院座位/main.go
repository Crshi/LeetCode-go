package main

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	reserverdMap := make(map[int]map[int]bool)
	for _, seat := range reservedSeats {
		if seat[1] == 1 || seat[1] == 10 {
			continue
		}
		if reserverdMap[seat[0]] == nil {
			reserverdMap[seat[0]] = make(map[int]bool)
		}
		reserverdMap[seat[0]][seat[1]] = true
	}
	count := 0
	for _, m := range reserverdMap {
		useMiddle := false
		if (!m[2]) && (!m[3]) && (!m[4]) && (!m[5]) {
			useMiddle = true
			count++
		}
		if (!m[6]) && (!m[7]) && (!m[8]) && (!m[9]) {
			useMiddle = true
			count++
		}
		if (!m[4]) && (!m[5]) && (!m[6]) && (!m[7]) && (!useMiddle) {
			count++
		}
	}
	count += 2 * (n - len(reserverdMap))
	return count
}

func main() {
	//3, reservedSeats = [[1,2],[1,3],[1,8],[2,6],[3,1],[3,10]]
	b := [][]int{{1, 2}, {1, 3}, {1, 8}, {2, 6}, {3, 1}, {3, 10}}

	println(maxNumberOfFamilies(3, b))
}
