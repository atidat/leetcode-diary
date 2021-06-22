package main

func main() {

}

func canCompleteCircuit(gas []int, cost []int) int {
	res := -1
	for i := 0; i < len(gas); i++ {
		if gas[i] < cost[i] {
			continue
		}

		if handle(gas, cost, i, &res) {
			//fmt.Println(res)
			break
		}
	}
	return res
}

// return true: find available start, end process
func handle(gas []int, cost []int, s int, res *int) bool {
	rest := gas[s] - cost[s]
	for i := s+1; i != s; i++{
		if i == len(gas) {
			i = -1
			continue
		}
		rest +=  gas[i] - cost[i]
		if rest < 0 {
			break
		}
	}
	if rest < 0 {
		return false
	}
	*res = s
	return true
}