package main

func main() {

}

func robotSim(commands []int, obstacles [][]int) int {
	var stepLen int

	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	ind := 0

	trap := make(map[[2]int]bool, 0)
	for _, v := range obstacles {
		trap[[2]int{v[0], v[1]}] = true
	}
	var x, y int
	for i := 0; i < len(commands); i++ {
		if commands[i] == -1 {
			ind = (ind + 1)%4
		} else if commands[i] == -2 {
			ind = (ind + 4 - 1)%4
		} else {
			for e := 0; e < commands[i]; e++ {
				x += dx[ind]
				y += dy[ind]
				if _, ok := trap[[2]int{x, y}]; ok {
					x -= dx[ind]
					y -= dy[ind]
				}
				stepLen = mmax(stepLen, euler(x,y))
			}
		}
	}
	return stepLen
}

func mmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func euler(a, b int) int {
	return a*a + b*b
}