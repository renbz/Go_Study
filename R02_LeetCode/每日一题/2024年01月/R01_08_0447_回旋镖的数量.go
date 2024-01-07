package main

func numberOfBoomerangs(points [][]int) (ans int) {
	if len(points) == 1 {
		return 0
	}
	for _, p1 := range points {
		cntMap := make(map[int]int, len(points))
		for _, p2 := range points {
			dx, dy := p1[0]-p2[0], p1[1]-p2[1]
			cntMap[dx*dx+dy*dy]++
		}
		for _, m := range cntMap {
			if m <= 1 {
				continue
			}
			ans += m * (m - 1)
		}
	}
	return ans
}
