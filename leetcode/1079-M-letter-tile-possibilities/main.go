package main

func main() {

}

/*
tiles = "AAB"
then f(tiles) = executor(tiles[:len(tiles)-1], tiles[len(tiles)-1])
			  = executor(
*/
func numTilePossibilities(tiles string) int {
	res := executor(tiles[:len(tiles)-1], tiles[len(tiles)-1])
	store := make(map[string]struct{})
	for i := range res {
		if _, ok := store[res[i]]; !ok {
			store[res[i]] = struct{}{}
		}
	}
	//fmt.Println(store)
	return len(store)-1
}

func executor(tiles string, b byte) []string {
	//fmt.Printf("input: %v, %c\n", tiles, b)
	if len(tiles) == 0 {
		return []string{"", string(b)}
	}
	res := executor(tiles[:len(tiles)-1], tiles[len(tiles)-1])
	//fmt.Printf("ha 1: %v", res)
	resLen := len(res)
	for i := 0; i < resLen; i++ {
		res = append(res, convert(res[i], b)...)
	}
	//fmt.Printf("ha 2: %v", res)
	return res
}

func convert(tar string, b byte) []string {
	tars := []byte(tar)
	res := []string{}
	for i := 0; i <= len(tars); i++ {
		tmp := string(tars[:i]) + string(b) + string(tars[i:len(tars)])
		res = append(res, tmp)
	}
	return res
}