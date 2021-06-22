package main

func main() {

}

func predictPartyVictory(senate string) string {
	raQueue := make([]int, 0)
	diQueue := make([]int, 0)
	senLen := len(senate)
	for i := 0; i < senLen; i++ {
		if senate[i] == 'R' {
			raQueue = append(raQueue, i)
		} else {
			diQueue = append(diQueue, i)
		}
	}

	for len(raQueue) != 0 && len(diQueue) != 0 {
		if raQueue[0] < diQueue[0] {
			tmp := raQueue[0]
			raQueue = raQueue[1:]
			raQueue = append(raQueue, tmp+senLen)
			diQueue = diQueue[1:]
		} else {
			tmp := diQueue[0]
			diQueue = diQueue[1:]
			diQueue = append(diQueue, tmp+senLen)
			raQueue = raQueue[1:]
		}
	}
	if len(raQueue) == 0 {
		return "Dire"
	}
	return "Radiant"
}

