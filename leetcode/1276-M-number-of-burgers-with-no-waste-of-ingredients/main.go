package main

func main() {

}
func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	if tomatoSlices%2 == 1 || (tomatoSlices <= cheeseSlices && tomatoSlices != 0) {
		return []int{}
	}
	for i := 0; i <= cheeseSlices; i++ {
		if 4*i + 2*(cheeseSlices-i) == tomatoSlices {
			return []int{i, cheeseSlices-i}
		}
	}
	return []int{}
}